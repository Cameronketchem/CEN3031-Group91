package db

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/mattn/go-sqlite3"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"sync"
	"time"
)

// busy is the time to wait for a sqlite lock from another process, in ms.
// This causes sqlite to wait before returning SQLITE_BUSY.  On the other
// hand, conflicts with other connections from the same process
// might contend on the shared cache, which corresponds to SQLITE_LOCKED and
// is not covered by the busy timeout.  We rely on sqlite_unlock_notify()
// to wait for the shared cache lock to be released.  This is enabled in
// go-sqlite3 with the "sqlite_unlock_notify" Go build tag.
const busy = 1000

// How many times to retry a db transaction before an error is issued.
const infoTxRetries = 5

// Warn about repeatedly failing transactions every 5 failures
const warnTxRetriesInterval = 5

// enableFullfsyncStatements is a list of statements we execute to enable a fullfsync.
// Currently, it's only supported by MacOSX.
var enableFullfsyncStatements []string
var sqliteInitOnce sync.Once

// A function for the database Accessor to execute. Errors must be of the sql.Error type
type idemFn func(ctx context.Context, tx *sql.Tx) error

// txExecutionContext contains the data that is associated with every created transaction
// before sending it to the user-defined callback. This allows the callback function to
// make changes to the execution setting of an ongoing transaction.
type txExecutionContext struct {
	deadline time.Time
}

// URI converts a filename and db options into an SQLITE3 connection URI.
func URI(filename string, readOnly bool, memory bool) string {
	filename = filepath.FromSlash(filename)
	uri := fmt.Sprintf("file:%s?_busy_timeout=%d&_synchronous=full", filename, busy)
	if !readOnly {
		uri += "&_txlock=immediate"
	}
	if memory {
		uri += "&mode=memory"
		uri += "&cache=shared"
	}
	return uri
}

// An Accessor manages a sqlite database handle and any outstanding
// batching operations.
type Accessor struct {
	Handle   *sql.DB
	readOnly bool
	inMemory bool
}

// MakeAccessor creates a new Accessor.
func MakeAccessor(dbfilename string, readOnly bool, inMemory bool) (Accessor, error) {
	return makeAccessorImpl(dbfilename, readOnly, inMemory, []string{"_journal_mode=wal"})
}

// MakeErasableAccessor creates a new Accessor with the secure_delete pragma set;
// see https://www.sqlite.org/pragma.html#pragma_secure_delete
// It is not read-only and not in-memory (otherwise, erasability doesn't matter)
func MakeErasableAccessor(dbfilename string) (Accessor, error) {
	return makeErasableAccessor(dbfilename, false)
}

func makeErasableAccessor(dbfilename string, readOnly bool) (Accessor, error) {
	return makeAccessorImpl(dbfilename, readOnly, false, []string{"_secure_delete=on", "_journal_mode=wal"})
}

func makeAccessorImpl(dbfilename string, readOnly bool, inMemory bool, params []string) (Accessor, error) {
	var db Accessor
	db.readOnly = readOnly
	db.inMemory = inMemory

	var err error
	db.Handle, err = sql.Open("sqlite3", URI(dbfilename, readOnly, inMemory)+"&"+strings.Join(params, "&"))

	if err == nil {
		// create a connection to safely initialize SQLite once
		initFn := func() {
			var conn *sql.Conn
			if conn, err = db.Handle.Conn(context.Background()); err != nil {
				db.Close()
				return
			}
			if err = conn.Close(); err != nil {
				db.Close()
			}
		}
		sqliteInitOnce.Do(initFn)
		if err != nil {
			// init failed, db closed and err is set
			return db, err
		}
		err = db.SetSynchronousMode(context.Background(), SynchronousModeFull, true)
		if err != nil {
			db.Close()
		}
	}

	return db, err
}

// Close closes the connection.
func (db *Accessor) Close() {
	db.Handle.Close()
	db.Handle = nil
}

