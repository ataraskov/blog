---
title: Regexp is not free
date: 2023-06-16T12:28:34+02:00
draft: false
tags:
  - go
  - regexp
  - benchmark
---
Years ago I was writing `perl` scripts and regexp became a default string manipulation approach for me. I even had a book dedicated to regexp intricacies.
A decade and few languages later, regexp is still a tempting tool.
Going with `go`, one can assume that statically compiled language with regexp matcher pre-compiled, should be blazing fast. Isn't it?
Lets test. Testing framework is a first class citizen in `go` not just for fun, but for profit as well.

Will test on a quite simple schenario of checking if string is empty or end with a question mark. Full code can be found at [github.com/ataraskov](https://github.com/ataraskov/blog/tree/main/regexp-vs-strings).

Regexp function to test:
``` go {linenos=table}
func WhatIsItRegexp(input string) string {
	isQuestion := regexp.MustCompile(`\?$`)
	isEmpty := regexp.MustCompile(`^\s*$`)
	switch {
	case isQuestion.MatchString(input):
		return question
	case isEmpty.MatchString(input):
		return empty
	default:
		return other
	}
}
```
&nbsp;

Strings function to test:
``` go {linenos=table}
func WhatIsItStrings(input string) string {
  switch {
    case strings.HasSuffix(input, "?"):
		return question
	case strings.TrimSpace(input) == "":
		return empty
	default:
		return other
	}
}
```
&nbsp;

Benchmark results:
```
BenchmarkRegexp-16      120072   9792 ns/op   11030 B/op   177 allocs/op
BenchmarkStrings-16   71365939     16 ns/op       0 B/op     0 allocs/op
```
&nbsp;
The winner is obvious. Strings approach is more readable, faster, and cheaper. Though discarding regexp is not an option either. But going with regexp should be a last resort, IMHO ;)
