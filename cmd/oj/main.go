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
	flag.BoolVar(&lazy, "z", lazy, "lazy mode accepts Simple Encoding Notation (quotes and commas mostly optional)")
	flag.BoolVar(&senOut, "sen", senOut, "output in Simple Encoding Notation")
	flag.BoolVar(&tab, "t", tab, "indent with tabs")
	flag.Var(&exValue{}, "x", "extract path")
	flag.Var(&matchValue{}, "m", "match equation/script")
	flag.Var(&delValue{}, "d", "delete path")
	flag.BoolVar(&showVersion, "version", showVersion, "display version and exit")
	flag.StringVar(&planDef, "a", planDef, "assembly plan or plan file using @<plan>")
	flag.BoolVar(&showRoot, "r", showRoot, "print root if an assemble plan provided")
	flag.StringVar(&prettyOpt, "p", prettyOpt, `pretty print with the width, depth, and align as <width>.<max-depth>.<align>`)
	flag.BoolVar(&html, "html", html, "output colored output as HTML")
	flag.BoolVar(&safe, "safe", safe, "escape &, <, and > for HTML inclusion")
	flag.StringVar(&confFile, "f", confFile, "configuration file (see -help-config), - indicates no file")
	flag.BoolVar(&showFnDocs, "fn", showFnDocs, "describe assembly plan functions")
	flag.BoolVar(&showFnDocs, "help-fn", showFnDocs, "describe assembly plan functions")
	flag.BoolVar(&showFilterDocs, "help-filter", showFilterDocs, "describe filter operators like [?(@.x == 3)]")
	flag.BoolVar(&showConf, "help-config", showConf, "describe .oj-config.sen format")
	flag.BoolVar(&mongo, "mongo", mongo, "parse mongo Javascript output")
	flag.StringVar(&convName, "conv", convName, `apply converter before writing. Supported values are:
  nano - converts integers over 946684800000000000 (2000-01-01) to time
  rcf3339 - converts string in RFC3339 or RFC3339Nano to time
  mongo - converts mongo wrapped values e.g.,  {$numberLong: "123"} => 123
  <with-numbers> - if digits are included then time layout is assumed
  <other> - any other is taken to be a key in a map with a string or nano time
`)
}

func main() {
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, `
usage: %s [<options>] [@<extraction>]... [(<match>)]... [<json-file>]...

The default behavior it to write the JSON formatted according to the color
options and the indentation option. If no files are specified JSON input is
expected from stdin.

Filtering and extraction of elements is supported using JSONPath and the
scripting that is part of JSONPath filters.

Extraction paths can be provided either with the -x option or an argument
starting with a $ or @. A Expr.Get() is executed and all the results are
either written or wrapped with an array and written depending on the value of
the wrap option (-w).

  oj -x abc.def myfile.json "@.x[?(@.y > 1)]"

To filter JSON documents the match option (-m) is used. If a JSON document
matches at least one match option the JSON will be written. In addition to the
-m option an argument starting with a '(' is assumed to be a match script that
follows the oj.Script format.

  oj -m "(@.name == 'Pete')" myfile.json "(@.name == "Makie")"

An argument that starts with a { or [ marks the start of a JSON document that
is composed of the remaining argument concatenated together. That document is
then used as the input.

  oj -i 0 -z {a:1, b:two}
  => {"a":1,"b":"two"}

Elements can be deleted from the JSON using the -d option. Multiple
occurrences of -d are supported.

Oj can also be used to assemble new JSON output from input data. An assembly
plan that describes how to assemble the new JSON if specified by the -a
option. The -fn option will display the documentation for assembly.

Pretty mode output can be used with JSON or the -sen option. It indents
according to a defined width and maximum depth in a best effort approach. The
-p takes a pattern of <width>.<max-depth>.<align> where width and max-depth
are integers and align is a boolean.

`, filepath.Base(os.Args[0]))
		flag.PrintDefaults()
		fmt.Fprintln(os.Stderr)
	}
	flag.Parse() // get config file if specified
	if showVersion {
		fmt.Printf("oj %s\n", version)
		os.Exit(0)
	}
	if showConf {
		displayConf()
		os.Exit(0)
	}
	if showFnDocs {
		displayFnDocs()
		os.Exit(0)
	}
	if showFilterDocs {
		displayFilterDocs()
		os.Exit(0)
	}
	extracts = extracts[:0]
	matches = matches[:0]
	dels = dels[:0]
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "*-*-* %s\n", err)
		os.Exit(1)
	}
}

fu