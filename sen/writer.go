// Copyright (c) 2021, Peter Ohler, All rights reserved.

package sen

import (
	"encoding"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"reflect"
	"sort"
	"strconv"
	"strings"
	"time"
	"unsafe"

	"github.com/ohler55/ojg"
	"github.com/ohler55/ojg/alt"
)

const (
	spaces = "\n                                                                " +
		"                                                                "
	tabs = "\n\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t"
)

// Writer is a SEN writer that includes a reused buffer for reduced
// allocations for repeated encoding calls.
type Writer struct {
	ojg.Options
	buf           []byte
	w             io.Writer
	appendArray   func(wr *Writer, data []any, depth int)
	appendObject  func(wr *Writer, data map[string]any, depth int)
	appendDefault func(wr *Writer, data any, depth int)
	appendString  func(buf []byte, s string, htmlSafe bool) []byte
	findex        byte
	needSep       bool
}

// SEN writes data, SEN encoded. On error, an empty string is returned.
func (wr *Writer) SEN(data any) string {
	defer func() {
		if r := recover(); r != nil {
			wr.buf = wr.buf[:0]
		}
	}()
	return string(wr.MustSEN(data))
}

// MustSEN writes data, SEN encoded as a []byte and not a string like the
// SEN() function. On error a panic is called with the error. The returned
// buffer is the Writer buffer and is reused on the next call to write. If
// returned value is to be preserved past a second invocation then the buffer
// should be copied.
func (wr *Writer) MustSEN(data any) []byte {
	wr.w = nil
	if wr.InitSize <= 0 {
		wr.InitSize = 256
	}
	if cap(wr.buf) < wr.InitSize {
		wr.buf = make([]byte, 0, wr.InitSize)
	} else {
		wr.buf = wr.buf[:0]
	}
	wr.calcFieldsIndex()
	if wr.Color {
		wr.colorSEN(data, 0)
	} else {
		wr.appendString = ojg.AppendSENString
		if wr.Tab || 0 < wr.Indent {
			wr.appendArray = appendArray
			if wr.Sort {
				wr.appendObject = appendSortObject
			} else {
				wr.appendObject = appendObject
			}
			wr.appendDefault = appendDefault
		} else {
			wr.appendArray = tightArray
			if wr.Sort {
				wr.appendObject = tightSortObject
			} else {
				wr.appendObject = tightObject
			}
			wr.appendDefault = tightDefault
		}
		wr.appendSEN(data, 0)
	}
	return wr.buf
}

// Write a SEN string for the data provided.
func (wr *Writer) Write(w io.Writer, data any) (err error) {
	defer func() {
		if r := recover(); r != nil {
			wr.buf = wr.buf[:0]
			err = ojg.NewError(r)
		}
	}()
	wr.MustWrite(w, data)
	return
}

// MustWrite a SEN string for the data provided. If an error occurs panic is
// called with the error.
func (wr *Writer) MustWrite(w io.Writer, data any) {
	wr.w = w
	if wr.InitSize <= 0 {
		wr.InitSize = 256
	}
	if wr.WriteLimit <= 0 {
		wr.WriteLimit = 1024
	}
	if cap(wr.buf) < wr.InitSize {
		wr.buf = make([]byte, 0, wr.InitSize)
	} else {
		wr.buf = wr.buf[:0]
	}
	wr.calcFieldsIndex()
	if wr.Color {
		wr.colorSEN(data, 0)
	} else {
		wr.appendString = ojg.AppendSENString
		if wr.Tab || 0 < wr.Indent {
			wr.appendArray = appendArray
			if wr.Sort {
				wr.appendObject = appendSortObject
			} else {
				wr.appendObject = appendObject
			}
			wr.appendDefault = appendDefault
		} else {
			wr.appendArray = tightArray
			if wr.Sort {
				wr.appendObject = tightSortObject
			} else {
				wr.appendObject = tightObject
			}
			wr.appendDefault = tightDefault
		}
		wr.appendSEN(data, 0)
	}
	if 0 < len(wr.buf) {
		if _, err := wr.w.Write(wr.buf); err != nil {
			panic(err)
		}
	}
}

