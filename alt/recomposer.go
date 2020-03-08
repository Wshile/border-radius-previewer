// Copyright (c) 2020, Peter Ohler, All rights reserved.

package alt

import (
	"encoding/json"
	"fmt"
	"math"
	"reflect"
	"strconv"
	"strings"
	"time"

	"github.com/ohler55/ojg"
	"github.com/ohler55/ojg/gen"
)

// DefaultRecomposer provides a shared Recomposer. Note that this should not
// be shared across go routines unless all types that will be used are
// registered first. That can be done explicitly or with a warm up run.
var DefaultRecomposer = Recomposer{
	composers: map[string]*composer{},
}

// RecomposeFunc should build an object from data in a map returning the
// recomposed object or an error.
type RecomposeFunc func(map[string]any) (any, error)

// RecomposeAnyFunc should build an object from data in an any
// returning the recomposed object or an error.
type RecomposeAnyFunc func(any) (any, error)

// Recomposer is used to recompose simple data into stru