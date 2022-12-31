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
	match  = &op{prec: 0, code: 'M', name: "match", cnt: 2}
	search = &op{prec: 0, code: 'S', name: "search", cnt: 2}

	opMap = map[string]*op{
		eq.name:     eq,
		neq.name:    neq,
		lt.name:     lt,
		gt.name:     gt,
		lte.name:    lte,
		gte.name:    gte,
		or.name:     or,
		and.name:    and,
		not.name:    not,
		add.name:    add,
		sub.name:    sub,
		mult.name:   mult,
		divide.name: divide,
		in.name:     in,
		empty.name:  empty,
		has.name:    has,
		exists.name: exists,
		rx.name:     rx,
		rxa.name:    rx,

		length.name: length,
		count.name:  count,
		match.name:  match,
		search.name: search,
	}
	// Nothing can be used in scripts to indicate no value as in a script such
	// as [?(@.x == Nothing)] this indicates there was no value as @.x. It is
	// the same as [?(@.x has false)] or [?(@.x exists false)].
	Nothing = nothing(0)
)

type op struct {
	name     string
	prec     byte
	cnt      byte
	code     byte
	getLeft  bool
	getRight bool
}

type precBuf struct {
	prec byte
	buf  []byte
}

// Script represents JSON Path script used in filters as well.
type Script struct {
	template []any
}

// NewScript parses the string argument and returns a script or an error.
func NewScript(str string) (s *Script, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = ojg.NewError(r)
		}
	}()
	s = MustNewScript(str)
	return
}

// MustNewScript parses the string argument and returns a script or an error.
func MustNewScript(str string) (s *Script) {
	p := &parser{buf: []byte(str)}
	if 0 < len(p.buf) && p.buf[0] == '(' {
		p.pos = 1
	}
	eq := p.readEquation()

	return eq.Script()
}

// Append a string representation of the fragment to the buffer and then
// return the expanded buffer.
func (s *Script) Append(buf []byte) []byte {
	buf = append(buf, '(')
	if 0 < len(s.template) {
		bstack := make([]any, len(s.template))
		copy(bstack, s.template)

		for i := len(bstack) - 1; 0 <= i; i-- {
			o, _ := bstack[i].(*op)
			if o == nil {
				continue
			}
			var (
				left  any
				right any
			)
			if 1 < len(bstack)-i {
				left = bstack[i+1]
			}
			if 2 < len(bstack)-i {
				right = bstack[i+2]
			}
			bstack[i] = s.appendOp(o, left, right)
			if i+int(o.cnt)+1 <= len(bstack) {
				copy(bstack[i+1:], bstack[i+int(o.cnt)+1:])
			}
		}
		if pb, _ := bstack[0].(*precBuf); pb != nil {
			buf = append(buf, pb.buf...)
		}
	}
	buf = append(buf, ')')

	return buf
}

// String representation of the script.
func (s *Script) String() string {
	return string(s.Append([]byte{}))
}

// Match returns true if the script returns true when evaluated against the
// data argument.
func (s *Script) Match(data any) bool {
	stack := []any{}
	if node, ok := data.(gen.Node); ok {
		stack, _ = s.EvalWithRoot(stack, gen.Array{node}, data).([]any)
	} else {
		stack, _ = s.EvalWithRoot(stack, []any{data}, data).([]any)
	}
	return 0 < len(stack)
}

// Eval is primarily used by the Expr parser but is public for testing.
func (s *Script) Eval(stack any, data any) any {
	return s.EvalWithRoot(stack, data, nil)
}

// EvalWithRoot is primarily used by the Expr parser but is public for testing.
func (s *Script) EvalWithRoot(stack any, data, root any) any {
	// Checking the type each iteration adds 2.5% but allows code not to be
	// duplicated and not to call a separate function. Using just one more
	// function call for each iteration adds 6.5%.
	var dlen int
	switch td := data.(type) {
	case []any:
		dlen = len(td)
	case gen.Array:
		dlen = len(td)
	case map[string]any:
		dlen = len(td)
		da := make([]any, 0, dlen)
		for _, v := range td {
			da = append(da, v)
		}
		data = da
	case gen.Object:
		dlen = len(td)
		da := make(gen.Array, 0, dlen)
		for _, v := range td {
			da = append(da, v)
		}
		data = da
	default:
		rv := reflect.ValueOf(td)
		if rv.Kind() != reflect.Slice && rv.Kind() != reflect.Array {
			return stack
		}
		dlen = rv.Len()
		da := make([]any, 0, dlen)
		for i := 0; i < dlen; i++ {
			da = append(da, rv.Index(i).Interface())
		}
		data = da
	}
	sstack := make([]any, len(s.template))
	var v any
	for vi := dlen - 1; 0 <= vi; vi-- {
		switch td := data.(type) {
		case []any:
			v = td[vi]
		case gen.Array:
			v = td[vi]
		}
		// Eval script for each member of the list.
		copy(sstack, s.template)
		