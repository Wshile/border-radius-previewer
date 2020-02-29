// Copyright (c) 2021, Peter Ohler, All rights reserved.

package alt

import (
	"reflect"
	"strconv"
	"unsafe"
)

var int64ValFuncs = [8]valFunc{
	valInt64,
	valInt64AsString,
	valInt64NotEmpty,
	valInt64NotEmptyAsString,
	ivalInt64,
	ivalInt64AsString,
	ivalInt64NotEmpty,
	ivalInt64NotEmptyAsString,
}

func valInt64(fi *finfo, rv reflect.Value, addr uintptr) (any, reflect.Value, bool) {
	return *(*int64)(unsafe.Pointer(addr + fi.offset)), nilValue, false
}

func valInt64AsString(fi *finfo, rv reflect.Value, addr uintptr) (any, reflect.Value, bool) {
	return strconv.FormatInt(*(*int64)(unsafe.Pointer(addr + fi.offset)), 10), nilValue, false
}

func valInt64NotEmpty(fi *finfo, rv reflect.Value, addr uintptr) (any, reflect.Value, bool) {
	v := *(*int64)(unsafe.Pointer(addr + fi.offset))
	return v, nilValue, v == 0
}

func valInt64NotEmptyAsString(fi *finfo, rv reflect.Value, addr uintptr) (any, reflect.Value, bool) {
	v := *(*int64)(unsafe.Pointer(addr + fi.offset))
	if v == 0 {
		return nil, nilValue, true
	}
	return strconv.FormatInt(v, 10), nilVa