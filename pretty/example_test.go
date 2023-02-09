// Copyright (c) 2021, Peter Ohler, All rights reserved.

package pretty_test

import (
	"fmt"
	"strings"

	"github.com/ohler55/ojg/pretty"
	"github.com/ohler55/ojg/sen"
)

func ExampleJSON() {
	val := sen.MustParse([]byte("[true false [3 2 1] {a:1 b:2 c:3 d:[x y z []]}]"))
	// Pretty JSON format with a edge of 80 characters and a max depth of 2 per line.
	s := pretty.JSON(val, 80.2)
	fmt.Println(s)

	// Output:
	// [
	//   true,
	//   false,
	//   [3, 2, 1],
	//   {
	//     "a": 1,
	//     "b": 2,
	//     "c": 3,
	//     "d": ["x", "y", "z", []]
	//   }
	// ]
}

func ExampleSE