// dbretry returns true if the error might be temporary
func dbretry(obj error) bool {
	var sqliteErr sqlite3.Error
	if errors.As(obj, &sqliteErr) {
		return sqliteErr.Code == sqlite3.ErrLocked || sqliteErr.Code == sqlite3.ErrBusy
	}

	return false // Not an sqlite error type
}

// LoggedRetry executes a function repeatedly as long as it returns an error
// that indicates database contention that warrants a retry.
// Prints warnings and errors.
func LoggedRetry(fn func() error) (err error) {
	for i := 0; (i == 0) || dbretry(err); i++ {
		if i > 0 {
			if i < infoTxRetries {
				fmt.Printf("[INFO] db.LoggedRetry: %d retries (last err: %v)\n", i, err)
			} else if i >= 1000 {
				fmt.Printf("[FAILURE] db.LoggedRetry: %d retries (last err: %v)\n", i, err)
				return
			} else if i%warnTxRetriesInterval == 0 {
				fmt.Printf("[WARNING] db.LoggedRetry: %d retries (last err: %v)\n", i, err)
			}
		}
		err = fn()
	}
	return
}

// Retry executes a function repeatedly as long as it returns an error
// that indicates database contention that warrants a retry.
func Retry(fn func() error) (err error) {
	return LoggedRetry(fn)
}

// Atomic executes a piece of code with respect to the database atomically.
// For transactions where readOnly is false, sync determines whether or not to wait for the result.
// The return error of fn should be a native sqlite3.Error type or an error wrapping it.
// DO NOT return a custom error - the internal logic of Atomic expects an sqlite error and uses that value.
func (db *Accessor) Atomic(fn idemFn, extras ...interface{}) (err error) {
	return db.AtomicContext(context.Background(), fn, extras...)
}

// AtomicContext executes a piece of code with respect to the database atomically.
// For transactions where readOnly is false, sync determines whether or not to wait for the result.
// Like for Atomic, the return error of fn should be a native sqlite3.Error type or an error wrapping it.
func (db *Accessor) AtomicContext(ctx context.Context, fn idemFn, extras ...interface{}) (err error) {
	atomicDeadline := time.Now().Add(time.Second)

	// note that the sql library will drop panics inside an active transaction
	guardedFn := func(ctx context.Context, tx *sql.Tx) (err error) {
		defer func() {
			if r := recover(); r != nil {
				var ok bool
				err, ok = r.(error)
				if !ok {
					err = fmt.Errorf("%v", r)
				}

				buf := make([]byte, 16*1024)
				stlen := runtime.Stack(buf, false)
				errstr := string(buf[:stlen])
				fmt.Fprintf(os.Stderr, "recovered panic in atomic: %s", errstr)

			}
		}()

		err = fn(ctx, tx)
		return
	}

	var tx *sql.Tx
	var conn *sql.Conn

	// Establish a connection to the database
	for i := 0; (i == 0) || dbretry(err); i++ {
		if i > 0 {
			if i < infoTxRetries {
				fmt.Printf("[INFO] db.atomic: %d connection retries (last err: %v)\n", i, err)
			} else if i >= 1000 {
				fmt.Printf("[FAILURE] db.atomic: %d connection retries (last err: %v)\n", i, err)
				break
			} else if i%warnTxRetriesInterval == 0 {
				fmt.Printf("[WARNING] db.atomic: %d connection retries (last err: %v)\n", i, err)
			}
		}
		conn, err = db.Handle.Conn(ctx)
	}

	// fail case - unable to create database connection
	if err != nil {
		return
	}
	defer conn.Close()

	for i := 0; ; i++ {
		if i > 0 {
			if i < infoTxRetries {
				fmt.Printf("[INFO] db.atomic: %d retries (last err: %v)\n", i, err)
			} else if i >= 1000 {
				fmt.Printf("[FAILURE] db.atomic: %d retries (last err: %v)\n", i, err)
				break
			} else if i%warnTxRetriesInterval == 0 {
				fmt.Printf("[WARNING] db.atomic: %d retries (last err: %v)\n", i, err)
			}
		}

		tx, err = conn.BeginTx(ctx, &sql.TxOptions{Isolation: sql.LevelSerializable, ReadOnly: db.readOnly})
		if dbretry(err) {
			continue
		} else if err != nil {
			break
		}

		// create a transaction context data
		txContextData := &txExecutionContext{
			deadline: atomicDeadline,
		}

		err = guardedFn(context.WithValue(ctx, tx, txContextData), tx)
		if err != nil {
			tx.Rollback()
			if dbretry(err) {
				continue
			} else {
				break
			}
		}

		err = tx.Commit()
		if err == nil {
			// update the deadline, as it might have been updated.
			atomicDeadline = txContextData.deadline
			break
		} else if !dbretry(err) {
			break
		}
	}

	if time.Now().After(atomicDeadline) {
		fmt.Printf("[WARNING] dbatomic: tx surpassed expected deadline by %v\n", time.Now().Sub(atomicDeadline))
	}

	return
}

