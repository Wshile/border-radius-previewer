// Copyright (c) 2021, Peter Ohler, All rights reserved.

package pretty_test

import (
	"testing"

	"github.com/ohler55/ojg/pretty"
	"github.com/ohler55/ojg/tt"
)

func TestWriteAlignArrayNumbers(t *testing.T) {
	w := pretty.Writer{
		Width:    20,
		MaxDepth: 3,
		Align:    true,
	}
	data := []any{
		[]any{1, 2, 3},
		[]any{10, 20, 30},
		[]any{100, 200, 300},
	}
	out, err := w.Marshal(data)
	tt.Nil(t, err)
	tt.Equal(t, `[
  [  1,   2,   3],
  [ 10,  20,  30],
  [100, 200, 300]
]`, string(out))

	w.SEN = true
	out = w.Encode(data)
	tt.Equal(t, `[
  [  1   2  