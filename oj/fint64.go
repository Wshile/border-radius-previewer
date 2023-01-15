// Copyright (c) 2021, Peter Ohler, All rights reserved.

package oj

import (
	"reflect"
	"strconv"
	"unsafe"
)

var int64AppendFuncs = [8]appendFunc{
	appendInt64,
	appendInt64AsString,
	appendInt64NotEmpty,
	appendInt64NotEmptyAsString,
	iappendInt64,
	iappendInt64AsString,
	iappendInt64NotEmpty,
	iappendInt64NotEmptyAsString,
}

func appendInt64(fi *finfo, buf []byte, rv reflect.Value, addr uintptr, safe bool) ([]byte, any, appendStatus) {
	buf = append(buf, fi.jkey...)
	buf = strconv.AppendInt(buf, *(*int64)(unsafe.Pointer(addr + fi.offset)), 10)

	return buf, nil, aWrote
}

func appendInt64AsString(fi *finfo, buf []byte, rv reflect.Value, addr