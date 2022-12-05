# A Journey building a fast JSON parser and full JSONPath, Oj for Go

I had a dream. I'd write a fast JSON parser, generic data, and a
JSONPath implementation and it would be beautiful, well organized, and
something to be admired. Well, reality kicked in and laughed at those
dreams. A Go JSON parser and tools could be high performance but to
get that performance compromises in beauty would have to be made. This
is a tale of journey that ended with a Parser that leaves the Go JSON
parser in the dust and resulted in some useful tools including a
comp