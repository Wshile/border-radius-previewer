// Copyright (c) 2021, Peter Ohler, All rights reserved.

package asm_test

import (
	"sort"
	"testing"

	"github.com/ohler55/ojg/asm"
	"github.com/ohler55/ojg/sen"
	"github.com/ohler55/ojg/tt"
)

func TestGetall(t *testing.T) {
	root := testPlan(t,
		`[
           {one:1 two:2}
           [set $.at [getall "@.*"]]
           [set $.root [getall "$.src.*"]]
           [set $.arg [getall "@.*" {x:1 y:2}]]
         ]`,
		"{src: {a:1 b:2 c:3}}",
	)
	got, _ := root["at"].([]any)
	sort.Slice(got, func(i, j int) bool {
		a, _ := got[i].(in