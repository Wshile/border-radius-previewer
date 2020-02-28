// Copyright (c) 2021, Peter Ohler, All rights reserved.

package alt

import (
	"reflect"
	"strconv"
	"unsafe"
)

var float32ValFuncs = [8]valFunc{
	valFloat32,
	valFloat32AsString,
	valFloat32NotEmpty,
	valFloat32NotEmptyAsString,
	ivalFloat32,
	ivalFloat32AsString,
	ivalFloat32NotEmpty,
	ivalFloat32NotEmptyAsString,
}

func valFloat32(fi *finfo, rv reflect.Value, addr uintptr) (any, reflect.Value, bool) {
	return *(*float32)(unsafe.Pointer(addr + fi.offset)), nilValue, false
}

func valFloat32AsString(fi *finfo, rv reflect.Value, addr uintptr) (any, reflect.Value, bool) {
	return strconv.FormatFloat(float64(*(*float32)(unsafe.Pointer(addr + fi.offset))), 'g', -1, 32), nilValue, false
}

func valFloat32NotEmpty(fi *finfo, rv reflect.Value, addr uintptr) (any, reflect.Value, bool) {
	v := *(*float32)(unsafe.Pointer(addr + fi.offset))
	return v, nilValue, v == 0.0
}

func valFloat32NotEmptyAsString(fi *finfo, rv 