package strfs

import (
	"github.com/reiver/go-erorr"
)

const (
	errClosed        = erorr.Error("closed")
	errEmptyContent  = erorr.Error("empty content")
	errInternalError = erorr.Error("internal error")
	errNilByteSlice  = erorr.Error("nil byte slice")
	errNilReadSeeker = erorr.Error("nil read-seeker")
	errNilReceiver   = erorr.Error("nil receiver")
)
