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
			expect: []any{1, 2, 3,