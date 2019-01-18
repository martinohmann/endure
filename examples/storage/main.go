package main

import (
	"flag"
	"fmt"
	"os"

	yaml "gopkg.in/yaml.v2"

	"github.com/martinohmann/endure"
)

type someType struct {
	Foo string  `xml:"foo"`
	Bar float64 `xml:"bar"`
}

var (
	foo      string
	bar      float64
	filename string

	usage = func() {
		fmt.Fprintf(os.Stderr, "usage: %s\n", os.Args[0])
		flag.PrintDefaults()
	}
)

func init() {
	// change the global marshaller to yaml
	endure.Marshal = yaml.Marshal
	endure.Unmarshal = yaml.Unmarshal

	flag.Usage = usage
	flag.StringVar(&foo, "foo", "", "foo")
	flag.Float64Var(&bar, "bar", 0, "bar")
	flag.StringVar(&filename, "filename", "/tmp/endure-example.xml", "filename")
}

func main() {
	flag.Parse()

	storage := endure.NewXMLFileStorage(filename)

	var v someType

	err := storage.Load(&v)
	if err != nil && !os.IsNotExist(err) {
		panic(err)
	}

	fmt.Println(v)

	v.Foo = foo
	v.Bar = bar

	endure.Must(storage.Store(v))
}
