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

// Options for writing data to JSON.
type Options struct {

	// Indent for the output.
	Indent int

	// Tab if true will indent using tabs and ignore the Indent member.
	Tab bool

	// Sort object members if true.
	Sort bool

	// OmitNil skips the writing of nil values in an object.
	OmitNil bool

	// OmitEmpty skips the writing of empty string, slices, maps, and zero
	// values although maps with all empty members will not be skipped on
	// writing but will be with alt.Decompose and alter.
	OmitEmpty bool

	// InitSize is the initial buffer size.
	InitSize int

	// WriteLimit is the size of the buffer that will trigger a write when
	// using a writer.
	WriteLimit int

	// TimeFormat defines how time is encoded. Options are to use a
	// time. layout string format such as time.RFC3339Nano, "second" for a
	// decimal representation, "nano" for a an integer. For decompose setting
	// to "time" will leave it unchanged.
	TimeFormat string

	// TimeWrap if not empty encoded time as an object with a single member. For
	// example if set to "@" then and TimeFormat is RFC3339Nano then the encoded
	// time will look like '{"@":"2020-04-12T16:34:04.123456789Z"}'
	TimeWrap string

	// TimeMap if true will encode time as a map with a create key and a
	// 'value' member formatted according to the TimeFormat options.
	TimeMap bool

	// CreateKey if set is the key to use when encoding objects that can later
	// be reconstituted with an Unmarshall call. This is only use when writing
	// simple types where one of the object in an array or map is not a
	// Simplifier. Reflection is used to encode all public members of the
	// object if possible. For example, is CreateKey is set to "type" this
	// might be the encoding.
	//
	//   { "type": "MyType", "a": 3, "b": true }
	//
	CreateKey string

	// NoReflect if true does not use reflection to encode an object. This is
	// only considered if the CreateKey is empty.
	NoReflect bool

	// FullTypePath if true includes the full type name and path when used
	// with the CreateKey.
	FullTypePath bool

	// Color if true will colorize the output.
	Color bool

	// SyntaxColor is the color for syntax in the JSON output.
	Syn