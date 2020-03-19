// Copyright (c) 2021, Peter Ohler, All rights reserved.

/*
Package asm provides a means of building JSON or simple types using JSON
encoded scripts. The assembly scripts are encapsuled in the Plan type.

An assembly plan is described by a JSON document or a SEN document. The format
is much like LISP but with brackets instead of parenthesis. A plan is
evaluated by evaluating the plan function which is usually an 'asm'
function. The plan operates on a data map which is the root during
evaluation. The source data is in the $.src and the expected assembled output
should be in $.asm.

An example of a plan in SEN format is (the first asm is optional):

	[ asm
	  [set $.asm {good: bye}]  // set output to {good: bad}
	  [set $.asm.