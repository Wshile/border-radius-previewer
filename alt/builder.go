// Copyright (c) 2020, Peter Ohler, All rights reserved.

package alt

import (
	"fmt"

	"github.com/ohler55/ojg/gen"
)

var emptySlice = []any{}

// Builder is a basic type builder. It uses a stack model to build where maps
// (objects) and slices (arrays) add pushed on the stack and closed with a
// pop.
type Builder struct {
	stack  []any
	starts []int
}

// Reset the builder.
func (b *Builder) Reset() {
	if 0 < cap(b.stack) && 0 < len(b.stack) {
		b.stack = b.stack[:0]
		b.starts = b.starts[:0]
	} else {
		b.stack = make([]any, 0, 64)
		b.starts = make([]int, 0, 16)
	}
}

// Object pushs a map[string]any onto the stack. A key must be
// provided if the top of the stack is an object (map) and must not be
// provided if the op of the stack is an array or slice.
func (b *Builder) Object(key ...string) error {
	newObj := map[string]any{}
	if 0 < len(key) {
		if len(b.starts) == 0 || 0 <= b.starts[len(b.starts)-1] {
			return fmt.Errorf("can not use a key when pushing to an array")
		}
		if obj, _ := b.stack[len(b.stack)-1].(map[string]any); obj != nil {
			obj[key[0]] = newObj
		}
	} else if 0 < len(b.starts) && b.starts[len(b.starts)-1] < 0 {
		return fmt.Errorf("must have a key when pushing to an object")
	}
	b.starts = append(b.starts, -1)
	b.stack = append(b.stack, newObj)

	return nil
}

// Array pushs a []any onto the stack. A key must be provided if the
// top of the stack is an object (map) and must not be provided if the op of
// the stack is an array or slice.
func (b *Builder) Array(key ...string) error {
	if 0 < len(key) {
		if len(b.starts) == 0 || 0 <= b.starts[len(b.starts)-1] {
			return fmt.Errorf("can not use a key when pushing to an array")
		}
		b.stack = append(b.stack, gen.Key(key[0]))
	} else if 0 < len(b.starts) && b.starts[len(b.starts)-1] < 0 {
		return fmt.Errorf("must have a key when pushing to an object")
	}
	b.starts = append(b.starts, len(b.stack))
	b.stack = append(b.stack, emptySlice)

	return nil
}

// V