package strfs

import (
	"io"
	"strings"
)

// Content represents the content part of a file.
//
// Content does NOT map to anything in Go's built-in "fs" package.
//
// But, for example, is used to create a RegularFile (which maps to a fs.File).
//
// Example usage:
//
//	var content strfs.Content = strfs.CreateContent("<!DOCTYPE html>"+"\n"+"<html></html>")
//	
//	var regularfile strfs.RegularFile = strfs.RegularFile{
//		FileContent: content,
//		FileName:    "notice.html",
//		FileModTIme: time.Date(2022, 12, 12, 10, 30, 14, 2, time.UTC),
//	}
type Content struct{
	reader io.Reader
	size int64
	closed bool
}

// A trick to make sure strfs.Content fits the io.ReadCloser interface.
// This is a compile-time check.
var _ io.ReadCloser = &Content{}

// CreateContent returns a strfs.Content whose content is the string given to it.
//
// Example usage:
//
//	var content strfs.Content = strfs.CreateContent("# Hello world!"+"\r\r"+"Welcome to my document."+"\n")
//	
//	var regularfile strfs.RegularFile = strfs.RegularFile{
//		FileContent: content,
//		FileName:    "message.md",
//		FileModTIme: time.Now(),
//	}
func CreateContent(s string) Content {
	var reader io.Reader = strings.NewReader(s)
	var size int64 = int64(len(s))

	return Content{
		reader:reader,
		size:size,
	}
}

// EmptyContent is used to see if a given strfs.Content is empty.
//
// Note that a strfs.Content being empty is NOT the same as containing the empty string!
//
// A strfs.Content is empty when it hasn't been initialized.
//
// Example usage:
//
//	var content strfs.Content
//	
//	// ...
//	
//	if strfs.EmptyContent() == content {
//		//@TODO
//	}
func EmptyContent() Content {
	return Content{}
}

// Close makes strfs.Content fit the io.Closer interface.
//
// Call close will stop the Read method from working.
//
// Close can safely be called more than once.
func (receiver *Content) Close() error {
	if nil == receiver {
		return errNilReceiver
	}

	if receiver.closed {
		return nil
	}

	receiver.closed = true
	return nil
}

// Closed returns whether a strfs.Content is closed or not.
func (receiver *Content) Closed() bool {
	if nil == receiver {
		return true
	}

	return receiver.closed
}
// Read makes strfs.Content fit the io.Reader interface.
//
// Read reads up to len(p) bytes into 'p'.
// Read returns the number of bytes actually read, and any errors it encountered.
//
// Example usage:
//
//	var content strfs.Content = strfs.CreateContent("ABCDEFGHIJKLMNOPQRSTUVWXYZ")
//	
//	var b1 [5]byte
//	var p1 []byte = b1[:]
//	
//	n, err := content.Read(p1)
//	
//	// n == 5
//	// b1 == [5]byte{'A', 'B', 'C', 'D', 'E'}
//	
//	// ...
//	
//	var b2 [4]byte
//	var p2 []byte = b2[:]
//	
//	n, err := content.Read(p2)
//	
//	// n == 4
//	// b2 == [5]byte{'F', 'G', 'H', 'I'}
func (receiver *Content) Read(p []byte) (int, error) {
	if nil == receiver {
		return 0, errNilReceiver
	}
	if nil == p {
		return 0, errNilByteSlice
	}

	var reader io.Reader = receiver.reader
	if nil == reader {
		return 0, errInternalError
	}

	if receiver.closed {
		return 0, nil
	}

	return receiver.reader.Read(p)
}

// Size returns the of the strnig given to it as the number of bytes.
func (receiver *Content) Size() int64 {
	if nil == receiver {
		return 0
	}

	return receiver.size
}
