// Copyright (c) 2021, Peter Ohler, All rights reserved.

package pretty

import (
	"io"

	"github.com/ohler55/ojg"
)

// JSON encoded output. Arguments can be used to set the writer options. An
// int sets the width while a float64 is separated into a width as the integer
// portion of the float and the 10ths sets the maximum depth per line. A bool
// sets the align option and a *ojg.Options replaces the options portion of
// the writer.
func JSON(data any, args ...any) string {
	w := Writer{
		Options:  ojg.DefaultOptions,
		Width:    80,
		MaxDepth: 3,
		SEN:      false,
	}
	w.config(args)
	b, _ := w.encode(data)

	return 