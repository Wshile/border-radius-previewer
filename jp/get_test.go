// Copyright (c) 2020, Peter Ohler, All rights reserved.

package jp_test

import (
	"fmt"
	"sort"
	"testing"

	"github.com/ohler55/ojg/alt"
	"github.com/ohler55/ojg/gen"
	"github.com/ohler55/ojg/jp"
	"github.com/ohler55/ojg/oj"
	"github.com/ohler55/ojg/pretty"
	"github.com/ohler55/ojg/tt"
)

type getData struct {
	path   string
	data   any
	expect []any
}

type Sample struct {
	A int
	B string
}

type One struct {
	A int
}

type Any struct {
	X any
}

var (
	getTestData = []*getData{
		{path: "", expect: []any{}},
		{path: "$.a.*.b", expect: []any{112, 122, 132, 142}},
		{path: "@.b[1].c", expect: []any{223}},
		{path: "..[1].b", expect: []any{122, 222, 322, 422}},
		{path: "[-1]", expect: []any{3}, data: []any{0, 1, 2, 3}},
		{path: "[1,'a']['b',2]['c',3]", expect: []any{133}},
		{path: "a[1::2].a", expect: []any{121, 141}},
		{path: "a[?(@.a > 135)].b", expect: []any{142}},
		{path: "[?(@[1].a > 230)][1].b", expect: []any{322, 422}},
		{path: "[?(@ > 1)]", expect: []any{2, 3}, data: []any{1, 2, 3}},
		{path: "$[?(1==1)]", expect: []any{1, 2, 3}, data: []any{1, 2, 3}},
		{path: "$.*[*].a", expect: []any{111, 121, 131, 141, 211, 221, 231, 241, 311, 321, 331, 341, 411, 421, 431, 441}},
		{path: "$.a[*].y",
			expect: []any{2, 4},
			data: map[string]any{
				"a": []any{
					map[string]any{"x": 1, "y": 2, "z": 3},
					map[string]any{"x": 2, "y": 4, "z": 6},
				},
			},
		},
		{path: "$..x",
			expect: []any{map[string]any{"x": 2}, 1, 2, 3, 4},
			data: map[string]any{
				"o": map[string]any{
					"a": []any{
						map[string]any{"x": 1},
						map[string]any{
							"x": map[string]any{
								"x": 2,
							},
						},
					},
					"x": 3,
				},
				"x": 4,
			},
		},
		{path: "$..[1].x",
			expect: []any{42, 200, 500},
			data: map[string]any{
				"x": []any{0, 1},
				"y": []any{
					map[string]any{"x": 0},
					map[string]any{"x": 42},
				},
				"z": []any{
					[]any{
						map[string]any{"x": 100},
						map[string]any{"x": 200},
						map[string]any{"x": 300},
					},
					[]any{
						map[string]any{"x": 400},
						map[string]any{"x": 500},
						map[string]any{"x": 600},
					},
				},
			},
		},
		{path: "$['a-b']",
			expect: []any{1},
			data:   map[string]any{"a-b": 1, "c-d": 2},
		},
		{path: "$.a..x",
			expect: []any{1, 2, 3, 4, 5},
			data: map[string]any{
				"a": map[string]any{
					"b": []any{
						map[string]any{"x": 1, "y": true},
						map[string]any{"x": 2, "y": false},
						map[string]any{"x": 3, "y": true},
						map[string]any{"x": 4, "y": false},
					},
					"c": map[string]any{"x": 5, "y": nil},
				},
			},
		},
		{path: "a[2].*", expect: []any{131, 132, 133, 134}},
		{path: "[*]", expect: []any{1, 2, 3}, data: []any{1, 2, 3}},
		{path: "$", expect: []any{map[string]any{"x": 1}}, data: map[string]any{"x": 1}},
		{path: "@", expect: []any{map[string]any{"x": 1}}, data: map[string]any{"x": 1}},
		{path: "['x',-1]", expect: []any{3}, data: []any{1, 2, 3}},
		{path: "$[1:3]", expect: []any{2, 3}, data: []any{1, 2, 3, 4, 5, 6}},
		{path: "$[::0]", expect: []any{}, data: []any{1, 2, 3, 4, 5, 6}},
		{path: "$[10:]", expect: []any{}, data: []any{1, 2, 3, 4, 5, 6}},
		{path: "$[:-10:-1]", expect: []any{1}, data: []any{1, 2, 3, 4, 5, 6}},
		{path: "$[1:10]", expect: []any{2, 3, 4, 5, 6}, data: []any{1, 2, 3, 4, 5, 6}},
		{path: "$[-4:-4]", expect: []any{}, data: []any{1, 2, 3, 4, 5, 6}},
		{path: "$[-4:-3]", expect: []any{3}, data: []any{1, 2, 3, 4, 5, 6}},
		{path: "$[-4:2]", expect: []any{}, data: []any{1, 2, 3, 4, 5, 6}},
		{path: "$[-4:3]", expect: []any{3}, data: []any{1, 2, 3, 4, 5, 6}},
		{path: "$[:2]", expect: []any{1, 2}, data: []any{1, 2, 3, 4, 5, 6}},
		{path: "$[-4:]", expect: []any{1, 2, 3}, data: []any{1, 2, 3}},
		{path: "$[0:3:1]", expect: []any{1, 2, 3}, data: []any{1, 2, 3, 4, 5}},
		{path: "$[0:4:2]", expect: []any{1, 3}, data: []any{1, 2, 3, 4, 5}},
		{path: "[-4:-1:2]", expect: []any{3, 5}, data: []any{1, 2, 3, 4, 5, 6}},
		{path: "[-4:]", expect: []any{1, 2, 3}, data: []any{1, 2, 3}},
		{path: "[-1:1:-2]", expect: []any{4, 6}, data: []any{1, 2, 3, 4, 5, 6}},
		{path: "c[-1:1:-1].a", expect: []any{331, 341}},
		{path: "a[2]..", expect: []any{map[string]any{"a": 131, "b": 132, "c": 133, "d": 134}, 131, 132, 133, 134}},
		{path: "..", expect: []any{[]any{1, 2}, 1, 2}, data: []any{1, 2}},
		{path: "..a", expect: []any{}, data: []any{1, 2}},
		{path: "a..b", expect: []any{112, 122, 132, 142}},
		{path: "[1]", expect: []any{2}, data: []int{1, 2, 3}},
		{path: "[-1]", expect: []any{3}, data: []int{1, 2, 3}},
		{path: "[-1,'a']", expect: []any{3}, data: []int{1, 2, 3}},
		{path: "$[::]", expect: []any{1, 2, 3}, data: []int{1, 2, 3}},
		{path: "[-1,'a'].x",
			expect: []any{2},
			data: []any{
				map[string]any{"x": 1, "y": 2, "z": 3},
				map[string]any{"x": 2, "y": 4, "z": 6},
			},
		},
		{path: "$[1:3:]", expect: []any{2, 3}, data: []any{1, 2, 3, 4, 5}},
		{path: "$[01:03:01]", expect: []any{2, 3}, data: []any{1, 2, 3, 4, 5}},
		{path: "$[:]['x','y']",
			expect: []any{1, 2, 4, 5},
			data: []any{
				map[string]any{"x": 1, "y": 2, "z": 3},
				map[string]any{"x": 4, "y": 5, "z": 6},
			},
		},
		{path: "a.b", expect: []any{}, data: map[string]any{"a": nil}},
		{path: "*.*", expect: []any{}, data: map[string]any{"a": nil}},
		{path: "*.*", expect: []any{}, data: []any{nil}},
		{path: "[0][0]", expect: []any{}, data: []any{nil}},
		{path: "['a','b'].c", expect: []any{}, data: map[string]any{"a": nil}},
		{path: "[1:0:-1].c", expect: []any{}, data: []any{nil, nil}},
		{path: "[0:1][0]", expect: []any{}, data: []any{nil}},
	}
	getTestReflectData = []*getData{
		{path: "['a','b']", expect: []any{"sample", 3}, data: &Sample{A: 3, B: "sample"}},
		{path: "$.*", expect: []any{"sample", 3}, data: &Sample{A: 3, B: "sample"}},
		{path: "$.a", expect: []any{3}, data: &Sample{A: 3, B: "sample"}},
		{path: "x.a", expect: []any{3}, data: map[string]any{"x": &Sample{A: 3, B: "sample"}}},
		{path: "[0,'x'].a", expect: []any{3}, data: map[string]any{"x": &Sample{A: 3, B: "sample"}}},
		{path: "[0].a", expect: []any{3}, data: []any{&Sample{A: 3, B: "sample"}}},
		{path: "[*].*", expect: []any{"sample", 3}, data: []*Sample{{A: 3, B: "sample"}}},
		{path: "[*].a", expect: []any{3}, data: []any{&Sample{A: 3, B: "sample"}}},
		{path: "$.*.a", expect: []any{3}, data: map[string]any{"x": &Sample{A: 3, B: "sample"}}},
		{path: "$..a", expect: []any{3}, data: map[string]any{"x": &Sample{A: 3, B: "sample"}}},
		{path: "$..a", expect: []any{3}, data: []any{&Sample{A: 3, B: "sample"}}},
		{path: "$[1:2].a", expect: []any{2}, data: []any{&One{A: 1}, &One{A: 2}, &One{A: 3}}},
		{path: "$[2:1:-1].a", expect: []any{3}, data: []any{&One{A: 1}, &One{A: 2}, &One{A: 3}}},
		{path: "[0::2].a", expect: []any{1, 3}, data: []*One{{A: 1}, {A: 2}, {A: 3}}},
		{path: "[-1:0:-2].a", expect: []any{3}, data: []*One{{A: 1}, {A: 2}, {A: 3}}},
		{path: "[4:0:-2].a", expect: []any{}, data: []*One{{A: 1}, {A: 2}, {A: 3}}},
		{path: "$.*[0]", expect: []any{3}, data: &Any{X: []any{3}}},
		{path: "$[1:2]", expect: []any{2}, data: []int{1, 2, 3}},
		{path: "$[1:2][0]", expect: []any{gen.Int(2)},
			data: []gen.Array{{gen.Int(1)}, {gen.Int(2)}, {gen.Int(3)}}},
		{path: "$[-10:]", expect: []any{1, 2, 3}, data: []int{1, 2, 3}},
		{path: "$[1:-10:-1]", expect: []any{1, 2}, data: []int{1, 2, 3}},
		{path: "$[2:10]", expect: []any{3}, data: []int{1, 2, 3}},
		// filter with map
		{
			path:   "$.x[?(@.b=='sample1')].a",
			expect: []any{3},
			data:   map[string]any{"x": []any{map[string]any{"a": 3, "b": "sample1"}}},
		},
		{
			path:   "$.x[?(@.a==3)].b",
			expect: []any{"sample1"},
			data:   map[string]any{"x": []any{map[string]any{"a": 3, "b": "sample1"}}},
		},
		// filter with struct
		{
			path:   "$.x[?(@.b=='sample2')].a",
			expect: []any{3},
			data:   Any{X: []*Sample{{A: 3, B: "sample2"}}},
		},
		{
			path:   "$.x[?(@.a==4)].b",
			expect: []any{"sample2"},
			data:   Any{X: []*Sample{{A: 4, B: "sample2"}}},
		},
		{path: "$.*", expect: []any{}, data: &one},
		{path: "['a',-1]", expect: []any{3}, data: []any{1, 2, 3}},
		{path: "['a','b']", expect: []any{}, data: []any{1, 2, 3}},
		{path: "$.*.x", expect: []any{}, data: &Any{X: 5}},
		{path: "$.*.x", expect: []any{}, data: &Any{X: 5}},
		{path: "[0:1].z", expect: []any{}, data: []*Any{nil, {X: 5}}},
		{path: "[0:1].z", expect: []any{}, data: []int{1}},
	}
)

var (
	firstData1    = map[string]any{"a": []any{map[string]any{"b": 2}}}
	one           = &One{A: 3}
	firstTestData = []*getData{
		{path: "", expect: []any{nil}, data: map[string]any{"x": 1}},
		{path: "$", expect: []any{map[string]any{"x": 1}}, data: map[string]any{"x": 1}},
		{path: "@", expect: []any{map[string]any{"x": 1}}, da