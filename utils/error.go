package utils

import "github.com/pkg/errors"

var (
	InvalidSlice       = errors.New("slice type error")
	SliceIndexRangeOut = errors.New("slice index range out")
	SliceMustPointer   = errors.New("slice must pointer")
)
