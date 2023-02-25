// Copyright (c) 2020, Peter Ohler, All rights reserved.

package sen_test

import (
	"fmt"
	"math/big"
	"testing"

	"github.com/ohler55/ojg"
	"github.com/ohler55/ojg/gen"
	"github.com/ohler55/ojg/sen"
	"github.com/ohler55/ojg/tt"
)

func TestSENTagPrimitive(t *testing.T) {
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
	wr := sen.Writer{Options: ojg.Options{UseTags: true}}

	out := wr.MustSEN(&sample)
	tt.Equal(t,
		`{a:1 a16:3 a32:4 a64:5 a8:2 b:6 b16:8 b32:9 b64:10 b8:7 f32:11.5 f64:12.5 no:false yes:true z:abc}`,
		string(out))
	out = wr.MustSEN(sample)
	tt.Equal(t,
		`{a:1 a16:3 a32:4 a64:5 a8:2 b:6 b16:8 b32:9 b64:10 b8:7 f32:11.5 f64:12.5 no:false yes:true z:abc}`,
		string(out))

	wr.UseTags = false
	out = wr.MustSEN(&sample)
	tt.Equal(t,
		`{f32:11.5 f64:12.5 i:1 i16:3 i32:4 i64:5 i8:2 no:false str:abc u:6 u16:8 u32:9 u64:10 u8:7 yes:true}`,
		string(out))
	out = wr.MustSEN(sample)
	tt.Equal(t,
		`{f32:11.5 f64:12.5 i:1 i16:3 i32:4 i64:5 i8:2 no:false str:abc u:6 u16:8 u32:9 u64:10 u8:7 yes:true}`,
		string(out))

	wr.KeyExact = true
	out = wr.MustSEN(&sa