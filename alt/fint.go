// Copyright (c) 2021, Peter Ohler, All rights reserved.

package alt

import (
	"reflect"
	"strconv"
	"unsafe"
)

var intValFuncs = [8]valFunc{
	valInt,
	valIntAsString,
	valIntNotEmpty,
	valIntNotEmptyAsString,
	ivalInt,
	ivalIntAsString,
	ivalIntNotEmpty,
	ivalIntNotEmptyAsString,
}

func valInt(fi *finfo, rv reflect.Value, addr uintptr) (any, reflect.Value, bool) {
	return *(*int)(unsafe.Pointer(addr + fi.offset)), nilValue, false
}

func valIntAsString(fi *finfo, rv reflect.Value, addr uintptr) (any, reflect.Value, bool) {
	return strconv.FormatInt(int64(*(*int)(unsafe.Pointer(addr + fi.offset))), 10