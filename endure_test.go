package endure

import (
	"bytes"
	"errors"
	"io/ioutil"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

type testData struct {
	Foo string
}

func TestMust(t *testing.T) {
	assert.Panics(t, func() { Must(errors.New("yikes")) })
}

func TestLoad(t *testing.T) {
	val := &testData{}

	err := Load("testdata/foo.golden", val)

	if assert.NoError(t, err) {
		assert.Equal(t, "BAR", val.Foo)
	}
}

func TestLoadNonexistent(t *testing.T) {
	val := &testData{}

	err := Load("testdata/nonexistent", val)

	assert.Error(t, err)
	assert.True(t, os.IsNotExist(err))
}

func TestLoadWithReader(t *testing.T) {
	buf := bytes.NewBuffer([]byte(`{"foo": "bar"}`))

	val := &testData{}

	err := LoadWithReader(buf, val)

	if assert.NoError(t, err) {
		assert.Equal(t, "bar", val.Foo)
	}
}

type badReadWriter struct{}

func (rw *badReadWriter) Write(p []byte) (n int, err error) {
	return 0, errors.New("yikes")
}

func (rw *badReadWriter) Read(p []byte) (n int, err error) {
	return 0, errors.New("yikes")
}

func TestLoadWithBadReader(t *testing.T) {
	assert.Error(t, LoadWithReader(&badReadWriter{}, &testData{}))
}

func TestStore(t *testing.T) {
	f, err := ioutil.TempFile("", "endure-test")
	if err != nil {
		t.Fatal(err)
	}

	f.Close()

	filename := f.Name()

	defer os.Remove(filename)

	err = Store(filename, &testData{Foo: "baz"})

	if assert.NoError(t, err) {
		contents, err := ioutil.ReadFile(filename)

		if assert.NoError(t, err) {
			assert.Equal(t, `{"Foo":"baz"}`, string(contents))
		}
	}
}

func TestStoreWithWriter(t *testing.T) {
	buf := bytes.NewBuffer(nil)

	err := StoreWithWriter(buf, &testData{Foo: "baz"})

	if assert.NoError(t, err) {
		assert.Equal(t, `{"Foo":"baz"}`, buf.String())
	}
}

func TestStoreWithBadWriter(t *testing.T) {
	assert.Error(t, StoreWithWriter(&badReadWriter{}, &testData{}))
}
