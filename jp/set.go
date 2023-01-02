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
	retur