// Copyright (c) 2020, Peter Ohler, All rights reserved.

package gen

import (
	"encoding/json"
	"math"
	"strconv"
)

// BigLimit is the limit before a number is converted into a Big
// instance. (9223372036854775807 / 10 = 922337203685477580)
const BigLimit = math.MaxInt64 / 10

// Number is used internally by parsers.
type Number struct {
	I          uint64
	Frac       uint64
	Div        uint64
	Exp        uint64
	Neg        bool
	NegExp     bool
	BigBuf     []byte
	ForceFloat bool
}

// Reset the number.
func (n *Number) Reset() {
	n.I = 0
	n.Frac = 0
	n.Div = 1
	n.Exp = 0
	n.Neg = false
	n.NegExp = false
	if 0 < len(n.BigBuf) {
		n.BigBuf = n.BigBuf[:0]
	}
}

// AddDigit to a number.
func (n *Number) AddDigit(b byte) {
	switch {
	case 0 < len(n.BigBuf):
		n.BigBuf = append(n.BigBuf, b)
	case n.I <= BigLimit