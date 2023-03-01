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
func (wr *Writer) Write(w io.Writer, data any) (err e