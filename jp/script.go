// Copyright (c) 2020, Peter Ohler, All rights reserved.

package jp

import (
	"reflect"
	"regexp"
	"strconv"

	"github.com/ohler55/ojg"
	"github.com/ohler55/ojg/gen"
)

type nothing int

var (
	eq     = &op{prec: 3, code: '=', name: "==", cnt: 2}
	neq    = &op{prec: 3, code: 'n', name: "!=", cnt: 2}
	lt     = &op{prec: 3, code: '<', name: "<", cnt: 2}
	gt     = &op{prec: 3, code: '>', name: ">", cnt: 2}
	lte    = &op{prec: 3, code: 'l', name: "<=", cnt: 2}
	gte    = &op{prec: 3, code: 'g', name: ">=", cnt: 2}
	or     = &op{prec: 4, code: '|', name: "||", cnt: 2}
	and    = &op{prec: 4, code: '&', name: "&&", cnt: 2}
	not    = &op{prec: 0, code: '!', name: "!", cnt: 1}
	add    = &op{prec: 2, code: '+', name: "+", cnt: 2}
	sub    = &op{prec: 2, code: '-', name: "-", cnt: 2}
	mult   = &op{prec: 1, code: '*', name: "*", cnt: 2}
	divide = &op{prec: 1, code: '/', name: "/", cnt: 2}
	get    = &op{prec: 0, code: 'G', name: "get", cnt: 1}
	in     = &op{prec: 3, code: 'i', name: "in", cnt: 2}
	empty  = &op{prec: 3, code: 'e', name: "empty", cnt: 2}
	rx     = &op{prec: 0, code: '~', name: "~=", cnt: 2}
	rxa    = &op{prec: 0, code: '~', name: "=~", cnt: 2}
	has    = &op{prec: 3, code: 'h', name: "has", cnt: 2}
	exists = &op{prec: 3, code: 'x', name: "exists", cnt: 2}
	// functions
	length = &op{prec: 0, code: 'L', name: "length", cnt: 1}
	count  = &op{prec: 0, code: 'C', name: "count", cnt: 1, getLeft: true}
	match  = &op{prec: 0, 