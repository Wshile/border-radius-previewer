// Copyright (c) 2020, Peter Ohler, All rights reserved.

package alt_test

import (
	"testing"

	"github.com/ohler55/ojg"
	"github.com/ohler55/ojg/alt"
	"github.com/ohler55/ojg/gen"
	"github.com/ohler55/ojg/tt"
)

func TestDecomposeTagPrimitive(t *testing.T) {
	type Sample struct {
		Yes bool    `json:"yes"`
		No  bool    `json:"no"`
		I   int     `json:"a"`
		I8  int8    `json:"a8"`
		I16 int16   `json:"a16"`
		I32 int32   `json:"a32"`
		I64 int64   `json:"a64"`
		U   uint    `json:"b"`
		U8  uint8   `json:"b8"`
		U16 uint16  `json:"b16"`
		U32 uint32  `json:"b32"`
		U64 uint64  `json:"b64"`
		F32 float32 `json:"f32"`
		F64 float64 `json:"f64"`
		Str string  `json:"z"`
	}
	sample := Sample{
		Yes: true,
		No:  false,
		I:   1,
		I8:  2,
		I16: 3,
		I32: 4,
		I64: 5,
		U:   6,
		U8:  7,
		U16: 8,
		U32: 9,
		U64: 10,
		F32: 11.5,
		F64: 12.5,
		Str: "abc",
	}
	opt := ojg.Options{UseTags: true}
	out := alt.Decompose(&sample, &opt)
	tt.Equal(t,
		map[string]any{
			"a":   1,
			"a16": 3,
			"a32": 4,
			"a64": 5,
			"a8":  2,
			"b":   6,
			"b16": 8,
			"b32": 9,
			"b64": 10,
			"b8":  7,
			"f32": 11.5,
			"f64": 12.5,
			"no":  false,
			"yes": true,
			"z":   "abc",
		}, out)

	out = alt.Decompose(sample, &opt)
	tt.Equal(t,
		map[string]any{
			"a":   1,
			"a16": 3,
			"a32": 4,
			"a64": 5,
			"a8":  2,
			"b":   6,
			"b16": 8,
			"b32": 9,
			"b64": 10,
			"b8":  7,
			"f32": 11.5,
			"f64": 12.5,
			"no":  false,
			"yes": true,
			"z":   "abc",
		}, out)

	opt.UseTags = false
	out = alt.Decompose(&sample, &opt)
	tt.Equal(t,
		map[string]any{
			"f32": 11.5,
			"f64": 12.5,
			"i":   1,
			"i16": 3,
			"i32": 4,
			"i64": 5,
			"i8":  2,
			"no":  false,
			"str": "abc",
			"u":   6,
			"u16": 8,
			"u32": 9,
			"u64": 10,
			"u8":  7,
			"yes": true,
		}, out)
	out = alt.Decompose(sample, &opt)
	tt.Equal(t,
		map[string]any{
			"f32": 11.5,
			"f64": 12.5,
			"i":   1,
			"i16": 3,
			"i32": 4,
			"i64": 5,
			"i8":  2,
			"no":  false,
			"str": "abc",
			"u":   6,
			"u16": 8,
			"u32": 9,
			"u64": 10,
			"u8":  7,
			"yes": true,
		}, out)

	opt.KeyExact = true
	out = alt.Decompose(&sample, &opt)
	tt.Equal(t,
		map[string]any{
			"F32": 11.5,
			"F64": 12.5,
			"I":   1,
			"I16": 3,
			"I32": 4,
			"I64": 5,
			"I8":  2,
			"No":  false,
			"Str": "abc",
			"U":   6,
			"U16": 8,
			"U32": 9,
			"U64": 10,
			"U8":  7,
			"Yes": true,
		}, out)
	out = alt.Decompose(sample, &opt)
	tt.Equal(t,
		map[string]any{
			"F32": 11.5,
			"F64": 12.5,
			"I":   1,
			"I16": 3,
			"I32": 4,
			"I64": 5,
			"I8":  2,
			"No":  false,
			"Str": "abc",
			"U":   6,
			"U16": 8,
			"U32": 9,
			"U64": 10,
			"U8":  7,
			"Yes": true,
		}, out)
}

func TestDecomposeTagAsString(t *testing.T) {
	type Sample struct {
		Yes bool    `json:"yes,string"`
		No  bool    `json:"no,string"`
		I   int     `json:"a,string"`
		I8  int8    `json:"a8,string"`
		I16 int16   `json:"a16,string"`
		I32 int32   `json:"a32,string"`
		I64 int64   `json:"a64,string"`
		U   uint    `json:"b,string"`
		U8  uint8   `json:"b8,string"`
		U16 uint16  `json:"b16,string"`
		U32 uint32  `json:"b32,string"`
		U64 uint64  `json:"b64,string"`
		F32 float32 `json:"f32,string"`
		F64 float64 `json:"f64,string"`
		Str string  `json:"z,string"`
	}
	sample := Sample{
		Yes: true,
		No:  false,
		I:   1,
		I8:  2,
		I16: 3,
		I32: 4,
		I64: 5,
		U:   6,
		U8:  7,
		U16: 8,
		U32: 9,
		U64: 10,
		F32: 11.5,
		F64: 12.5,
		Str: "abc",
	}
	opt := ojg.Options{UseTags: true}

	out := alt.Decompose(&sample, &opt)
	tt.Equal(t,
		map[string]any{
			"a":   "1",
			"a16": "3",
			"a32": "4",
			"a64": "5",
			"a8":  "2",
			"b":   "6",
			"b16": "8",
			"b32": "9",
			"b64": "10",
			"b8":  "7",
			"f32": "11.5",
			"f64": "12.5",
			"no":  "false",
			"yes": "true",
			"z":   "abc",
		}, out)
	out = alt.Decompose(sample, &opt)
	tt.Equal(t,
		m