func (wr *Writer) calcFieldsIndex() {
	wr.findex = 0
	if wr.NestEmbed {
		wr.findex |= maskNested
	}
	if 0 < wr.Indent {
		wr.findex |= maskPretty
	}
	if wr.UseTags {
		wr.findex |= maskByTag
	} else if wr.KeyExact {
		wr.findex |= maskExact
	}
}

func (wr *Writer) appendSEN(data any, depth int) {
	wr.needSep = true
	switch td := data.(type) {
	case nil:
		wr.buf = append(wr.buf, "null"...)

	case bool:
		if td {
			wr.buf = append(wr.buf, "true"...)
		} else {
			wr.buf = append(wr.buf, "false"...)
		}

	case int:
		wr.buf = strconv.AppendInt(wr.buf, int64(td), 10)
	case int8:
		wr.buf = strconv.AppendInt(wr.buf, int64(td), 10)
	case int16:
		wr.buf = strconv.AppendInt(wr.buf, int64(td), 10)
	case int32:
		wr.buf = strconv.AppendInt(wr.buf, int64(td), 10)
	case int64:
		wr.buf = strconv.AppendInt(wr.buf, td, 10)
	case uint:
		wr.buf = strconv.AppendUint(wr.buf, uint64(td), 10)
	case uint8:
		wr.buf = strconv.AppendUint(wr.buf, uint64(td), 10)
	case uint16:
		wr.buf = strconv.AppendUint(wr.buf, uint64(td), 10)
	case uint32:
		wr.buf = strconv.AppendUint(wr.buf, uint64(td), 10)
	case uint64:
		wr.buf = strconv.AppendUint(wr.buf, td, 10)

	case float32:
		wr.buf = strconv.AppendFloat(wr.buf, float64(td), 'g', -1, 32)
	case float64:
		wr.buf = strconv.AppendFloat(wr.buf, td, 'g', -1, 64)

	case string:
		wr.buf = wr.appendString(wr.buf, td, !wr.HTMLUnsafe)

	case []byte:
		switch wr.BytesAs {
		case ojg.BytesAsBase64:
			wr.buf = wr.appendString(wr.buf, base64.StdEncoding.EncodeToString(td), !wr.HTMLUnsafe)
		case ojg.BytesAsArray:
			a := make([]any, len(td))
			for i, m := range td {
				a[i] = int64(m)
			}
			wr.appendArray(wr, a, depth)
		default:
			wr.buf = wr.appendString(wr.buf, string(td), !wr.HTMLUnsafe)
		}

	case time.Time:
		wr.buf = wr.AppendTime(wr.buf, td, true)

	case []any:
		wr.appendArray(wr, td, depth)
		wr.needSep = false

	case map[string]any:
		wr.appendObject(wr, td, depth)
		wr.needSep = false

	case alt.Simplifier:
		wr.appendSEN(td.Simplify(), depth)
	case alt.Genericer:
		wr.appendSEN(td.Generic().Simplify(), depth)
	case json.Marshaler:
		out, err := td.MarshalJSON()
		if err != nil {
			panic(err)
		}
		wr.buf = append(wr.buf, out...)
	case encoding.TextMarshaler:
		out, err := td.MarshalText()
		if err != nil {
			panic(err)
		}
		wr.buf = wr.appendString(wr.buf, string(out), !wr.HTMLUnsafe)

	default:
		wr.appendDefault(wr, data, depth)
		if 0 < len(wr.buf) {
			switch wr.buf[len(wr.buf)-1] {
			case '}', ']':
				wr.needSep = false
			default:
			}
		}
	}
	if wr.w != nil && wr.WriteLimit < len(wr.buf) {
		if _, err := wr.w.Write(wr.buf); err != nil {
			panic(err)
		}
		wr.buf = wr.buf[:0]
	}
}

func appendDefault(wr *Writer, data any, depth int) {
	if !wr.NoReflect {
		rv := reflect.ValueOf(data)
		kind := rv.Kind()
		if kind == reflect.Ptr {
			rv = rv.Elem()
			kind = rv.Kind()
		}
		switch kind {
		case reflect.Struct:
			wr.appendStruct(rv, depth, nil)
		case reflect.Slice, reflect.Array:
			wr.appendSlice(rv, depth, nil)
		case reflect.Map:
			wr.appendMap(rv, depth, nil)
		default:
			// Not much should get here except Complex and n