// Copyright (c) 2021, Peter Ohler, All rights reserved.

package ojg_test

import (
	"testing"

	"github.com/ohler55/ojg"
	"github.com/ohler55/ojg/tt"
)

func TestStringJSON(t *testing.T) {
	type Data struct {
		src      string
		expect   string
		htmlSafe bool
	}
	for i, td := range []*Data{
		{src: "abc", expect: `"abc"`},
		{src: "a\tbc", expect: `"a\tbc"`},
		{src: "a<b>c", expect: `"a<b>c"`},
		{src: "a<b>c", expect: `"a\u003cb\u003ec"`, htmlSafe: true},
		{src: "a 𝄢 note", expect: `"a 𝄢 note"`},
		{src: "a\u001ec", expect: `"a\u001ec"`},
		{src: "a\u2028b\u2029c", expect: `"a\u2028b\u2029c"`},
		{src: "abc\ufffd", expect: `"abc\ufffd"`},
	} {
		var buf []byte
		buf = ojg.Appe