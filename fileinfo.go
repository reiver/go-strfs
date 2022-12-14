package strfs

import (
	"io/fs"
	"time"
)

type internalFileInfo struct {
	sys string
	mode fs.FileMode
	modtime time.Time
	name string
	size int64
}

var _ fs.FileInfo = &internalFileInfo{}

func (receiver internalFileInfo) IsDir() bool {
	return 0 != (receiver.mode & fs.ModeDir)
}

func (receiver internalFileInfo) Mode() fs.FileMode {
	return receiver.mode
}

func (receiver internalFileInfo) ModTime() time.Time {
	return receiver.modtime
}

func (receiver internalFileInfo) Name() string {
	return receiver.name
}

func (receiver internalFileInfo) Size() int64 {
	return receiver.size
}

func (receiver internalFileInfo) Sys() any {
	return receiver.sys
}
