---
title: "Don't lose an error during panic"
date: 2023-12-25T10:47:10+01:00
draft: false
tags:
  - golang
  - panic
---

It can be useful to print the error message returned by panic (if one is returned). That should be simple enough. Here were are not concerned that panic can return any, just to keep it simple.

``` go {linenos=table}
func job() {
	fmt.Println("in job")
	panic(errors.New("error in job"))
}

func executor() error {
	var err error
	fmt.Println("in executor")
	defer func() {
		if r := recover(); r != nil {
			err = r.(error)
		}
	}()
	job()

	return err
}

func main() {
	fmt.Println("in main")
	err := executor()
	fmt.Printf("error is %v", err)
}
```
&nbsp;

But above example reports error to be `nil`. Which is not true.
The 'problem' is in the `defer` statement, actually. Defer behavior, as described in the Go Blog [post](https://go.dev/blog/defer-panic-and-recover):
>    1. A deferred function’s arguments are evaluated when the defer statement is evaluated.
>    2. Deferred function calls are executed in Last In First Out order after the surrounding function returns.
>    3. Deferred functions may read and assign to the returning function’s named return values.

In our case, we have to use a named return value, to actually return err value from the function:
``` go {linenos=table}
func executor() (err error) {
	fmt.Println("in executor")
	defer func() {
		if r := recover(); r != nil {
			err = r.(error)
		}
	}()
	job()

	return err
}
```
&nbsp;

Even in panic, one needs to keep a good posture ;).

References:
* [GO: Defer, Panic, and Recover](https://go.dev/blog/defer-panic-and-recover)
