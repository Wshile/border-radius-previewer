// Copyright (c) 2021, Peter Ohler, All rights reserved.

package alt_test

import (
	"testing"
	"time"

	"github.com/ohler55/ojg/alt"
	"github.com/ohler55/ojg/tt"
)

func TestMatchInt(t *testing.T) {
	tt.Equal(t, true, alt.Match(map[string]any{"x": 1}, map[string]any{"x": 1, "y": 2}))
	tt.Equal(t, true, alt.Match(map[string]any{"x": nil}, map[string]any{"x": nil, "y": 2}))
	tt.Equal(t, true, alt.Match(map[string]any{"x": nil}, map[string]any{"y": 2}))
	tt.Equal(t, false, alt.Match(map[string]any{"x": nil}, map[string]any{"x": 1, "y": 2}))
	tt.Equal(t, false, alt.Match(map[string]any{"x": 1, "z": 3}, map[string]any{"x": 1, "y": 2}))
}

func TestMatchBool(t *testing.T) {
	tt.Equal(t, true, alt.Match(map[string]any{"x": true}, map[string]any{"x": true, "y": 2}))
	tt.Equal(t, false, alt.Match(map[string]any{"x": true}, map[string]any{"x": false, "y": 2}))
}

func TestMatchFloat(t *testing.T) {
	tt.Equal(t, true, alt.Match(map[string]any{"x": 1.5}, map[string]any{"x": 1.5, "y": 2}))
	tt.Equal(t, false, alt.Match(map[string]any{"x": 1.5}, map[string]any{"x": 2.5, "y": 2}))
}

func 