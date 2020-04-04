// Copyright (c) 2021, Peter Ohler, All rights reserved.

package asm

import (
	"fmt"
	"time"

	"github.com/ohler55/ojg/sen"
)

func init() {
	Define(&Fn{
		Name: "string?",
		Eval: stringCheck,
		Desc: `Returns true if the single required argumement is a string
otherwise false is returned.`,
	})
	Define(&Fn{
		Name: "string",
		Eval: stringConv,
		De