// Copyright (c) 2021, Peter Ohler, All rights reserved.

package asm_test

import (
	"testing"

	"github.com/ohler55/ojg/asm"
	"github.com/ohler55/ojg/sen"
	"github.com/ohler55/ojg/tt"
)

func TestStringCheck(t *testing.T) {
	root := testPlan(t,
		`[
           [set $.asm.a [string? abc]]
           [set $.asm.b [string? 123]]
           [set $.asm.c [string? true]]
         ]`,
		"{src: []}",
	)
	tt.Equal(t, "{a:true b:false c:false}", sen.String(root["asm"], &sopt))
}

func TestStringCheckArgCount(t *testing.T) {
	p := asm.NewPlan([]any{
		[]any{"string?", 1, 2},
	})
	err := p.Execute(map[string]any{})
	tt.NotNil(t, err)
}

func TestStringConv(t *testing.T) {
	root := testPlan(t,
		`[
           [set $.asm.a [string 1]]
           [set $.asm.b [string 1 "0x%02x"]]
           [set $.asm.c [string 2.2]]
           [set $.asm.d [string 2.2 "%e"]]
           [set $.asm.e [string [