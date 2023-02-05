package db

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/stretchr/testify/require"
	"math/rand"
	"os"
	"path/filepath"
	"testing"
)

// Delete a test database file
func cleanupSqliteDb(t *testing.T, path string) {
	parts, err := filepath.Glob(path + "*")
	if err != nil {
		t.Errorf("%s*: could not glob, %s", path, err)
		return
	}
	for _, part := range parts {
		err = os.Remove(part)
		if err != nil {
			t.Errorf("%s: error cleaning up, %s", part, err)
		}
	}
}

func TestInMemoryDisposal(t *testing.T) {
	acc, err := MakeAccessor("fn.db", false, true)
	require.NoError(t, err)
	err = acc.Atomic(func(ctx context.Context, tx *sql.Tx) error {
		_, err := tx.Exec("create table Service (data blob)")
		return err
	})
	require.NoError(t, err)

	err = acc.Atomic(func(ctx context.Context, tx *sql.Tx) error {
		raw := []byte{0, 1, 2}
		_, err := tx.Exec("insert or replace into Service (rowid, data) values (1, ?)", raw)
		return err
	})
	require.NoError(t, err)

	anotherAcc, err := MakeAccessor("fn.db", false, true)
	require.NoError(t, err)
	err = anotherAcc.Atomic(func(ctx context.Context, tx *sql.Tx) error {
		var nrows int
		row := tx.QueryRow("select count(*) from Service")
		err := row.Scan(&nrows)
		return err
	})
	require.NoError(t, err)
	anotherAcc.Close()

	acc.Close()

	acc, err = MakeAccessor("fn.db", false, true)
	require.NoError(t, err)
	err = acc.Atomic(func(ctx context.Context, tx *sql.Tx) error {
		var nrows int
		row := tx.QueryRow("select count(*) from Service")
		err := row.Scan(&nrows)
		if err == nil {
			return errors.New("table `Service` presents while it should not")
		}
		return nil
	})
	require.NoError(t, err)

	acc.Close()
}

func TestInMemoryUniqueDB(t *testing.T) {
	acc, err := MakeAccessor("fn.db", false, true)
	require.NoError(t, err)
	defer acc.Close()
	err = acc.Atomic(func(ctx context.Context, tx *sql.Tx) error {
		_, err := tx.Exec("create table Service (data blob)")
		return err
	})
	require.NoError(t, err)

	err = acc.Atomic(func(ctx context.Context, tx *sql.Tx) error {
		raw := []byte{0, 1, 2}
		_, err := tx.Exec("insert or replace into Service (rowid, data) values (1, ?)", raw)
		return err
	})
	require.NoError(t, err)

	anotherAcc, err := MakeAccessor("fn2.db", false, true)
	require.NoError(t, err)
	defer anotherAcc.Close()
	err = anotherAcc.Atomic(func(ctx context.Context, tx *sql.Tx) error {
		var nrows int
		row := tx.QueryRow("select count(*) from Service")
		err := row.Scan(&nrows)
		if err == nil {
			return errors.New("table `Service` presents while it should not")
		}
		return nil
	})
	require.NoError(t, err)
}

func TestDBConcurrency(t *testing.T) {
	fn := fmt.Sprintf("/tmp/%s.%d.sqlite3", t.Name(), rand.Intn(100))
	defer cleanupSqliteDb(t, fn)

	acc, err := MakeAccessor(fn, false, false)
	require.NoError(t, err)
	defer acc.Close()

	acc2, err := MakeAccessor(fn, true, false)
	require.NoError(t, err)
	defer acc2.Close()

	err = acc.Atomic(func(ctx context.Context, tx *sql.Tx) error {
		_, err := tx.Exec("CREATE TABLE foo (a INTEGER, b INTEGER)")
		return err
	})
	require.NoError(t, err)

	err = acc.Atomic(func(ctx context.Context, tx *sql.Tx) error {
		_, err := tx.Exec("INSERT INTO foo (a, b) VALUES (?, ?)", 1, 1)
		return err
	})
	require.NoError(t, err)

	c1 := make(chan struct{})
	c2 := make(chan struct{})
	go func() {
		err := acc.Atomic(func(ctx context.Context, tx *sql.Tx) error {
			<-c2

			_, err := tx.Exec("INSERT INTO foo (a, b) VALUES (?, ?)", 2, 2)
			if err != nil {
				return err
			}

			c1 <- struct{}{}
			<-c2

			_, err = tx.Exec("INSERT INTO foo (a, b) VALUES (?, ?)", 3, 3)
			if err != nil {
				return err
			}

			return nil
		})

		require.NoError(t, err)
		c1 <- struct{}{}
	}()

	err = acc2.Atomic(func(ctx context.Context, tx *sql.Tx) error {
		var nrows int64
		err := tx.QueryRow("SELECT COUNT(*) FROM foo").Scan(&nrows)
		if err != nil {
			return err
		}

		if nrows != 1 {
			return fmt.Errorf("row count mismatch: %d != 1", nrows)
		}

		return nil
	})
	require.NoError(t, err)

	c2 <- struct{}{}
	<-c1

	err = acc2.Atomic(func(ctx context.Context, tx *sql.Tx) error {
		var nrows int64
		err := tx.QueryRow("SELECT COUNT(*) FROM foo").Scan(&nrows)
		if err != nil {
			return err
		}

		if nrows != 1 {
			return fmt.Errorf("row count mismatch: %d != 1", nrows)
		}

		return nil
	})
	require.NoError(t, err)

	c2 <- struct{}{}
	<-c1

	err = acc2.Atomic(func(ctx context.Context, tx *sql.Tx) error {
		var nrows int64
		err := tx.QueryRow("SELECT COUNT(*) FROM foo").Scan(&nrows)
		if err != nil {
			return err
		}

		if nrows != 3 {
			return fmt.Errorf("row count mismatch: %d != 3", nrows)
		}

		return nil
	})
	require.NoError(t, err)
}

