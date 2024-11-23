---
title: "Semantic Naming"
date: 2024-02-13T10:16:08+01:00
draft: false
tags:
  - golang
  - programming
  - methodology
---
Reading literature is easier than writing it. That goes as a common knowledge, with few exceptions. One can assume that the same applies to code. Does it?
Authors have a goal to be read, and understood, more often than not. Some authors are trying hard to keep a reader entertained and engaged. Some say, "It's my style".

With code, the goal is to accomplish some task via computation or bytes manipulation, usually. Our code will not be reviewed by anyone from NY Times. A machine will be happy to execute it, as far as it's syntactically correct.
"Reading Code Is Harder Than Writing It" try to find an opposite thesis online.

We cannot change basic language syntax. But we name our variables, functions, etc.

> "Good naming is like a good joke. If you have to explain it, it’s not funny." — [Dave Cheney](https://twitter.com/davecheney/status/997155238929842176)

What can be done to improve readability?
* Read ~~classics~~ [std](https://pkg.go.dev/std) lib's source code
* Follow [naming conventions](https://go.dev/doc/effective_go#names)
* Be consistent

Let's pay attention to how we code, so our writing is not painful to read ;)

References:
* [Effective Go](https://go.dev/doc/effective_go#names)
* [Practical Go: Real world advice for writing maintainable Go programs](https://dave.cheney.net/practical-go/presentations/qcon-china.html)
