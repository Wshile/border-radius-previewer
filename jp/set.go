// Copyright (c) 2020, Peter Ohler, All rights reserved.

package jp

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/ohler55/ojg/alt"
	"github.com/ohler55/ojg/gen"
)

type delFlagType struct{}

var delFlag = &delFlagType{}

// MustDel removes matching nodes and pinics on error.
func (x Expr) MustDel(data any) {
	if err := x.set(data, delFlag, "delete", false); err != nil {
		panic(err)
	}
}

// Del removes matching nodes.
func (x Expr) Del(data any) error {
	return x.set(data, delFlag, "delete", false)
}

// MustDelOne removes one matching node and pinics on error.
func (x Expr) MustDelOne(data any) {
	if err := x.set(data, delFlag, "delete", true); err != nil {
		panic(err)
	}
}

// DelOne removes at most one node.
func (x Expr) DelOne(data any) error {
	return x.set(data, delFlag, "delete", true)
}

// MustSet all matching child node values. If the path to the child does not
// exist array and map elements are added. Panics on error.
func (x Expr) MustSet(data, value any) {
	if err := x.set(data, value, "set", false); err != nil {
		panic(err)
	}
}

// Set all matching child node values. An error is returned if it is not
// possible. If the path to the child does not exist array and map elements
// are added.
func (x Expr) Set(data, value any) error {
	return x.set(data, value, "set", false)
}

// SetOne child node value. An error is returned if it is not possible. If the
// path to the child does not exist array and map elements are added.
func (x Expr) SetOne(data, value any) error {
	return x.set(data, value, "set", true)
}

// MustSetOne child node value. If the path to the child does not exist array
// and map elements are added. Panics on error.
func (x Expr) MustSetOne(data, value any) {
	if err := x.set(data, value, "set", true); err != nil {
		panic(err)
	}
}

func (x Expr) set(data, value any, fun string, one bool) error {
	if len(x) == 0 {
		return fmt.Errorf("can not %s with an empty expression", fun)
	}
	switch x[len(x)-1].(type) {
	case Root, At, Bracket, Descent, Slice, *Filter:
		ta := strings.Split(fmt.Sprintf("%T", x[len(x)-1]), ".")
		return fmt.Errorf("can not %s with an expression ending with a %s", fun, ta[len(ta)-1])
	}
	var v any
	var nv gen.Node
	_, isNode := data.(gen.Node)
	nodeValue, ok := value.(gen.Node)
	if isNode && !ok {
		if value != nil {
			if v = alt.Generify(value); v == nil {
				return fmt.Errorf("can not %s a %T in a %T", fun, value, data)
			}
			nodeValue, _ = v.(gen.Node)
		}
	}
	var prev any
	stack := make([]any, 0, 64)
	stack = append(stack, data)

	f := x[0]
	fi := fragIndex(0) // frag index
	stack = append(stack, fi)

	for 1 < len(stack) {
		prev = stack[len(stack)-2]
		if ii, up := prev.(fragIndex); up {
			stack[len(stack)-1] = nil
			stack = stack[:len(stack)-1]
			fi = ii & fragIndexMask
			f = x[fi]
			continue
		}
		stack[len(stack)-2] = stack[len(stack)-1]
		stack[len(stack)-1] = ni