func TestSetSynchronousMode(t *testing.T) {
	setSynchrounousModeHelper := func(mem bool, ctx context.Context, mode SynchronousMode, fullfsync bool) error {
		acc, err := MakeAccessor("fn.db", false, mem)
		require.NoError(t, err)
		if !mem {
			defer os.Remove("fn.db")
			defer os.Remove("fn.db-shm")
			defer os.Remove("fn.db-wal")
		}
		defer acc.Close()
		return acc.SetSynchronousMode(ctx, mode, fullfsync)
	}
	// check with canceled context.
	ctx, cancelFunc := context.WithCancel(context.Background())
	cancelFunc()

	require.Error(t, context.Canceled, setSynchrounousModeHelper(true, ctx, SynchronousModeOff, false))
	require.Error(t, context.Canceled, setSynchrounousModeHelper(false, ctx, SynchronousModeOff, false))

	require.Contains(t, setSynchrounousModeHelper(false, context.Background(), SynchronousModeOff-1, false).Error(), "invalid value")
	require.Contains(t, setSynchrounousModeHelper(false, context.Background(), SynchronousModeExtra+1, false).Error(), "invalid value")

	// try all success permutations -
	for _, mode := range []SynchronousMode{SynchronousModeOff, SynchronousModeNormal, SynchronousModeFull, SynchronousModeExtra} {
		for _, disk := range []bool{true, false} {
			for _, fullfsync := range []bool{true, false} {
				require.NoError(t, setSynchrounousModeHelper(disk, context.Background(), mode, fullfsync))
			}
		}
	}
}

func TestReadingWhileWriting(t *testing.T) {
	writeAcc, err := MakeAccessor("fn.db", false, false)
	require.NoError(t, err)
	defer os.Remove("fn.db")
	defer os.Remove("fn.db-shm")
	defer os.Remove("fn.db-wal")
	defer writeAcc.Close()
	readAcc, err := MakeAccessor("fn.db", true, false)
	require.NoError(t, err)
	defer readAcc.Close()
	_, err = writeAcc.Handle.Exec("CREATE TABLE foo (a INTEGER, b INTEGER)")
	require.NoError(t, err)

	_, err = writeAcc.Handle.Exec("INSERT INTO foo(a,b) VALUES (1,1)")
	require.NoError(t, err)

	var count int
	err = readAcc.Handle.QueryRow("SELECT count(*) FROM foo").Scan(&count)
	require.NoError(t, err)
	require.Equal(t, 1, count)

	err = writeAcc.Atomic(func(ctx context.Context, tx *sql.Tx) error {
		_, err = tx.Exec("INSERT INTO foo(a,b) VALUES (2,2)")
		if err != nil {
			return err
		}
		err = readAcc.Handle.QueryRow("SELECT count(*) FROM foo").Scan(&count)
		if err != nil {
			return err
		}
		return nil
	})
	require.NoError(t, err)
	// this should be 1, since it was queried before the commit.
	require.Equal(t, 1, count)
	err = readAcc.Handle.QueryRow("SELECT count(*) FROM foo").Scan(&count)
	require.NoError(t, err)
	require.Equal(t, 2, count)
}
