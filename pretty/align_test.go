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
  [  1   2   3]
  [ 10  20  30]
  [100 200 300]
]`, string(out))
}

func TestWriteAlignArrayStrings(t *testing.T) {
	w := pretty.Writer{
		Width:    30,
		MaxDepth: 3,
		Align:    true,
	}
	data := []any{
		[]any{"alpha", "bravo", "charlie"},
		[]any{"a", "b", "c"},
		[]any{"andy", "betty"},
	}
	out, err := w.Marshal(data)
	tt.Nil(t, err)
	tt.Equal(t, `[
  ["alpha", "bravo", "charlie"],
  ["a"    , "b"    , "c"      ],
  ["andy" , "betty"]
]`, string(out))

	w.SEN = true
	out = w.Encode(data)
	tt.Equal(t, `[
  [alpha bravo charlie]
  [a     b     c      ]
  [andy  betty]
]`, string(out))
}

func TestWriteAlignArrayNested(t *testing.T) {
	w := pretty.Writer{
		Width:    40,
		MaxDepth: 3,
		Align:    true,
	}
	data := []any{
		[]any{1, 2, 3, []any{100, 200, 300}},
		[]any{1, 2, 3, "fourth"},
		[]any{10, 20, 30, []any{1, 20, 300}},
	}
	out := w.Encode(data)
	tt.Equal(t, `[
  [ 1,  2,  3, [100, 200, 300]],
  [ 1,  2,  3, "fourth"       ],
  [10, 20, 30, [  1,  20, 300]]
]`, string(out))

	w.SEN = true
	out = w.Encode(data)
	tt.Equal(t, `[
  [ 1  2  3 [100 200 300]]
  [ 1  2  3 fourth       ]
  [10 20 30 [  1  20 300]]
]`, string(out))
}

func TestWriteAlignMixed(t *testing.T) {
	w := pretty.Writer{
		Width:    20,
		MaxDepth: 3,
		Align:    true,
		SEN:      true,
	}
	out, err := w.Marshal([]any{
		[]any{1, 2, 3},
		map[string]any{"x": 1, "y": 2},
	})
	tt.Nil(t, err)
	tt.Equal(t, `[
  [1 2 3]
  {x: 1 y: 2}
]`, string(out))
}

func TestWriteAlignMapNumber(t *testing.T) {
	w := pretty.Writer{
		Width:    50,
		MaxDepth: 3,
		Align:    true,
	}
	data := []any{
		map[string]any{"x": 1, "y": 2},
		map[string]any{"z": 3, "y": 2},
		map[string]any{"x": 100, "y": 200, "z": 300},
		map[string]any{"x": 10, "z": 30},
	}
	out := w.Encode(data)
	tt.Equal(t, `[
  {"x":   1, "y":   2,         },
  {          "y":   2, "z":   3},
  {"x": 100, "y": 200, "z": 300},
  {"x":  10,           "z":  30}
]`, string(out))

	w.SEN = true
	out = w.Encode(data)
	tt.Equal(t, `[
  {x:   1 y:   2       }
  {       y:   2 z:   3}
  {x: 100 y: 200 z: 300}
  {x:  10        z:  30}
]`, string(out))
}

