// Copyright (c) 2021, Peter Ohler, All rights reserved.

package sen

import (
	"reflect"
	"strconv"
	"unsafe"
)

var intAppendFuncs = [8]appendFunc{
	appendInt,
	appendIntAsString,
	appendIntNotEmpty,
	appendIntNotEmptyAsString,
	iappendInt,
	iappendIntAsString,
	iappendIntNotEmpty,
	iappendIntNotEmptyAsString,
}

func appendInt(fi *finfo, buf []byte, rv reflect.Value, addr uintptr, safe bool) ([]byte, any, append