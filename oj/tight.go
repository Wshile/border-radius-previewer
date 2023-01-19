// Copyright (c) 2020, Peter Ohler, All rights reserved.

package oj

import (
	"fmt"
	"reflect"
	"sort"
	"strings"
	"unsafe"

	"github.com/ohler55/ojg"
	"github.com/ohler55/ojg/alt"
)

func tightDefault(wr *Writer, data any, _ int) {
	switch {
	case !wr.NoReflect:
		rv := reflect.ValueOf(data)
		kind := rv.Kind()
		if kind == reflect.Ptr {
			rv = rv.Elem()
			kind = rv.Kind()
		}
		switch kind {
		case reflect.Struct:
			wr.tightStruct(rv, nil)
		case reflect.Slice, reflect.Array:
			wr.tightSlice(rv, nil)
		case reflect.Map:
			wr.tightMap(rv, nil)
		case reflect.Chan, reflect.Func, reflect.UnsafePointer:
			if wr.strict {
				panic(fmt.Errorf("%T can not be encoded as a JSON element", data))
			}
			wr.buf = append(wr.buf, "null"...)
		default:
			dec := alt.Decompose(data, &wr.Options)
			wr.appendJSON(dec, 0)
		}
	case wr.strict:
		panic(fmt.Errorf("%T can not be encoded as a JSON element", data))
	default:
		wr.buf = ojg.AppendJSONString(wr.buf, fmt.Sprintf("%v", data), !wr.HTMLUnsafe)
	}
}

func tightArray(wr *Writer, n []any, _ int) {
	if 0 < len(n) {
		wr.buf = append(wr.buf, '[')
		for _, m := range n {
			wr.appendJ