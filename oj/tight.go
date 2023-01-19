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
			wr.appendJSON(m, 0)
			wr.buf = append(wr.buf, ',')
		}
		wr.buf[len(wr.buf)-1] = ']'
	} else {
		wr.buf = append(wr.buf, "[]"...)
	}
}

func tightObject(wr *Writer, n map[string]any, _ int) {
	comma := false
	wr.buf = append(wr.buf, '{')
	for k, m := range n {
		switch tm := m.(type) {
		case nil:
			if wr.OmitNil {
				continue
			}
		case string:
			if wr.OmitEmpty && len(tm) == 0 {
				continue
			}
		case map[string]any:
			if wr.OmitEmpty && len(tm) == 0 {
				continue
			}
		case []any:
			if wr.OmitEmpty && len(tm) == 0 {
				continue
			}
		}
		wr.buf = ojg.AppendJSONString(wr.buf, k, !wr.HTMLUnsafe)
		wr.buf = append(wr.buf, ':')
		wr.appendJSON(m, 0)
		wr.buf = append(wr.buf, ',')
		comma = true
	}
	if comma {
		wr.buf[len(wr.buf)-1] = '}'
	} else {
		wr.buf = append(wr.buf, '}')
	}
}

func tightSortObject(wr *Writer, n map[string]any, _ int) {
	comma := false
	wr.buf = append(wr.buf, '{')
	keys := make([]string, 0, len(n))
	for k := range n {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for _, k := range keys {
		m := n[k]
		switch tm := m.(type) {
		case nil:
			if wr.OmitNil {
				continue
			}
		case string:
			if wr.OmitEmpty && len(tm) == 0 {
				continue
			}
		case map[string]any:
			if wr.OmitEmpty && len(tm) == 0 {
				continue
			}
		case []any:
			if wr.OmitEmpty && len(tm) == 0 {
				continue
			}
		}
		wr.buf = ojg.AppendJSONString(wr.buf, k, !wr.HTMLUnsafe)
		wr.buf = append(wr.buf, ':')
		wr.appendJSON(m, 0)
		wr.buf = append(wr.buf, ',')
		comma = true
	}
	if comma {
		wr.buf[len(wr.buf)-1] = '}'
	} else {
		wr.buf = append(wr.buf, '}')
	}
}

func (wr *Writer) tightStruct(rv reflect.Value, si *sinfo) {
	if si == nil {
		si = getSinfo(rv.Interface(), wr.OmitEmpty)
	}
	fields := si.fields[wr.findex]
	wr.buf = append(wr.buf, '{')
	var v any
	comma := false
	if 0 < len(wr.CreateKey) {
		wr.buf = wr.appendString(wr.buf, wr.CreateKey, !wr.HTMLUnsafe)
		wr.buf = append(wr.buf, `:"`...)
		if wr.FullTypePath {
			wr.buf = append(wr.buf, si.rt.PkgPath()...)
			wr.buf = append(wr.buf, '/')
			wr.buf = append(wr.buf, si.rt.Name()...)
		} else {
			wr.buf = append(wr.buf, si.rt.Name()...)
		}
		wr.buf = a