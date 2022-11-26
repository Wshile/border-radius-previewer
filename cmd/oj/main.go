// Copyright (c) 2020, Peter Ohler, All rights reserved.

package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/ohler55/ojg"
	"github.com/ohler55/ojg/alt"
	"github.com/ohler55/ojg/asm"
	"github.com/ohler55/ojg/jp"
	"github.com/ohler55/ojg/oj"
	"github.com/ohler55/ojg/pretty"
	"github.com/ohler55/ojg/sen"
)

const version = "1.17.5"

var (
	indent         = 2
	color          = false
	bright         = false
	sortKeys       = false
	lazy           = false
	senOut         = false
	tab            = false
	showFnDocs     = false
	showFilterDocs = false
	showConf       = false
	safe           = false
	mongo          = false
	omit           = false

	// If true wrap extracts with an array.
	wrapExtract = false
	extracts    = []jp.Expr{}
	matches     = []*jp.Script{}
	dels        = []jp.Expr{}
	planDef     = ""
	showVersion bool
	plan        *asm.Plan
	root        = map[string]any{}
	showRoot    bool
	prettyOpt   = ""
	width       = 80
	maxDepth    = 3
	prettyOn    = false
	align       = false
	html        = false
	convName    = ""
	confFile    = ""

	conv    *alt.Converter
	options *ojg.Options
)

func init() {
	flag.IntVar(&indent, "i", indent, "indent")
	flag.BoolVar(&color, "c", color, "color")
	flag.BoolVar(&sortKeys, "s", sortKeys, "sort")
	flag.BoolVar(&bright, "b", bright, "bright color")
	flag.BoolVar(&omit, "o", omit, "omit nil and empty")
	flag.BoolVar(&wrapExtract, "w", wrapExtract, "wrap extracts in an array")
	flag.BoolVar(&lazy, "z", lazy, "lazy mode accepts