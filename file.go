package endure

import (
	"sync"
)

// FileStorage defines a storage that loads from and saves to a file.
type FileStorage struct {
	mu        sync.Mutex
	Filename  string
	Marshal   MarshalFunc
	Unmarshal UnmarshalFunc
}

// NewFileStorage creates a new value of the FileStorage type with
// the default Marshal and Unmarshal funcs.
func NewFileStorage(filename string) Storage {
	return &FileStorage{
		Filename:  filename,
		Marshal:   Marshal,
		Unmarshal: Unmarshal,
	}
}

// Load loads data from storage into v. Returns an error if loading fails.
func (s *FileStorage) Load(v interface{}) error {
	return load(s.Filename, v, s.Unmarshal, &s.mu)
}

// Store saves v to the storage. Returns an error if saving fails.
func (s *FileStorage) Store(v interface{}) error {
	return store(s.Filename, v, s.Marshal, &s.mu)
}
