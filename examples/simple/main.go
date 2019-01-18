package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/martinohmann/endure"
)

type someType struct {
	Foo string  `json:"foo"`
	Bar float64 `json:"bar"`
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
	flag.Usage = usage
	flag.StringVar(&foo, "foo", "", "foo")
	flag.Float64Var(&bar, "bar", 0, "bar")
	flag.StringVar(&filename, "filename", "/tmp/endure-example.json", "filename")
}

func main() {
	flag.Parse()

	var v someType

	err := endure.Load(filename, &v)
	if err != nil && !os.IsNotExist(err) {
		panic(err)
	}

	fmt.Println(v)

	v.Foo = foo
	v.Bar = bar

	endure.Must(endure.Store(filename, v))
}
