// Copyright (c) 2021, Peter Ohler, All rights reserved.

package asm_test

import (
	"testing"

	"github.com/ohler55/ojg/asm"
	"github.com/ohler55/ojg/sen"
	"github.com/ohler55/ojg/tt"
)

func TestOrTrue(t *testing.T) {
	root := testPlan(t,
		`[
           [set $.asm [or false "$.src[0]"]]
         ]`,
		"{src: [true false]}",
	)
	tt.Equal(t, "true",