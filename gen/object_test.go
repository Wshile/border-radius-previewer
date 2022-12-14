// Copyright (c) 2020, Peter Ohler, All rights reserved.

package gen_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/ohler55/ojg/gen"
	"github.com/ohler55/ojg/tt"
)

func TestObjectString(t *testing.T) {
	gen.Sort = true
	o := gen.Object{"a": gen.Int(3), "b": gen.Object{"c": gen.Int(5)}, "d": gen.Int(7), "n": nil}
	tt.Equal(t, `{"a":3,"b":{"c":5},"d":7,"n":null}`, o.String())

	gen.Sort = false
	o