
// Copyright (c) 2020, Peter Ohler, All rights reserved.

package alt

import (
	"fmt"
	"reflect"
	"time"
	"unsafe"

	"github.com/ohler55/ojg/gen"
)

// Genericer is the interface for the Generic() function that converts types
// to generic types.
type Genericer interface {

	// Generic should return a Node that represents the object. Generally this
	// includes the use of a creation key consistent with call to the
	// reflection based Generic() function.
	Generic() gen.Node
}

// Generify converts a value into Node compliant data. A best effort is made
// to convert values that are not simple into generic Nodes.
func Generify(v any, options ...*Options) (n gen.Node) {
	opt := &DefaultOptions
	if 0 < len(options) {
		opt = options[0]
	}
	if v != nil {
		switch tv := v.(type) {
		case bool:
			n = gen.Bool(tv)