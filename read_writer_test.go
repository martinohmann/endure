package endure_test

import (
	"bytes"
	"fmt"

	"github.com/martinohmann/endure"
)

func ExampleReadWriterStorage_Load() {
	rw := bytes.NewBuffer([]byte(`"foo"`))

	storage := endure.NewReadWriterStorage(rw)

	var s string

	endure.Must(storage.Load(&s))

	fmt.Println(s)

	// Output:
	// foo
}

func ExampleReadWriterStorage_Store() {
	rw := bytes.NewBuffer(nil)

	storage := endure.NewReadWriterStorage(rw)

	s := "foo"

	endure.Must(storage.Store(s))

	fmt.Println(rw.String())

	// Output:
	// "foo"
}
