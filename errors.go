package strfs

import (
	"sourcecode.social/reiver/go-erorr"
)

const (
	errClosed        = erorr.Error("closed")
	errEmptyContent  = erorr.Error("empty content")
	errInternalError = erorr.Error("internal error")
	errNilByteSlice  = erorr.Error("nil byte slice")
	errNilReader     = erorr.Error("nil reader")
	errNilReceiver   = erorr.Error("nil receiver")
)
