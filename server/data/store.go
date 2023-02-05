package data

import (
	"context"
	"fmt"
	"github.com/Cameronketchem/CEN3031-Group91/server/utils/db"
)

// A Store is a database driver
type Store struct {
	Handler   db.Pair
	ctx       context.Context
	sync      db.SynchronousMode
	fullfsync bool
	init      bool
}

// Create a new datastore from the specified filename.
// Memory (true) indicates whether to open the file in-memory.
func OpenStore(filename string, memory bool) (Store, error) {
	handler, err := db.OpenPair(filename, memory)
	if err != nil {
		return Store{}, err
	}

	var str Store
	str.Handler = handler
	return str, nil
}

// Create a new datastore from the specified filename.
// Opens the datastore handler with the secure_delete pragma
// set.
func OpenStoreErasable(filename string, sync db.SynchronousMode) (Store, error) {
	handler, err := db.OpenErasablePair(filename)
	if err != nil {
		return Store{}, err
	}

	var str Store
	str.Handler = handler
	return str, nil
}

// Close the database connection
func (str Store) Close() {
	str.Handler.Close()
}

func (str *Store) Init(sync db.SynchronousMode, fullfsync bool) error {
	str.ctx = context.Background()
	str.sync = sync
	str.fullfsync = fullfsync

	err := str.Handler.SetSynchronousMode(str.ctx, str.sync, fullfsync)
	if err != nil {
		return err
	}

	// Initialize tables (creates tables if they don't already exist)
	if err = str.Handler.Wdb.Atomic(initTables); err != nil {
		return err
	}

	str.init = true
	return nil
}

// Drop all tables and erase all data. To use the db after dropping all tables,
// the store will need to Close() and reinitialize  - use Init().
func (str *Store) DropAll() error {
	if !str.init {
		return fmt.Errorf("str.DropAll: datastore is uninitialized")
	}

	if err := str.Handler.Wdb.Atomic(dropTables); err != nil {
		return err
	}

	return nil
}

// Seed the database. Drops all tables!
func (str *Store) Seed() error {
	if !str.init {
		return fmt.Errorf("str.Seed: datastore is uninitialized")
	}

	str.DropAll()
	if err := str.Handler.Wdb.Atomic(initTables); err != nil {
		return err
	}

	if err := str.Handler.Wdb.Atomic(demoSeed); err != nil {
		return err
	}

	return nil
}
