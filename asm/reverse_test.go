// Copyright (c) 2021, Peter Ohler, All rights reserved.

package asm_test

import (
	"testing"

	"github.com/ohler55/ojg/asm"
	"github.com/ohler55/ojg/sen"
	"github.com/ohler55/ojg/tt"
)

func TestReverse(t *testing.T) {
	root := testPlan(t,
		`[
           [set $.asm.a [reverse [a b c]]]
           [set $.asm.b [reverse [1 b 3]]]
         ]`,
		"{src: []}",
	)
	tt.Equal(t, `{a:[c b a] b: