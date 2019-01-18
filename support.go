package endure

import (
	"encoding/json"
	"encoding/xml"
	"io"

	yaml "gopkg.in/yaml.v2"
)

// NewJSONFileStorage creates a new value of the FileStorage type with
// the json.Marshal and json.Unmarshal funcs.
func NewJSONFileStorage(filename string) Storage {
	return &FileStorage{
		Filename:  filename,
		Marshal:   json.Marshal,
		Unmarshal: json.Unmarshal,
	}
}

// NewXMLFileStorage creates a new value of the FileStorage type with
// the xml.Marshal and xml.Unmarshal funcs.
func NewXMLFileStorage(filename string) Storage {
	return &FileStorage{
		Filename:  filename,
		Marshal:   xml.Marshal,
		Unmarshal: xml.Unmarshal,
	}
}

// NewYAMLFileStorage creates a new value of the FileStorage type with
// the yaml.Marshal and yaml.Unmarshal funcs.
func NewYAMLFileStorage(filename string) Storage {
	return &FileStorage{
		Filename:  filename,
		Marshal:   yaml.Marshal,
		Unmarshal: yaml.Unmarshal,
	}
}

// NewJSONReadWriterStorage creates a new value of the ReadWriterStorage type
// with the json.Marshal and json.Unmarshal funcs.
func NewJSONReadWriterStorage(rw io.ReadWriter) Storage {
	return &ReadWriterStorage{
		RW:        rw,
		Marshal:   json.Marshal,
		Unmarshal: json.Unmarshal,
	}
}

// NewXMLReadWriterStorage creates a new value of the ReadWriterStorage type
// with the xml.Marshal and xml.Unmarshal funcs.
func NewXMLReadWriterStorage(rw io.ReadWriter) Storage {
	return &ReadWriterStorage{
		RW:        rw,
		Marshal:   xml.Marshal,
		Unmarshal: xml.Unmarshal,
	}
}

// NewYAMLReadWriterStorage creates a new value of the ReadWriterStorage type
// with the yaml.Marshal and yaml.Unmarshal funcs.
func NewYAMLReadWriterStorage(rw io.ReadWriter) Storage {
	return &ReadWriterStorage{
		RW:        rw,
		Marshal:   yaml.Marshal,
		Unmarshal: yaml.Unmarshal,
	}
}
