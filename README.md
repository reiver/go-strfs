# go-strfs

Package **strfs** provides a virtual file-system, whre a `fs.File` can be created from a Go `string`.

## Documention

Online documentation, which includes examples, can be found at: http://godoc.org/sourcecode.social/reiver/go-strfs

[![GoDoc](https://godoc.org/sourcecode.social/reiver/go-strfs?status.svg)](https://godoc.org/sourcecode.social/reiver/go-strfs)

## Example fs.File

Here is an example of turning a Go `string` into a `fs.File`:

```go
import "sourcecode.social/reiver/go-strfs"

// ...

var s string = "<!DOCTYPE html>"+"\n"+"<html><body>Hello world!</body></html>"

var content strfs.Content = strfs.CreateContent(s)

var regularfile strfs.RegularFile = strfs.RegularFile{
	FileContent: content,
	FileName:    "helloworld.html",
	FileModTIme: time.Date(2022, 12, 12, 10, 30, 14, 2, time.UTC),
}

var file fs.FS = &regularfile

```
