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
	  [set $.asm.hello world]  // output is now {good: bad, hello: world}
	]

The functions available are:

	      !=: Returns true if any the argument are not equal. An alias is !==.

	       *: Returns the product of all arguments. All arguments must be
	          numbers. If any of the arguments are not a number an error is
	          raised.

	       +: Returns the sum of all arguments. All arguments must be numbers
	          or strings. If any argument is a string then the result will be
	          a string otherwise the result will be a number. If any of the
	          arguments are not a number or a string an error is raised.

	       -: Returns the difference of all arguments. All arguments must be
	          numbers. If any of the arguments are not a number an error is
	          raised.

	       /: Returns the quotient of all arguments. All arguments must be
	          numbers. If any of the arguments are not a number an error is
	          raised. If an attempt is made to divide by zero and error will
	          be raised.

	       <: Returns true if each argument is less than any subsequent
	          argument. An alias is lt.

	      <=: Returns true if each argument is less than or equal to any
	          subsequent argument. An alias is lte.

	      ==: Returns true if all the argument are equal. Aliases are eq, ==,
	          and equal.

	       >: Returns true if each argument is greater than any subsequent
	          argument. An alias is gt.

	      >=: Returns true if each argument is greater than or equal to any
	          subsequent argument. An alias is gte.

	     and: Returns true if all argument evaluate to true. Any arguments
	          that do not evaluate to a boolean or null (false) raise an error.

	  append: Appends the second argument to the first argument which must be
	          an array.

	  array?: Returns true if the single required argumement is an array
	          otherwise false is returned.

	     asm: Processes all arguments in order using the return of each as
	          input for the next.

	      at: Forms a path starting with @. The remaining 