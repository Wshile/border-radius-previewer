// Copyright (c) 2020, Peter Ohler, All rights reserved.

package sen

const (
	skipChar     = 'a'
	skipNewline  = 'b'
	valSlash     = 'c'
	openParen    = 'd'
	valPlus      = 'e'
	valNeg       = 'f'
	val0         = 'g'
	valDigit     = 'h'
	valQuote     = 'i'
	tokenStart   = 'j'
	openArray    = 'k'
	openObject   = 'l'
	closeArray   = 'm'
	closeObject  = 'n'
	closeParen   = 'p'
	colonColon   = 'q'
	numSpc       = 'r'
	numNewline   = 's'
	numDot       = 't'
	tokenOk      = 'u'
	numFrac      = 'v'
	fracE        = 'w'
	expSign      = 'x'
	expDigit     = 'y'
	strQuote     = 'z'
	negDigit     = '-'
	strSlash     = 'A'
	escOk        = 'B'
	uOk          = 'E'
	tokenSpc     = 'G'
	tokenColon   = 'I'
	tokenNlColon = 'J'
	numDigit     = 'N'
	numZero      = 'O'
	strOk        = 'R'
	escU         = 'U'
	commentStart = 'K'
	commentEnd   = 'L'
	charErr      = '.'

	//   0123456789abcdef0123456789abcdef
	valueMap = "" +
		".........ab..a.................." + // 0x00
		"a.i.j..i.pjeafjcghhhhhhhhh..j.jj" + // 0x20
		"jjjjjjjjjjjjjjjjjjjjjjjjjjjk.mjj" + // 0x40
		".jjjjjjjjjjjjjjjjjjjjjjjjjjl.nj." + // 0x60
		"jjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjj" + // 0x80
		"jjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjj" + // 0xa0
		"jjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjj" + // 0xc0
		"jjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjv" //  0xe0
	//   0123456789abcdef0123456789abcdef
	tokenMap = "" +
		".........GJ..G.................." + // 0x00
		"G...u...dpuuGuucuuuuuuuuuuI.u.uu" + // 0x20
		"uuuuuuuuuuuuuuuuuuuuuuuuuuuk.muu" + // 0x40
		".uuuuuuuuuuuuuuuuuuuuuuuuuul.nu." + // 0x60
		"uuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuu" + // 0x80
		"uuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuu" + // 0xa0
		"uuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuu" + // 0xc0
		"uuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuut" //  0xe0
	//   0123456789abcdef0123456789abcdef
	colonMap = "" +
		".........ab..a.................." + // 0x00
		"a.........................q....." + // 0x20
		"................................" + // 0x40
		"................................" + // 0x60
		"................................" + // 0x80
		"................................" + // 0xa0
		"................................" + // 0xc0
		"................................" //   0xe0
	//   0123456789abcdef0123456789abcdef
	negMap = "" +
		"................................" + 