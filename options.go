// Copyright (c) 2020, Peter Ohler, All rights reserved.

package ojg

import (
	"fmt"
	"strconv"
	"time"
)

const (
	// Normal is the Normal ANSI encoding sequence.
	Normal = "\x1b[m"
	// Black is the Black ANSI encoding sequence.
	Black = "\x1b[30m"
	// Red is the Red ANSI encoding sequence.
	Red = "\x1b[31m"
	// Green is the Green ANSI encoding sequence.
	Green = "\x1b[32m"
	// Yellow is the Yellow ANSI encoding sequence.
	Yellow = "\x1b[33m"
	// Blue is the Blue ANSI encoding sequence.
	Blue = "\x1b[34m"
	// Magenta is the Magenta ANSI encoding sequence.
	Magenta = "\x1b[35m"
	// Cyan is the Cyan ANSI encoding sequence.
	Cyan = "\x1b[36m"
	// White is the White ANSI encoding sequence.
	White = "\x1b[37m"
	// Gray is the Gray ANSI encoding sequence.
	Gray = "\x1b[90m"
	// BrightRed is the BrightRed ANSI encoding sequence.
	BrightRed = "\x1b[91m"
	// BrightGreen is the BrightGreen ANSI encoding sequence.
	BrightGreen = "\x1b[92m"
	// BrightYellow is the BrightYellow ANSI encoding sequence.
	BrightYellow = "\x1b[93m"
	// BrightBlue is the BrightBlue ANSI encoding sequence.
	BrightBlue = "\x1b[94m"
	// BrightMagenta is the BrightMagenta ANSI encoding sequence.
	BrightMagenta = "\x1b[95m"
	// BrightCyan is the BrightCyan ANSI encoding sequence.
	BrightCyan = "\x1b[96m"
	// BrightWhite is the BrightWhite ANSI encoding sequence.
	BrightWhite = "\x1b[97m"

	// BytesAsString indicates []byte should be encoded as a string.
	BytesAsString = iota
	// BytesAsBase64 indicates []byte should be encoded as base64.
	BytesAsBase64
	// BytesAsArray indicates []byte should be encoded as an array if integers.
	BytesAsArray

	// MaskByTag is the mask for byTag fields.
	MaskByTag = b