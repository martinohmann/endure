endure
======

Easily load and store values of arbitrary go types in files or other persistent
storages. The main purpose of this library is to provide an easy-to-use
interface for reading and writing config and state files.

Endure supports marshalling to and unmarshalling from various encodings. It
provides ready-to-use helpers for json, xml and yaml. Other encodings can be
easily supported via pluggable marshal and unmarshal functions.

Installation
------------

```sh
go get -u github.com/martinohmann/endure
```

Usage
-----

Check out the [`examples`](examples) directory for common usage examples.

For an example of how to use custom encoding, have a look at [`custom_test.go`](custom_test.go).

License
-------

The source code of endure is released under the MIT License. See the bundled
LICENSE file for details.
