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

func TestTokenizerLoadErrR