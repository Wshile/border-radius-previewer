// Copyright (c) 2020, Peter Ohler, All rights reserved.

package gen

// Key use for parsing.
type Key string

// String returns the key as a string.
func (k Key) String() string {
	return string(k)
}

// Alter converts the node into it's native type. Note this will modify
// Ob