package endure_test

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/martinohmann/endure"
)

func ExampleFileStorage_Load() {
	f, err := ioutil.TempFile("", "endure-test")
	if err != nil {
		panic(err)
	}

	filename := f.Name()

	defer os.Remove(filename)

	f.WriteString(`42.5`)
	f.Close()

	storage := endure.NewFileStorage(filename)

	var v float64

	endure.Must(storage.Load(&v))

	fmt.Println(v)

	// Output:
	// 42.5
}

func ExampleFileStorage_Store() {
	f, err := ioutil.TempFile("", "endure-test")
	if err != nil {
		panic(err)
	}

	filename := f.Name()

	defer func() {
		f.Close()
		os.Remove(filename)
	}()

	storage := endure.NewFileStorage(filename)

	v := 42.5

	endure.Must(storage.Store(v))

	buf, err := ioutil.ReadAll(f)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(buf))

	// Output:
	// 42.5
}
