package data

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/Cameronketchem/CEN3031-Group91/server/utils/db"
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

func TestOpenStore(t *testing.T) {
	fn := fmt.Sprintf("/tmp/%s.%d.sqlite3", t.Name(), rand.Intn(100))
	str, err := OpenStore(fn, false)
	defer cleanupSqliteDb(t, fn)
	defer str.Close()

	require.NoError(t, err)
}

func TestOpenStoreErasable(t *testing.T) {
	fn := fmt.Sprintf("/tmp/%s.%d.sqlite3", t.Name(), rand.Intn(100))
	str, err := OpenStoreErasable(fn, db.SynchronousModeExtra)
	defer cleanupSqliteDb(t, fn)
	defer str.Close()

	require.NoError(t, err)
}

func contains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}

	return false
}

func TestInit(t *testing.T) {
	fn := fmt.Sprintf("/tmp/%s.%d.sqlite3", t.Name(), rand.Intn(100))
	str, err := OpenStore(fn, false)
	defer cleanupSqliteDb(t, fn)
	defer str.Close()
	require.NoError(t, err)

	err = str.Init(db.SynchronousModeExtra, false)
	require.NoError(t, err)

	err = str.Handler.Rdb.Atomic(func(ctx context.Context, tx *sql.Tx) error {
		tables := []string{}

		// Assets table
		rows, err := tx.Query("SELECT name FROM sqlite_master")
		defer rows.Close()
		require.NoError(t, err)

		for rows.Next() {
			var exists string
			rows.Scan(&exists)
			tables = append(tables, exists)
		}

		require.Equal(t, true, contains(tables, "users"))
		require.Equal(t, true, contains(tables, "contracts"))
		require.Equal(t, true, contains(tables, "assets"))
		require.Equal(t, true, contains(tables, "contributions"))
		return err
	})

	require.NoError(t, err)
}

func TestDropAll(t *testing.T) {
	fn := fmt.Sprintf("/tmp/%s.%d.sqlite3", t.Name(), rand.Intn(100))
	str, err := OpenStore(fn, false)
	defer cleanupSqliteDb(t, fn)
	defer str.Close()
	require.NoError(t, err)

	err = str.Init(db.SynchronousModeExtra, false)
	require.NoError(t, err)

	// Drop all tables
	str.DropAll()

	err = str.Handler.Rdb.Atomic(func(ctx context.Context, tx *sql.Tx) error {
		tables := []string{}

		// Assets table
		rows, err := tx.Query("SELECT name FROM sqlite_master")
		defer rows.Close()
		require.NoError(t, err)

		for rows.Next() {
			var exists string
			rows.Scan(&exists)
			if exists == "sqlite_sequence" {
				continue
			}
			tables = append(tables, exists)
		}

		require.Equal(t, 0, len(tables))
		return err
	})

	require.NoError(t, err)
}

func TestSeed(t *testing.T) {
	fn := fmt.Sprintf("/tmp/%s.%d.sqlite3", t.Name(), rand.Intn(100))
	str, err := OpenStore(fn, false)
	defer cleanupSqliteDb(t, fn)
	defer str.Close()
	require.NoError(t, err)

	err = str.Init(db.SynchronousModeExtra, false)
	require.NoError(t, err)

	// Seed database
	err = str.Seed()
	require.NoError(t, err)
}
