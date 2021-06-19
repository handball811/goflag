package goflag

import "errors"

var (
	ErrFlagSize   = errors.New("flag size you specified does not supported")
	ErrOutOfBound = errors.New("index you specified is out of range")
	ErrNoSpace    = errors.New("Cannot find down flag")
)
