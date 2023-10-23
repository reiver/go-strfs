package strfs

import (
	"io/fs"
	"time"
)

// RegularFile lets you turn a string into a fs.File.
//
// Example usage:
//
//	var content strfs.Content = strfs.CreateContent("<!DOCTYPE html>"+"\n"+"<html><body>Hello world!</body></html>")
//	
//	var regularfile strfs.RegularFile = strfs.RegularFile{
//		FileContent: content,
//		FileName:    "helloworld.html",
//		FileModTime: time.Date(2022, 12, 12, 10, 30, 14, 2, time.UTC),
//	}
type RegularFile struct {
	FileContent Content
	FileName string
	FileModTime time.Time
}

// A trick to make sure strfs.RegularFile fits the fs.File interface.
// This is a compile-time check.
var _ fs.File = &RegularFile{}

// Close will stop the Read method from working.
//
// Close can safely be called more than once.
//
// Close helps strfs.RegularFile fit the fs.File interface.
// Close makes strfs.RegularFile fit the io.Closer interface.
func (receiver *RegularFile) Close() error {
	if nil == receiver {
		return errNilReceiver
	}

	return receiver.FileContent.Close()
}

// Read reads up to len(p) bytes into 'p'.
// Read returns the number of bytes actually read, and any errors it encountered.
//
// Read helps strfs.RegularFile fit the fs.File interface.
// Read makes strfs.Content fit the io.Reader interface.
//
// Example usage:
//
//	var content strfs.Content = strfs.CreateContent("ABCDEFGHIJKLMNOPQRSTUVWXYZ")
//	
//	var regularfile strfs.RegularFile = strfs.RegularFile{
//		FileContent: content,
//		FileName: "alphabet.txt",
//		FileModTime: time.Now(),
//	}
//	
//	var b1 [5]byte
//	var p1 []byte = b1[:]
//	
//	n, err := regularfile.Read(p1)
//	
//	// n == 5
//	// b1 == [5]byte{'A', 'B', 'C', 'D', 'E'}
//	
//	// ...
//	
//	var b2 [4]byte
//	var p2 []byte = b2[:]
//	
//	n, err := regularfile.Read(p2)
//	
//	// n == 4
//	// b2 == [5]byte{'F', 'G', 'H', 'I'}
func (receiver *RegularFile) Read(p []byte) (int, error) {
	if nil == receiver {
		return 0, errNilReceiver
	}

	return receiver.FileContent.Read(p)
}

// Stat returns a fs.FileInfo for a *strfs.RegularFile.
//
// Stat helps strfs.RegularFile fit the fs.File interface.
func (receiver *RegularFile) Stat() (fs.FileInfo, error) {
	if nil == receiver {
		return nil, errNilReceiver
	}

	if EmptyContent() == receiver.FileContent {
		return nil, errEmptyContent
	}

	const modeRegularFile = 0

	return internalFileInfo{
		sys:     receiver.FileContent.String(),
		name:    receiver.FileName,
		size:    receiver.FileContent.Size(),
		mode:    modeRegularFile,
		modtime: receiver.FileModTime,
	}, nil
}
