package endure_test

import (
	"bytes"
	"fmt"

	"github.com/martinohmann/endure"
)

func ExampleLoad() {
	var m map[string]string

	endure.Must(endure.Load("testdata/foo.golden", &m))

	fmt.Println(m)

	// Output:
	// map[foo:BAR]
}

func ExampleLoadWithReader() {
	buf := bytes.NewBuffer([]byte("42.5"))

	var v float64

	endure.Must(endure.LoadWithReader(buf, &v))

	fmt.Println(v)

	// Output:
	// 42.5
}

func ExampleStoreWithWriter() {
	buf := bytes.NewBuffer(nil)

	v := "foo"

	endure.Must(endure.StoreWithWriter(buf, v))

	fmt.Println(buf.String())

	// Output:
	// "foo"
}
