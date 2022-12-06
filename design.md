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
a collection an