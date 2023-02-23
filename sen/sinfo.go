// Copyright (c) 2021, Peter Ohler, All rights reserved.

package sen

import (
	"bytes"
	"reflect"
	"sort"
	"strings"
	"sync"
	"unsafe"
)

const (
	maskByTag  = byte(0x01)
	maskExact  = byte(0x02) // exact key vs lowwer case first letter
	maskNested = byte(0x04)
	maskPretty = byte(0x08)
	maskMax    = byte(0x10)
)

type sinfo struct {
	rt     reflect.Type
	fields [16][]*finfo
}

v