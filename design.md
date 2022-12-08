# A Journey building a fast JSON parser and full JSONPath, Oj for Go

I had a dream. I'd write a fast JSON parser, generic data, and a
JSONPath implementation and it would be beautiful, well organized, and
something to be admired. Well, reality kicked in and laughed at those
dreams. A Go JSON parser and tools could be high performance but to
get that performance compromises in beauty would have to be made. This
is a tale of journey that ended with a Parser that leaves the Go JSON
parser in the dust and resulted in some useful tools including a
complete and efficient JSONPath implementation.

In all fairness I did embark on with some previous experience. Having
written two JSON parser before. Both the Ruby
[Oj](https://github.com/ohler55/oj) and the C parser
[OjC](https://github.com/ohler55/ojc). Why not an
[OjG](https://github.com/ohler55/ojg) for go.

## Planning

Like any journey it starts with the planning. Yeah, I know, it's called
requirement gathering but casting it as planning a journey is more fun
and this was all about enjoying the discoveries on the journey. The
journey takes place in the land of OjG which stands for Oj for
Go. [Oj](https://github.com/ohler55/oj) or Optimized JSON being a
popular gem I wrote for Ruby.

First, JSON parsing and any frequently used operations such as
JSONPath evaluation had to be fast over everything else. With the
luxury of not having to follow the existing Go json package API the
API could be designed for the best performance.

The journey would visit several areas each with its own landscape and
different problems to solve.

### Generic Data

The first visit was to generic data. Not to be confused with the
proposed Go generics. Thats a completely different animal and has
nothing to do with whats being referred to as generic data here. In
building tools or packages for reuse the data acted on by those tools
needs to be navigable.

Reflection can be used but that gets a bit tricky when dealing with
private fields or field that can't be converted to something that can
say be written as a JSON element. Other options are often better.

Another approach is to use simple Go types such as `bool`, `int64`,
`[]any`, and other types that map directly on to JSON or some
other subset of all possible Go types. If too open, such as with
`[]any` it is still possible for the user to put unsupported
types into the data. Not to pick out any package specifically but it
is frustrating to see an argument type of `any` in an API and
then no documentation describing that the supported types are.

There is another approach though: Define a set of types that can be in
a collection and use those types. With this approach, the generic data
implementation has to support the basic JSON types of `null`,
`boolean`, `int64`, `float64`, `string`, array, and object. In
addition time should be supported. From experience in both JSON use in
Ruby and Go time has always been needed. Time is just too much a part
of any set of data to leave it out.

The generic data had to be type safe. It would not do to have an
element that could not be encoded as JSON in the data.

A frequent operation for generic data is to store that data into a
JSON database or similar. That meant converting to simple Go types of
`nil`, `bool`, `int64`, `float64`, `string`, `[]any`, and
`map[string]any` had to be fast.

Also planned for this part of the journey was methods on the types to
support getting, setting, and deleting elements using JSONPath. The
hope was to have an object based approach to the generic nodes so
something like the following could be used but keeping generic data,
JSONPath, and parsing in separate packages.

```golang
    var n gen.Node
    n = gen.Int(123)
    i, ok := n.AsInt()
```

Unfortunately that part of the journey had to be cancelled as the Go
travel guide refuses to let packages talk back and forth. Imports are
one way only. After trying to put all the code in one package it
eventually got unwieldy. Function names started being prefixed with
what should really have been package names so the object and method
approach was dropped. A change in API but the journey would continue.

### JSON Parser and Validator

The next stop was the parser and validator. After some consideration
it seemed like starting with the validator would be best way to become
familiar with the territory. The JSON parser and validator need not be
the same and each should be as performant as possible. The parsers
needed to support parsing into simple Go types as well as the generic
data types.

When parsing files that include millions or more JSON elements in
files that might be over 100GB a streaming parser is necessary. It
would be nice to share some code with both the streaming and string
parsers of course. It's easier to pack light when the areas are
similar.

The parser must also allow parsing into native Go types. Furthermore
interfaces must be supported even though Go unmarshalling does not
support interface fields. Many data types make use of interfaces
that limitation was not acceptable for the OjG parser. A different
approach to support interfaces was possible.

JSON documents of any non-trivial size, especially if hand-edited, are
likely to have errors at some point. Parse errors must identify where
in the document the error occurred.

### JSONPath

Saving the most interesting part of the trip for last, the JSONPath
implementation promised to have all sorts of interesting problems to
solve with descents, wildcards, and especially filters.

A JSONPath is used to extract elements from data. That part of the
implementation had to be fast. Parsing really didn't have to be fast
but it would be nice to have a way of building a JSONPath in a
performant manner even if it was not as convenient as parsing a
string.

The JSONPath implementation had to implement all the features
described by the [Goessner
article](https://goessner.net/articles/JsonPath). There are other
descriptions of JSONPath but the Goessner description is the most
referenced. Since the implementation is in Go the scripting feature
described could be left out as long as similar functionality could be
provided for array indexes relative to the length of the
array. Borrowing from Ruby, using negative indexes would provide that
functionality.

## The Journey

The journey unfolded as planned to a degree. There were some false
starts and revisits but eventually each destination was reached and
the journey completed.

### Generic Data (`gen` package)

What better way to make generic type fast than to just define generic
types from simple Go types and then add methods on those types? A
`gen.Int` is just an `int64` and a `gen.Array` is just a
`[]gen.Node`. With that approach there are no extra allocations.

```golang
type Node any
type Int int64
type Array []Node
```

Since generic arrays and objects restrict the type of the values in
each collection to `gen.Node` types the collections are assured to
contain only elements that can be encoded as JSON.

Methods on the `Node` could not be implemented without import loops so
the number of functions in the `Node` interface were limited. It was
clear a parser specific to the generic data type would be needed but
that would have to wait until the parser part of the journey was
completed. Then the generic data package could be revisited and the
parser explored.

Peeking at the future to the generic data parser revisit it was not
very interesting after the deep dive into the simple data parser. The
parser for generic types is a copy of the oj package parser but
instead of simple types being created instances that support the
`gen.Node` interface are created.

### Simple Parser (`oj` package)

Looking back its hard to say what was the most interesting part of the
journey, the parser or JSONPath. Each had their own unique set of
issues. The parser was the best place to start though as some valuable
lessons were learned about what to avoid and what to gravitate toward
in trying to achieve high performance Go code.

#### Validator

From the start I knew that a single pass parser would be more
efficient than building tokens and then making a second pass to decide
what the tokens means. At least that approach as worked well in the
past. I dived in and used a `readValue` function that branched
depending on the next character read. It worked but it was slower than
the target of being on par with the Go `json.Validate`. That was the
bar to clear. The first attempt was off by a lot. Of course a
benchmark was needed to verify that so the `cmd/benchmark` command was
started. Profiling didn't help much. It turned out since much of the
overhead was in the function call setup which isn't obvious when
profiling.

Not knowing at the time that function calls were so expensive but
anticipating that there was some overhead in function calls I moved
some of the