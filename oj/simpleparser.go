// Copyright (c) 2020, Peter Ohler, All rights reserved.

package oj

import "io"

// SimpleParser is the interface shared by the package parsers.
type SimpleParser interface {
	// Parse a string in to simple types. An error is