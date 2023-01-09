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

type setData struct {
	path     string
	data     string // JSON
	value    any
	expect   string // JSON
	err      string
	noNode   bool
	noSimple bool
}

type setReflectData struct {
	path   string
	data   any
	value  any
	expect string // JSON
	err    string
}

var (
	setTestData = []*setData{
		{path: "@.a", data: `{}`, value: 3, expect: `{"a":3}`},
		{path: "a.b", data: `{}`, value: 3, expect: `{"a":{"b":3}}`},
		{path: "a.b", data: `{"a":{}}`, value: 3, expect: `{"a":{"b":3}}`},
		{path: "[1]", data: `[1,2,3]`, value: 5, expect: `[1,5,3]`},
		{path: "[-1]", data: `[1,2,3]`, value: 5, expect: `[1,2,5]`},
		{path: "[1].a", data: `[1,{},3]`, value: 5, expect: `[1,{"a":5},3]`},
		{path: "[*]", data: `[1,2,3]`, value: 5, expect: `[5,5,5]`},
		{path: ".*", data: `{"a":1,"b":2}`, value: 5, expect: `{"a":5,"b":5}`},
		{path: "$.*.a", data: `{"a":{"a":1,"b":2},"b":{"a":2}}`, value: 5, expect: `{"a":{"a":5,"b":2},"b":{"a":5}}`},
		{path: "[*].a", data: `[{"a":1,"b":2},{"a":2}]`, value: 5, expect: `[{"a":5,"b":2},{"a":5}]`},
		{path: "..a", data: `{"a":{"a":1,"b":2},"b":{"a":2}}`, value: 5, expect: `{"a":5,"b":{"a":5}}`},
		{path: "..a", data: `[{"a":1,"b":2},{"a":2}]`, value: 5, expect: `[{"a":5,"b":2},{"a":5}]`},
		{path: "[-1,'x'].a", data: `[{"a":1,"b":2},{"a":2}]`, value: 5, expect: `[{"a":1,"b":2},{"a":5}]`},
		{path: "[1,'a'].a", data: `{"a":{"a":1,"b":2},"b":{"a":2}}`, value: 5, expect: `{"a":{"a":5,"b":2},"b":{"a":2}}`},
		{path: "[:-1:2].a", data: `[{"a":1,"b":2},{"a":2},{"a":3}]`, value: 5, expect: `[{"a":5,"b":2},{"a":2},{"a":5}]`},
		{path: "[-1:0:-2].a", data: `[{"a":1,"b":2},{"a":2},{"a":3}]`, value: 5, expect: `[{"a":5,"b":2},{"a":2},{"a":5}]`},
		{path: "[:5].a", data: `[{"a":1,"b":2},{"a":2},{"a":3}]`, value: 5, expect: `[{"a":5,"b":2},{"a":5},{"a":5}]`},
		{path: "[?(@.b == 2)].a", data: `[{"a":1,"b":2},{"a":2},{"a":3}]`, value: 5, expect: `[{"a":5,"b":2},{"a":2},{"a":3}]`},
		{path: "a[0]", data: `{}`, value: 3, expect: `{"a":[3]}`},
		{path: "*.x", data: `{"a":null}`, value: 3, expect: `{"a":null}`},
		{path: "[*].x", data: "[null]", value: 3, expect: `[null]`},
		{path: "...x", data: "[null]", value: 3, expect: `[null]`},
		{path: "[0,1].x", data: "[null]", value: 3, expect: `[null]`},
		{path: "['a','b'].x", data: "[null]", value: 3, expect: `[null]`},
		{path: "['a','b'].x", data: `{"a":null}`, value: 3, expect: `{"a":null}`},
		{path: "a[1,2]", data: `{"a":[0,1,2,3]}`, value: 5, expect: `{"a":[0,5,5,3]}`},
		{path: "['a','b']", data: `{"a":1,"b":2,"c":3}`, value: 5, expect: `{"a":5,"b":5,"c":3}`},

		{path: "", data: `{}`, value: 3, err: "can not set with an empty expression"},
		{path: "$", data: `{}`, value: 3, err: "can not set with an expression ending with a Root"},
		{path: "@", data: `{}`, value: 3, err: "can not set with an expression ending with a At"},
		{path: "a", data: `{}`, value: func() {}, err: "can not set a func() in a gen.Object", noSimple: true},
		{path: "a.b", data: `{"a":4}`, value: 3, err: "/can not follow a .+ at 'a'/"},
		{path: "a[-1]", data: `{}`, value: 3, err: "can not deduce the length of the array to add at 'a'"},
		{path: "a[1,2].x", data: `{}`, value: 3, err: "can not deduce what element to add at 'a'"},
		{path: "[0].1", data: `[1]`, value: 3, err: "/can not follow a .+ at '\\[0\\]'/"},
		{path: "[1]", data: `[1]`, value: 3, err: "can not follow out of bounds array index at '[1]'"},
	}
	setOneTestData = []*setData{
		{path: "@.a", data: `{}`, value: 3, expect: `{"a":3}`},
		{path: "a.b", data: `{}`, value: 3, expect: `{"a":{"b":3}}`},
		{path: "a.b", data: `{"a":{}}`, value: 3, expect: `{"a":{"b":3}}`},
		{path: "a.b", data: `{"a":{"b":2}}`, value: 3, expect: `{"a":{"b":3}}`},
		{path: "[1]", data: `[1,2,3]`, value: 5, expect: `[1,5,3]`},
		{path: "[-1]", data: `[1,2,3]`, value: 5, expect: `[1,2,5]`},
		{path: "[1].a", data: `[1,{},3]`, value: 5, expect: `[1,{"a":5},3]`},
		{path: "[*]", data: `[1,2,3]`, value: 5, expect: `[5,2,3]`},
		{path: "*", data: `{"a":1}`, value: 5, expect: `{"a":5}`},
		{path: "$.*.a", data: `{"a":{"a":1}}`, value: 5, expect: `{"a":{"a":5}}`},
		{path: "[*].a", data: `[{"a":1,"b":2}]`, value: 5, expect: `[{"a":5,"b":2}]`},
		{path: "..a", data: `{"a":1,"b":2}`, value: 5, expect: `{"a":5,"b":2}`},
		{path: "..a", data: `[{"a":1,"b":2}]`, value: 5, expect: `[{"a":5,"b":2}]`},
		{path: "..b", data: `{"x":{"a":{}}}`, value: 5, expect: `{"x":{"a":{"b":5}}}`},
		{path: "[-1,'x'].a", data: `[{"a":1,"b":2},{"a":2}]`, value: 5, expect: `[{"a":1,"b":2},{"a":5}]`},
		{path: "[1,'a'].a", data: `{"a":{"a":1,"b":2},"b":{"a":2}}`, value: 5, expect: `{"a":{"a":5,"b":2},"b":{"a":2}}`},
		{path: "[:-1:2].a", data: `[{"a":1,"b":2},{"a":2},{"a":3}]`, value: 5, expect: `[{"a":5,"b":2},{"a":2},{"a":3}]`},
		{path: "[-1:0:-2].a", data: `[{"a":1,"b":2},{"a":2},{"a":3}]`, value: 5, expect: `[{"a":1,"b":2},{"a":2},{"a":5}]`},
		{path: "[:5].a", data: `[{"a":1,"b":2},{"a":2},{"a":3}]`, value: 5, expect: `[{"a":5,"b":2},{"a":2},{"a":3}]`},
		{path: "[?(@.b == 2)].a", data: `[{"a":1,"b":2},{"a":2},{"a":3}]`, value: 5, expect: `[{"a":5,"b":2},{"a":2},{"a":3}]`},
		{path: "[-5:3].a", data: `[{"a":1,"b":2},{"a":2},{"a":3}]`, value: 5, expect: `[{"a":1,"b":2},{"a":2},{"a":3}]`},
		{path: "a[0]", data: `{}`, value: 3, expect: `{"a":[3]}`},
		{path: "*.x", data: `{"a":null}`, value: 3, expect: `{"a":null}`},
		{path: "[*].x", data: "[null]", value: 3, expect: `[null]`},
		{path: "...x", data: "[null]", value: 3, expect: `[null]`},
		{path: "[0,1].x", data: "[null]", value: 3, expect: `[null]`},
		{path: "['a','b'].x", data: "[null]", value: 3, expect: `[null]`},
		{path: "['a','b'].x", data: `{"a":null}`, value: 3, expect: `{"a":null}`},
		{path: "a[1,2]", data: `{"a":[0,1,2,3]}`, value: 5, expect: `{"a":[0,5,2,3]}`},
		{path: "['a','b']", data: `{"a":1,"b":2,"c":3}`, value: 5, expect: `{"a":5,"b":2,"c":3}`},

		{path: "", data: `{}`, value: 3, err: "can not set with an empty expression"},
		{path: "$", data: `{}`, value: 3, err: "can not set with an expression ending with a Root"},
		{path: "@", data: `{}`, value: 3, err: "can not set with an expression ending with a At"},
		{path: "a", data: `{}`, value: func() {}, err: "can not set a func() in a gen.Object", noSimple: true},
		{path: "a.b", data: `{"a":4}`, value: 3, err: "/can not follow a .+ at 'a'/"},
		{path: "a[-1]", data: `{}`, value: 3, err: "can not deduce the length of the array to add at 'a'"},
		{path: "a[1,2].x", data: `{}`, value: 3, err: "can not deduce what element to add at 'a'"},
		{path: "[0].1", data: `[1]`, value: 3, err: "/can not follow a .+ at '\\[0\\]'/"},
		{path: "[1]", data: `[1]`, value: 3, err: "can not follow out of bounds array index at '[1]'"},
	}
	setReflectTestData = []*setReflectData{
		{path: "a", data: &Sample{A: 1, B: "a string"}, value: 3, expect: `{"^":"Sample","a":3,"b":"a string"}`},
		{path: "x.a", data: &Any{X: map[string]any{"a": 1}}, value: 3, expect: `{"^":"Any","x":{"a":3}}`},
		{path: "x.a", data: &Any{X: &Sample{A: 1}}, value: 3, expect: `{"^":"Any","x":{"^":"Sample","a":3,"b":""}}`},
		{path: "x.a", data: map[string]any{"x": &Sample{A: 1}}, value: 3, expect: `{"x":{"^":"Sample","a":3,"b":""}}`},
		{path: "[1]", data: []int{1, 2, 3}, value: 5, expect: `[1,5,3]`},
		{path: "[-2]", data: []int{1, 2, 3}, value: 5, expect: `[1,5,3]`},
		{path: "[1].x", data: []*Any{{X: 1}, {X: 2}, {X: 3}}, value: 5, expect: `[{"^":"Any","x":1},{"^":"Any","x":5},{"^":"Any","x":3}]`},
		{path: "$.*.x", data: []*Any{{X: 1}, {X: 2}, {X: 3}}, value: 5, expect: `[{"^":"Any","x":5},{"^":"Any","x":5},{"^":"Any","x":5}]`},
		{path: "[1].x",
			data: []map[string]any{
				{"x": 1},
				{"x": 2},
				{"x": 3},
			},
			value: 5, expect: `[{"x":1},{"x":5},{"x":3}]`},
		{path: "[*].x",
			data: []map[string]any{
				{"x": 1},
				{"x": 2},
				{"x": 3},
			},
			value: 5, expect: `[{"x":5},{"x":5},{"x":5}]`},
		{path: "[1,'a'].x",
			data: []*Any{{X: 1}, {X: 2}, {X: 3}}, value: 5,
			expect: `[{"^":"Any","x":1},{"^":"Any","x":5},{"^":"Any","x":3}]`},
		{path: "[1,'a'].x",
			data: []map[string]any{
				{"x": 1},
				{"x": 2},
				{"x": 3},
			},
			value: 5, expect: `[{"x":1},{"x":5},{"x":3}]`},
		{path: "[1,'a'].x",
			data: []any{&Any{X: 1}, &Any{X: 2}, &Any{X: 3}}, value: 5,
			expect: `[{"^":"Any","x":1},{"^":"Any","x":5},{"^":"Any","x":3}]`},
		{path: "[1,'a'].x",
			data:  map[string]any{"a": &Any{X: 1}, "b": &Any{X: 2}, "c": &Any{X: 3}},
			value: 5, expect: `{"a":{"^":"Any","x":5},"b":{"^":"Any","x":2},"c":{"^":"Any","x":3}}`},
		{path: "[1,'x'].x", data: &Any{X: &Any{X: 1}}, value: 5, expect: `{"^":"Any","x":{"^":"Any","x":5}}`},
		{path: "[1,'x'].x", data: &Any{X: map[string]any{"x": 1}}, value: 5, expect: `{"^":"Any","x":{"x":5}}`},

		{path: "[1].x", data: []any{&Any{X: 1}, &Any{X: 2}, &Any{X: 3}}, value: 5,
			expect: `[{"^":"Any","x":1},{"^":"Any","x":5},{"^":"Any","x":3}]`},
		{path: "[*].x", data: []any{&Any{X: 1}, &Any{X: 2}, &Any{X: 3}}, value: 5,
			expect: `[{"^":"Any","x":5},{"^":"Any","x":5},{"^":"Any","x":5}]`},
		{path: "$.*.x",
			data:  map[string]any{"a": &Any{X: 1}, "b": &Any{X: 2}, "c": &Any{X: 3}},
			value: 5, expect: `{"a":{"^":"Any","x":5},"b":{"^":"Any","x":5},"c":{"^":"Any","x":5}}`},
		{path: "..x", data: []any{&Any{X: 1}, &Any{X: 2}, &Any{X: 3}}, value: 5,
			expect: `[{"^":"Any","x":5},{"^":"Any","x":5},{"^":"Any","x":5}]`},
		{path: "..x",
			data:  map[string]any{"a": &Any{X: 1}, "b": &Any{X: 2}, "c": &Any{X: 3}},
			value: 5, expect: `{"a":{"^":"Any","x":5},"b":{"^":"Any","x":5},"c":{"^":"Any","x":5},"x":5}`},
		{path: "[:2:2].x",
			data: []map[string]any{
				{"x": 1},
				{"x": 2},
				{"x": 3},
			},
			value: 5, expect: `[{"x":5},{"x":2},{"x":3}]`},
		{path: "[:2:2].x",
			data: []*Any{{X: 1}, {X: 2}, {X: 3}}, value: 5,
			expect: `[{"^":"Any","x":5},{"^":"Any","x":2},{"^":"Any","x":3}]`},
		{path: "*.a", data: &Any{X: 1}, value: 3, expect: `{"^":"Any","x":1}`},
		{path: "[0,1].x", data: []int{1}, value: 3, expect: `[1]`},
		{path: "[0:1].x", data: []int{1}, value: 3, expect: `[1]`},
		{path: "['x','y'].a", data: &Any{X: 1}, value: 3, expect: `{"^":"Any","x":1}`},

		{path: "x.a", data: map[string]any{"x": func() {}}, value: 3, err: "can not follow a func() at 'x'"},
		{path: "x.a", data: &Any{X: 1}, value: 3, er