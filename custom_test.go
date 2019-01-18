package endure_test

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"

	"github.com/martinohmann/endure"
)

type someType struct {
	Bar string
	Baz float64
}

func marshalBase64JSON(v interface{}) ([]byte, error) {
	buf, err := json.Marshal(v)
	if err != nil {
		return nil, err
	}

	n := base64.StdEncoding.EncodedLen(len(buf))

	dst := make([]byte, n)

	base64.StdEncoding.Encode(dst, buf)

	return dst, nil
}

func unmarshalBase64JSON(buf []byte, v interface{}) error {
	n := base64.StdEncoding.DecodedLen(len(buf))

	dst := make([]byte, n)

	n, err := base64.StdEncoding.Decode(dst, buf)
	if err != nil {
		return err
	}

	return json.Unmarshal(dst[:n], v)
}

func NewBase64JSONStorage(rw io.ReadWriter) endure.Storage {
	return &endure.ReadWriterStorage{
		RW:        rw,
		Marshal:   marshalBase64JSON,
		Unmarshal: unmarshalBase64JSON,
	}
}

func Example() {
	rw := bytes.NewBuffer(nil)

	storage := NewBase64JSONStorage(rw)

	foo := someType{
		Bar: "qux",
		Baz: 42.0,
	}

	endure.Must(storage.Store(foo))

	fmt.Println(rw.String())

	bar := someType{}

	endure.Must(storage.Load(&bar))

	fmt.Println(bar)

	// Output:
	// eyJCYXIiOiJxdXgiLCJCYXoiOjQyfQ==
	// {qux 42}
}
