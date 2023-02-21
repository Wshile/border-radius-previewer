// Copyright (c) 2020, Peter Ohler, All rights reserved.

package sen_test

import (
	"fmt"
	"strings"
	"testing"
	"testing/iotest"

	"github.com/ohler55/ojg/sen"
	"github.com/ohler55/ojg/tt"
)

const callbackSEN = `
1
[2]
{x:3}
true false 123`

const tokenSEN = `
abc
def`

type rdata struct {
	src    string
	expect string
	value  any
}

func TestParserParseString(t *testing.T) {
	for i, d := range []rdata{
		{src: "null", value: nil},
		{src: "true", value: true},
		{src: "false", value: false},
		{src: "false \n ", value: false},
		{src: "hello", value: "hello"},
		{src: "hello ", value: "hello"},
		{src: `"hello"`, value: "hello"},
		{src: `'ab"cd'`, value: `ab"cd`},
		{src: `"ab'cd"`, value: `ab'cd`},
		{src: "[one two]", value: []any{"one", "two"}},
		{src: "123", value: 123},
		{src: "-12.3", value: -12.3},
		{src: "2e-7", value: 2e-7},
		{src: "-12.5e-2", value: -0.125},
		{src: "0", value: 0},
		{src: "0\n ", value: 0},
		{src: "-12.3 ", value: -12.3},
		{src: "-12.3\n", value: -12.3},
		{src: "-12.3e-5", value: -12.3e-5},
		{src: "12.3e+5 ", value: 12.3e+5},
		{src: "12.3e+5\n ", value: 12.3e+5},
		{src: "12.3e+05\n ", value: 12.3e+5},
		{src: "12.3e-05\n ", value: 12.3e-5},
		{src: `12345678901234567890`, value: "12345678901234567890"},
		{src: `9223372036854775807`, value: "9223372036854775807"},   // max int
		{src: `9223372036854775808`, value: "9223372036854775808"},   // max int + 1
		{src: `-9223372036854775807`, value: -9223372036854775807},   // min int
		{src: `-9223372036854775808`, value: "-9223372036854775808"}, // min int -1
		{src: `0.9223372036854775808`, value: "0.9223372036854775808"},
		{src: `-0.9223372036854775808`, value: "-0.9223372036854775808"},
		{src: `1.2e1025`, value: "1.2e1025"},
		{src: `-1.2e-1025`, value: "-1.2e-1025"},
		{src: `12345678901234567890.321e66`, value: "12345678901234567890.321e66"},
		{src: `321.12345678901234567890e66`, value: "321.12345678901234567890e66"},
		{src: `321.123e2345`, value: "321.123e2345"},

		{src: "\xef\xbb\xbf\"xyz\"", value: "xyz"},

		{src: "[]", value: []any{}},
		{src: "[0,\ntrue , false