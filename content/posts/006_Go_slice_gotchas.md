---
title: "Go slice gotchas"
date: 2023-06-29T08:26:20+02:00
draft: false
tags:
 - go
 - slice
 - gotchas
---
Gotcha - a valid construct in a system, program or programming language that works as documented but is counter-intuitive and almost invites mistakes because it is both easy to invoke and unexpected or unreasonable in its outcome [[wikipedia](https://en.wikipedia.org/wiki/Gotcha_%28programming%29)].

Slice is actually a struct (holding info about pointer to the underlaying array, it's lenght, and capacity), but it behaves like a pointer. Because it is just a "view" of some memory area.

Sub-slice modification can cause modification of the original slice:
``` go {linenos=table}
	a := []int{1, 2}
	b := a[:1]
	b[0] = 9            /* what can go wrong? */
	fmt.Println("a", a) /* [9 2] */
	fmt.Println("b", b) /* [9] */
```
&nbsp;
``` go {linenos=table}
	a := []int{1, 2}
	b := a[:1]
	b = append(b, 9)    /* what can go wrong? */
	fmt.Println("a", a) /* [1 9] */
	fmt.Println("b", b) /* [1 9] */
```
&nbsp;

What about `append` to create a new slice? Nope that's a bad idea:
``` go {linenos=table}
	a := make([]int, 2)
	a = append(a, 1)
	b := append(a, 2)
	c := append(a, 9)   /* what can go wrong? */
	fmt.Println("a", a) /* [0 0 1]            */
	fmt.Println("b", b) /* [0 0 1 9] <- oops  */
	fmt.Println("c", c) /* [0 0 1 9]          */
```
&nbsp;

We can see common pattern here. Even in a modern and [memory-safe](https://en.wikipedia.org/wiki/Go_(programming_language)) language there are options to mess with memory and pointers.

There is a good example of another gotcha with slices, realated to garbage collector. Check it out [here](https://go.dev/blog/slices-intro). In short, when you create a slice, in many cases it would be benificial and safer to create it as a copy:
``` go {linenos=table}
    src := []int{1,2}
    dst := make([]int, len(src))
    copy(dst, src)
```
&nbsp;

Just don't overdo copying, as memory is not infinite unfortunately ;)

References:
* [Go Slices: usage and internals](https://go.dev/blog/slices-intro)
* [CommonMistakes - golang/go Wiki](https://github.com/golang/go/wiki/CommonMistakes)
