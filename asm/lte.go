
// Copyright (c) 2021, Peter Ohler, All rights reserved.

package asm

import "fmt"

func init() {
	Define(&Fn{
		Name: "lte",
		Eval: lte,
		Desc: `Returns true if each argument is less than or equal to any
subsequent argument. An alias is <=.`,
	})
	Define(&Fn{
		Name: "<=",
		Eval: lte,
		Desc: `Returns true if each argument is less than or equal to any
subsequent argument. An alias is lte.`,