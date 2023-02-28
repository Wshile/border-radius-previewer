// Copyright (c) 2021, Peter Ohler, All rights reserved.

package sen_test

import (
	"fmt"
	"io"
	"strings"
	"testing"

	"github.com/ohler55/ojg/oj"
	"github.com/ohler55/ojg/sen"
	"github.com/ohler55/ojg/tt"
)

type testHandler struct {
	buf []byte
}

func (h *testHandler) Null() {
	h.buf = append(h.buf, "null "...)
}

func (h *testHandler) Bool(v bool) {
	h.buf = append(h.buf, fmt.Sprintf("%t ", v)...)
}

func (h *testHandler) Int(v int64) {
	h.buf = append(h.buf, fmt.Sprintf("%d ", v)...)
}

func (h *testHandler) Float(v float64) {
	h.buf = append(h.buf, fmt.Sprintf("%g ", v)...)
}

func (h *testHandler) Number(v string) {
	h.buf = append(h.buf, fmt.Sprintf("%s ", v)...)
}

func (h *testHandler) String(v string) {
	h.buf = append(h.buf, fmt.Sprintf("%s ", v)...)
}

func (h *testHandler) ObjectStart() {
	h.buf = append(h.buf, '{')
	h.buf = append(h.buf, ' ')
}

func (h *testHandler) ObjectEnd() {
	h.buf = append(h.buf, '}')
	h.buf = append(h.buf, ' ')
}

func (h *testHandler) Key(v string) {
	h.buf = append(h.buf, fmt.Sprintf("%s: ", v)...)
}

func (h *testHandler) ArrayStart() {
	h.buf = append(h.buf, '[')
	h.buf = append(h.buf, ' ')
}

func (h *testHandler) ArrayEnd() {
	h.buf = append(h.buf, ']')
	h.buf = append(h.buf, ' ')
}

type tokeTest struct {
	src    string
	expect string
	err    string
}

func TestTokenizerParseBasic(t *testing.T) {
	toker := sen.Tokenizer{}
	h := testHandler{}
	src := `[true,null,123,12.3]{x:12345678901234567890}`
	err := toker.Parse([]byte(src), &h)
	tt.Nil(t, err)
	tt.Equal(t, "[ true null 123 12.3 ] { x: 12345678901234567890 } ", string(h.buf))

	h.buf = h.buf[:0]
	err = sen.Tokenize([]byte(src), &h)
	tt.Nil(t, err)
	tt.Equal(t, "[ true null 123 12.3 ] { x: 12345678901234567890 } ", string(h.buf))

	h.buf = h.buf[:0]
	toker.OnlyOne = true
	err = toker.Parse([]byte("[1, 2, 3]  "), &h)
	tt.Nil(t, err)
	tt.Equal(t, "[ 1 2 3 ] ", string(h.buf))

	err = toker.Parse([]byte("[1, 2, 3]  4"), &h)
	tt.NotNil(t, err)
}

func TestTokenizerLoad(t *testing.T) {
	toker := sen.Tokenizer{}
	h := testHandler{}
	err := toker.Load(strings.NewReader("\xef\xbb\xbf"+`[true,null,123,12.3]{x:3}`), &h)
	tt.Nil(t, err)
	tt.Equal(t, "[ true null 123 12.3 ] { x: 3 } ", string(h.buf))
}

func TestTokenizerLoadErrRead(t *testing.T) {
	h := oj.ZeroHandler{}
	r := tt.ShortReader{Max: 5, Content: []byte("[1, 2, 3, true, false]")}
	err := sen.TokenizeLoad(&r, &h)
	tt.NotNil(t, err)

	r = tt.ShortReader{Max: 5000, Content: []byte("[ 123" + strings.Repeat(",  123", 120) + "]")}
	err = sen.TokenizeLoad(&r, &h)
	tt.NotNil(t, err)
}

type eofReader int

func (r eofReader) Read(b []byte) (int, error) {
	b[0] = '['
	return 1, io.EOF
}

func TestTokenizerLoadEOF(t *testing.T) {
	h := oj.ZeroHandler{}
	toker := sen.Tokenizer{}
	err := toker.Load(eofReader(0), &h)
	tt.NotNil(t, err)

	err = toker.Load(eofReader(0), &h)
	tt.NotNil(t, err)
}

func TestTokenizerLoadMany(t *testing.T) {
	h := oj.ZeroHandler{}
	for i, s := range []string{
		// The read buffer is 4096 so force a buffer read in the middle of
		// reading a token.
		strings.Repeat(" ", 4094) + "null ",
		strings.Repeat(" ", 4094) + "true ",
		strings.Repeat(" ", 4094) + "false ",
		strings.Repeat(" ", 4094) + `{x:1}`,
		strings.Repeat(" ", 4095) + `"x"`,
		strings.Repeat(" ", 4095) + "xyz",
		strings.Repeat(" ", 4094) + "[xyz]",
		strings.Repeat(" ", 4094) + "[xyz[]]",
		strings.Repeat(" ", 4094) + "[xyz{}]",
		strings.Repeat(" ", 4092) + "{x:abc}",
		strings.Repeat(" ", 4094) + "abc// comment\n",
		strings.Repeat(" ", 4094) + "[abc\n  def]",
	} {
		toker := sen.Tokenizer{}
		err := toker.Load(strings.NewReader(s), &h)
		tt.Nil(t, err, i)
	}
}

func TestTokenizerMany(t *testing.T) {
	for i, d := range []tokeTest{
		{src: "null", expect: "null"},
		{src: "true", expect: "true"},
		{src: "false", expect: "false"},
		{src: "false \n ", expect: "false"},
		{src: "hello", expect: "hello"},
		{src: "hello ", expect: "hello"},
		{src: `"hello"`, expect: "hello"},
		{src: "[one two]", expect: "[ one two ]"},
		{src: "123", expect: "123"},
		{src: "-12.3", expect: "-12.3"},
		{src: "2e-7", expect: "2e-07"},
		{src: "-12.5e-2", expect: "-0.125"},
		{src: "0", expect: "0"},
		{src: "0\n ", expect: "0"},
		{src: "-12.3 ", expect: "-12.3"},
		{src: "-12.3\n", expect: "-12.3"},
		{src: "-12.3e-5", expect: "-0.000123"},
		{src: "12.3e+5 ", expect: "1.23e+06"},
		{src: "12.3e+5\n ", expect: "1.23