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
	MaskByTag = byte(0x10)
	// MaskExact is the mask for Exact fields.
	MaskExact = byte(0x08) // exact key vs lowwer case first letter
	// MaskPretty is the mask for Pretty fields.
	MaskPretty = byte(0x04)
	// MaskNested is the mask for Nested fields.
	MaskNested = byte(0x02)
	// MaskSen is the mask for Sen fields.
	MaskSen = byte(0x01)
	// MaskSet is the mask for Set fields.
	MaskSet = byte(0x20)
	// MaskIndex is the mask for an index that has been set up.
	MaskIndex = byte(0x1f)
)

var (
	// DefaultOptions default options that can be set as desired.
	DefaultOptions = Options{
		InitSize:    256,
		SyntaxColor: Normal,
		KeyColor:    Blue,
		NullColor:   Red,
		BoolColor:   Yellow,
		NumberColor: Cyan,
		StringColor: Green,
		TimeColor:   Magenta,
		HTMLUnsafe:  true,
		WriteLimit:  1024,
	}

	// BrightOptions encoding options for color encoding.
	BrightOptions = Options{
		InitSize:    256,
		SyntaxColor: Normal,
		KeyColor:    BrightBlue,
		NullColor:   BrightRed,
		BoolColor:   BrightYellow,
		NumberColor: BrightCyan,
		StringColor: BrightGreen,
		TimeColor:   BrightMagenta,
		WriteLimit:  1024,
	}

	// GoOptions are the options closest to the go json package.
	GoOptions = Options{
		InitSize:     256,
		SyntaxColor:  Normal,
		KeyColor:     Blue,
		NullColor:    Red,
		BoolColor:    Yellow,
		NumberColor:  Cyan,
		StringColor:  Green,
		TimeColor:    Magenta,
		CreateKey:    "",
		FullTypePath: false,
		OmitNil:      false,
		OmitEmpty:    false,
		UseTags:      true,
		KeyExact:     true,
		NestEmbed:    false,
		BytesAs:      BytesAsBase64,
		WriteLimit:   1024,
	}

	// HTMLOptions defines color options for generating colored HTML. The
	// encoding is suitable for use in a <pre> element.
	HTMLOptions = Options{
		InitSize:    256,
		SyntaxColor: "<span>",
		KeyColor:    `<span style="color:#44f">`,
		NullColor:   `<span style="color:red">`,
		BoolColor:   `<span style="color:#a40">`,
		NumberColor: `<span style="color:#04a">`,
		StringColor: `<span style="color:green">`,
		TimeColor:   `<span style="color:#f0f">`,
		NoColor:     "</span>",
		HTMLUnsafe:  false,
		WriteLimit:  1024,
	}
)

//