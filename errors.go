package strfs

import (
	"github.com/reiver/go-fck"
)

const (
	errClosed        = fck.Error("closed")
	errInternalError = fck.Error("internal error")
	errNilByteSlice  = fck.Error("nil byte slice")
	errNilReader     = fck.Error("nil reader")
	errNilReceiver   = fck.Error("nil receiver")
)