// SynchronousMode is the synchronous modes supported by sqlite database.
type SynchronousMode int

const (
	// SynchronousModeOff (0), SQLite continues without syncing as soon as it has handed data off to the operating system. If the application running SQLite crashes,
	// the data will be safe, but the database might become corrupted if the operating system crashes or the computer loses power before that data has been written to the
	// disk surface. On the other hand, commits can be orders of magnitude faster with synchronous OFF.
	SynchronousModeOff SynchronousMode = 0
	// SynchronousModeNormal (1), the SQLite database engine will still sync at the most critical moments, but less often than in FULL mode. There is a very small
	// (though non-zero) chance that a power failure at just the wrong time could corrupt the database in journal_mode=DELETE on an older filesystem.
	// WAL mode is safe from corruption with synchronous=NORMAL, and probably DELETE mode is safe too on modern filesystems. WAL mode is always consistent with synchronous=NORMAL,
	// but WAL mode does lose durability. A transaction committed in WAL mode with synchronous=NORMAL might roll back following a power loss or system crash.
	// Transactions are durable across application crashes regardless of the synchronous setting or journal mode.
	// The synchronous=NORMAL setting is a good choice for most applications running in WAL mode.
	SynchronousModeNormal SynchronousMode = 1
	// SynchronousModeFull (2), the SQLite database engine will use the xSync method of the VFS to ensure that all content is safely written to the disk surface prior to continuing.
	// This ensures that an operating system crash or power failure will not corrupt the database. FULL synchronous is very safe, but it is also slower.
	// FULL is the most commonly used synchronous setting when not in WAL mode.
	SynchronousModeFull SynchronousMode = 2
	// SynchronousModeExtra synchronous is like FULL with the addition that the directory containing a rollback journal is synced after that journal is unlinked to commit a
	// transaction in DELETE mode. EXTRA provides additional durability if the commit is followed closely by a power loss.
	SynchronousModeExtra SynchronousMode = 3
)

// SetSynchronousMode updates the synchronous mode of the connection
func (db *Accessor) SetSynchronousMode(ctx context.Context, mode SynchronousMode, fullfsync bool) (err error) {
	if mode < SynchronousModeOff || mode > SynchronousModeExtra {
		return fmt.Errorf("invalid value(%d) was provided to mode", mode)
	}

	_, err = db.Handle.ExecContext(ctx, fmt.Sprintf("PRAGMA synchronous=%d", mode))
	if err != nil {
		return err
	}

	if fullfsync {
		for _, stmt := range enableFullfsyncStatements {
			_, err = db.Handle.ExecContext(ctx, stmt)
			if err != nil {
				break
			}
		}
	} else {
		_, err = db.Handle.ExecContext(ctx, "PRAGMA fullfsync=false")
	}

	return
}

// Append the necessary statements to enableFullfsyncStatements that are
// needed by MacOS to enable fullfsync.
func darwinFullfsyncInit() {
	// See https://www.sqlite.org/pragma.html#pragma_fullfsync
	enableFullfsyncStatements = append(enableFullfsyncStatements, "PRAGMA fullfsync=true")
}
