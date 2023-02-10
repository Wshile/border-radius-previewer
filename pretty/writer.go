// Copyright (c) 2021, Peter Ohler, All rights reserved.

package pretty

import (
	"fmt"
	"io"
	"math"

	"github.com/ohler55/ojg"
)

const (
	nullStr  = "null"
	trueStr  = "true"
	falseStr = "false"
	spaces   = "\n                                                                " +
		"                                                                "
)

// Writer writes data in either JSON or SEN format using setting to determine
// the output.
type Writer struct {
	ojg.Options

	// Width is the suggested maximum width. In some cases it may not be
	// possible to stay within the specified width.
	Width int

	// MaxDepth is the maximum depth of an element on a single line.
	MaxDepth int

	// Align if true attempts to align elements of children in list.
	Align bool

	// SEN format if true otherwise JSON encoding.
	SEN bool

	buf []byte
	w   io.Writer
}

// Encode data. Any panics during encoding will cause an empty return but will
// not fail.The returned buffer is the Writer buffer and is reused on the next
// call to write. If returned value is to be preserved past a second
// invocation then the buffer should be copied.
func (w *Writer) Encode(data any) []byte {
	b, _ := w.encode(data)

	return b
}

// Marshal data. The same as Encode but a panics during encoding will result
// in an error return.
func (w *Writer) Marshal(data any) ([]byte, error) {
	if _, err := w.encode(data); err != nil {
		return nil, err
	}
	out := make([]byte, len(w.buf))
	copy(out, w.buf)

	return out, nil
}

// Write encoded data to the op.Writer. The returned buffer is the Writer
// buffer and is reused on the next call to write. If returned value is to be
// preserved past a second invocation then the buffer should be copied.
func (w *Writer) Write(wr io.Writer, data any) (err error) {
	w.w = wr
	_, err = w.encode(data)

	return
}
func (w *Writer) config(args []any) {
	for _, arg := range args {
		switch ta := arg.(type) {
		case int:
			w.Width = ta
		case float64:
			if 0.0 < ta {
				if ta < 1.0 {
					w.MaxDepth = int(math.Round(ta * 10.0))
				} else {
					w.Width = int(ta)
					w.MaxDepth = int(math.Round((ta - float64(w.Width)) * 10.0))
				}
				if w.MaxDepth == 0 { // use the default
					w.MaxDepth = 2
				}
			}
		case bool:
			w.Align = ta
		case *ojg.Options:
			sw := w.w
			w.Options = *ta
			w.w = sw
		}
	}
}

func (w *Writer) encode(data any) (out []byte, err error) {
	if w.InitSize == 0 {
		w.InitSize = 256
	}
	if len(spaces)-1 < w.Width {
		w.Width = len(spaces) - 1
	}
	if w.WriteLimit == 0 {
		w.