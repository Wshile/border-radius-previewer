// Copyright (c) 2021, Peter Ohler, All rights reserved.

package alt

import (
	"reflect"
	"strconv"
	"unsafe"
)

var float64ValFuncs = [8]valFunc{
	valFloat64,
	valFloat64AsString,
	valFloat64NotEmpty,
	valFloat64NotEmptyAsString,
	ivalFloat64,
	ivalFloat64AsString,
	ivalFloat64NotEmpty,
	ivalFloat64NotEmptyAsString,
}

func valFloat64(fi *finfo, rv reflect.Value, addr uintptr) (any, reflect.Value, bool) {
	return *(*float64)(unsafe.Pointer(addr + fi.offset)), nilValue, false
}

func valFloat64AsString(fi *finfo, rv reflect.Value, addr uintptr) (any, reflect.Value, bool) {
	return strconv.FormatFloat(*(*float64)(unsafe.Pointer(addr + fi.offset)), 'g', -1, 64), nilValue, false
}

func valFloat64NotEmpty(fi *finfo, rv reflect.Value, addr uintptr) (any, reflect.Value, bool) {
	v := *(*float64)(unsafe.Pointer(addr + fi.offset))
	return v, nilValue, v == 0.0
}

func valFloat64NotEmptyAsString(fi *finfo, rv reflect.Value, addr uintptr) (any, reflect.Value, bool) {
	v := *(*float64