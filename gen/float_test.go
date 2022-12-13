// Copyright (c) 2020, Peter Ohler, All rights reserved.

package gen_test

import (
	"fmt"
	"testing"

	"github.com/ohler55/ojg/gen"
	"github.com/ohler55/ojg/tt"
)

func TestFloatString(t *testing.T) {
	b := gen.Float(12.34)

	tt.Equal(t, "12.34", b.String())
}

func TestFloatSimplify(t *testing.T)