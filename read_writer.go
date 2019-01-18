package endure

import (
	"io"
	"sync"
)

// ReadWriter defines a storage that reads from and writes to an io.ReadWriter.
type ReadWriterStorage struct {
	mu        sync.Mutex
	RW        io.ReadWriter
	Marshal   MarshalFunc
	Unmarshal UnmarshalFunc
}

// NewReadWriterStorage creates a new value of the ReadWriterStorage type with
// the default Marshal and Unmarshal funcs.
func NewReadWriterStorage(rw io.ReadWriter) Storage {
	return &ReadWriterStorage{
		RW:        rw,
		Marshal:   Marshal,
		Unmarshal: Unmarshal,
	}
}

// Load loads data from storage into v. Returns an error if loading fails.
func (s *ReadWriterStorage) Load(v interface{}) error {
	return loadWithReader(s.RW, v, s.Unmarshal, &s.mu)
}

// Store saves v to the storage. Returns an error if saving fails.
func (s *ReadWriterStorage) Store(v interface{}) error {
	return storeWithWriter(s.RW, v, s.Marshal, &s.mu)
}
