// Copyright (c) 2020, Peter Ohler, All rights reserved.

package alt

import (
	"reflect"
	"strings"
)

type composer struct {
	fun     RecomposeFunc
	any     RecomposeAnyFunc
	short   string
	full    string
	rtype  