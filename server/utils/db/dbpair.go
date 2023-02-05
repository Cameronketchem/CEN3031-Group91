package db

import "context"

// Pair represents two accessors - read and write
type Pair struct {
	Rdb Accessor
	Wdb Accessor
}

// Close the read and write accessors
func (p Pair) Close() {
	if p.Rdb.Handle != nil {
		p.Rdb.Close()
	}
	if p.Wdb.Handle != nil {
		p.Wdb.Close()
	}
}

// OpenPair opens the filename with both reading and writing accessors.
func OpenPair(filename string, memory bool) (p Pair, err error) {
	p.Rdb, err = MakeAccessor(filename, true, memory)
	if err != nil {
		return
	}

	p.Wdb, err = MakeAccessor(filename, false, memory)
	if err != nil {
		p.Rdb.Close()
		return
	}

	return
}

// OpenErasablePair opens the filename with both reading and writing accessors
// with the secure_delete pragma set, using MakeErasableAccessor.
func OpenErasablePair(filename string) (p Pair, err error) {
	p.Rdb, err = makeErasableAccessor(filename, true)
	if err != nil {
		return
	}

	p.Wdb, err = makeErasableAccessor(filename, false)
	if err != nil {
		p.Rdb.Close()
		return
	}

	return
}

// Set the synchronous mode for both the read and write accessors.
func (p Pair) SetSynchronousMode(ctx context.Context, sync SynchronousMode, fullfsync bool) error {
	err := p.Rdb.SetSynchronousMode(ctx, sync, fullfsync)
	if err != nil {
		return err
	}

	err = p.Wdb.SetSynchronousMode(ctx, sync, fullfsync)
	if err != nil {
		return err
	}

	return nil
}
