// Package endure provides an easy to use interface for loading and storing
// arbitrary values into files and io.ReadWriters. Functions for marshalling to
// and unmarshalling from a target encoding can be overridden at runtime.
package endure

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"os"
	"sync"
)

// Storage defines the interface for a storage provider.
type Storage interface {
	// Load loads data from storage into v. Returns an error if loading fails.
	Load(v interface{}) error

	// Store saves v to the storage. Returns an error if saving fails.
	Store(v interface{}) error
}

// MarshalFunc defines the signature of a function for marshalling a value v
// into a byte slice.
type MarshalFunc func(v interface{}) ([]byte, error)

// UnmarshalFunc defines the signature of a function for unmarshalling a byte
// slice into a value v.
type UnmarshalFunc func(buf []byte, v interface{}) error

var (
	// Marshal defines the default function for marshalling an object. Can be
	// overridden at runtime.
	Marshal MarshalFunc = json.Marshal

	// Unmarshal defines the default function for unmarshalling an object. Can
	// be overridden at runtime.
	Unmarshal UnmarshalFunc = json.Unmarshal

	// mu is a global mutex for synchronizing calls to Load(),
	// LoadWithReader(), Store() and StoreWithWriter()
	mu sync.Mutex
)

// Must panics if the given error is not nil. Can be wrapped around Load() and
// Store() to panic on errors.
func Must(err error) {
	if err != nil {
		panic(err)
	}
}

// Load loads the contents of filename and tries to unmarshal it into v using
// the default Unmarshal func. Will return an error if opening the file or
// unmarshalling fails.
func Load(filename string, v interface{}) error {
	return load(filename, v, Unmarshal, &mu)
}

// LoadWithReader reads the contents of r and tries to unmarshal it into v
// using the default Unmarshal func. Will return an error if reading r or
// unmarshalling fails.
func LoadWithReader(r io.Reader, v interface{}) error {
	return loadWithReader(r, v, Unmarshal, &mu)
}

// Store marshalls v into a byte slice using the default Marshal func and saves
// it to filename. Will return an error if any file operation or the
// marshalling fails.
func Store(filename string, v interface{}) error {
	return store(filename, v, Marshal, &mu)
}

// StoreWithWriter marshalls v into a byte slice using the default Marshal func
// and write it to w. Will return an error if writing w or the marshalling
// fails.
func StoreWithWriter(w io.Writer, v interface{}) error {
	return storeWithWriter(w, v, Marshal, &mu)
}

// load loads the contents of filename and tries to unmarshal it into v using
// the provided unmarshal func. Will return an error if opening the file or
// unmarshalling fails.
func load(filename string, v interface{}, unmarshal UnmarshalFunc, mu *sync.Mutex) error {
	f, err := os.Open(filename)
	if err != nil {
		return err
	}

	defer f.Close()

	return loadWithReader(f, v, unmarshal, mu)
}

// loadWithReader reads the contents of r and tries to unmarshal it into v
// using the provided Unmarshal func. Will return an error if reading r or
// unmarshalling fails.
func loadWithReader(r io.Reader, v interface{}, unmarshal UnmarshalFunc, mu *sync.Mutex) error {
	mu.Lock()
	defer mu.Unlock()

	buf, err := ioutil.ReadAll(r)
	if err != nil {
		return err
	}

	return unmarshal(buf, v)
}

// store marshalls v into a byte slice using the provided Marshal func and saves
// it to filename. Will return an error if any file operation or the
// marshalling fails.
func store(filename string, v interface{}, marshal MarshalFunc, mu *sync.Mutex) error {
	f, err := os.Create(filename)
	if err != nil {
		return err
	}

	defer f.Close()

	return storeWithWriter(f, v, marshal, mu)
}

// storeWithWriter marshalls v into a byte slice using the provided Marshal func
// and write it to w. Will return an error if writing w or the marshalling
// fails.
func storeWithWriter(w io.Writer, v interface{}, marshal MarshalFunc, mu *sync.Mutex) error {
	buf, err := marshal(v)
	if err != nil {
		return err
	}

	mu.Lock()
	defer mu.Unlock()

	_, err = w.Write(buf)

	return err
}
