// Copyright (c) 2020, Peter Ohler, All rights reserved.

package oj_test

import (
	"bytes"
	"strings"
	"testing"
	"testing/iotest"

	"github.com/ohler55/ojg/oj"
	"github.com/ohler55/ojg/tt"
)

func TestValidatorValidateString(t *testing.T) {
	for i, d := range []data{
		{src: "null"},
		{src: "true"},
		{src: "false"},
		{src: "false \n "},
		{src: "123"},
		{src: "-321"},
		{src: "12.3"},
		{src: "0 "},
		{src: "12\n"},
		{src: "[]"},
		{src: "0\n"},
		{src: "-12.3 "},
		{src: "2e-7"},
		{src: "-12.3\n"},
		{src: "-12.3e-5"},
		{src: "12.3e+5 "},
		{src: "12.3e+5\n"},
		{src: `12345678901234567890`},
		{src: `9223372036854775807`},
		{src: `9223372036854775808`},
		{src: `-9223372036854775807`},
		{src: `-9223372036854775808`},
		{src: `0.9223372036854775808`},
		{src: `-0.9223372036854775808`},
		{src: `1.2e1025`},
		{src: `-1.2e-1025`},

		{src: "\xef\xbb\xbf\"xyz\"", value: "xyz"},

		{src: "[]"},
		{src: "[0,\ntrue , false,null]"},
		{src: `[0.1e3,"x",-1,{}]`},
		{src: "[1.2,0]"},
		{src: "[1.2e2,0.1]"},
		{src: "[1.2e2,0]"},
		{src: "[true]"},
		{src: "[true,false]"},
		{src: "[[]]"},
		{src: "[[true]]"},
		{src: `"x\t\n\"\b\f\r\u0041\\\/y"`},
		{src: `"x\u004a\u004Ay"`},
		{src: "\"bass \U0001D122\""},
		{src: `""`},
		{src: `[1,"a\tb"]`},
		{src: `{"a\tb":1}`},
		{src: `{"x":1,"a\tb":2}`},
		{src: "[0\n,3\n,5.0e2\n]"},

		{src: "{}"},
		{src: `{"abc":true}`},
		{src: "{\"z\":0,\n\"z2\":0}"},
		{src: `{"z":1.2,"z2":0}`},
		{src: `{"abc":{"def":3}}`},
		{src: `{"x":1.2e3,"y":true}`},
		{src: `{"abc": [{"x": {"y": [{"b": true}]},"z": 7}]}`},

		{src: "null {}"},

		{src: "{}}", expect: "unexpected object close at 1:3"},
		{src: "{ \n", 