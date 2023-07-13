---
title: "The mistery of v := v"
date: 2023-07-13T12:23:30+02:00
draft: false
tags:
  - go
  - programming
  - gotchas
---

Somethings one can observe strange `v := v` in Go source code. And there is a reason behind it. Take this code from [Go FAQ](https://go.dev/doc/faq#closures_and_goroutines) for example:
``` go {linenos=table}
func f() {
	done := make(chan bool)
	values := []int{"a", "b", "c"}
	for _, v := range values {
		v := v  // <-- Yeap, right here
		go func() {
			fmt.Println(v)
			done <- true
		}()
	}
	for range values {
		<-done
	}
}
```
&nbsp;

Output of this code is predictable and contains all the values from `values` variable.:
```
c
b
a
```
But if we remove line 5, logic is broken:
```
c
c
c
```
This is simple case to remember, sure. Just use one of 'workarounds': assign value to inner variable (`v := v`), or use is an argument for the go routine (` go func(u ...) {...}(v)`).

But there are more subtle ways to shoot yourself. By tests for example. Unfortenately tests are code as well, and can be false positive (snippet from [Go Wiki](https://github.com/golang/go/wiki/LoopvarExperiment)):
``` go {linenos=table}
func TestFoo(t *testing.T) {
	testCases := []int{1, 2, 3, 4}
	for _, v := range testCases {
		t.Run("sub", func(t *testing.T) {
			t.Parallel()
			if v%2 != 0 {
				t.Fatal("odd v", v)
			}
		})
	}
}
```
And this test will pass :(. Sometimes odd numbers are even. And here only `v := v` appears to be applicable, as we can't easily redefine `t.Run()` for pass extra arguments.

References:
* [go.dev: FAQ - What happens with closures running as goroutines?](https://go.dev/doc/faq#closures_and_goroutines)
* [github: golang wiki LoopvarExperiment](https://github.com/golang/go/wiki/LoopvarExperiment)
* [github: golang issue 56010](https://github.com/golang/go/issues/56010)
* [github: golang issue 60078](https://github.com/golang/go/issues/60078)