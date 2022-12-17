// Copyright (c) 2020, Peter Ohler, All rights reserved.

package jp_test

import (
	"regexp"
	"testing"

	"github.com/ohler55/ojg/jp"
	"github.com/ohler55/ojg/tt"
)

func TestEquation(t *testing.T) {
	eq := jp.Neq(jp.ConstInt(3), jp.ConstFloat(1.5))
	tt.Equal(t, "(3 != 1.5)", eq.String())

	eq = jp.Eq(jp.ConstBool(true), j