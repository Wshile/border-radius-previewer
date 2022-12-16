// Copyright (c) 2020, Peter Ohler, All rights reserved.

package jp_test

import (
	"fmt"
	"testing"

	"github.com/ohler55/ojg/gen"
	"github.com/ohler55/ojg/jp"
	"github.com/ohler55/ojg/oj"
	"github.com/ohler55/ojg/tt"
)

type delData struct {
	path     string
	data     string // JSON
	expect   string // JSON
	err      string
	noNode   bool
	noSimple bool
}

var (
	delTestData = []*delData{
		{path: "@.a", data: `{}`, expect: `{}`},
		{path: "@.a", data: `{"a":3}`, expect: `{}`},
		{path: "[1]", data: `[1,2,3]`, expect: `[1,null,3]`},
		{path: "a.*", data: `{"a":{"x":1,"y":2}}`, expect: `{"a":{}}`},
		{path: "[*]", data: `[1,2,3]`, expect: `[null,null,null]`},
		{path: "a[0]", data: `{}`, expect: `{}`},
		{path: "a[1,2]", data: `{"a":[0,1,2,3]}`, expect: `{"a":[0,null,null,3]}`},
		{path: "['a','b']", data: `{"a":1,"b":2,"c":3}`, expect: `{"c":3}`},

		{path: "", data: `{}`, err: "can not delete with an empty expression"},
		{path: "$", data: `{}`, err: "can not delete with an expression ending with a Root"},
		{path: "@", data: `{}`, err: "can not delete with an expression ending with a At"},
		{path: "a.b", data: `{"a":4}`, err: "/can not follow a .+ at 'a'/"},
		{path: "[0].1", data: `[1]`, err: "/can not follow a .+ at '\\[0\\]'/"},
		{path: "[1]", data: `[1]`, err: "can not follow out of bounds array index at '[1]'"},
	}
	delOneTestData =