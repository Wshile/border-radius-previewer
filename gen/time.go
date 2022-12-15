// Copyright (c) 2020, Peter Ohler, All rights reserved.

package gen

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

// TimeFormat defines how time is encoded. Options are to use a time. layout
// string format such as time.RFC3339Nano, "second" for a decimal
// representation, "nano" for a an integer.
var TimeFormat = ""

// TimeWrap if not empty encoded time as an object with a single member. For
// example if set to "@" then and TimeFormat is RFC3339Nano then the encoded
// time will look like '{"@":"2020-04-12T16:34:04.123456789Z"}'
var TimeWrap = ""

// Time is a time.Time Node.
type Time time.Time

// String returns a string representation of the Node.
func (n Time) String() string {
