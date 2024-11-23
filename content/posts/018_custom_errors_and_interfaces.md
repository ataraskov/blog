---
title: Custom errors and interfaces
date: 2023-11-26T16:47:42+01:00
draft: false
tags:
  - golang
  - programming
  - gotchas
---

Go error handling is almost a meme. Who hasn't seen a joke or two about the ocean of if statements for `err != nil`?

Let's consider a simple example:
``` go {linenos=table}
type CustomError struct{}

func (e CustomError) Error() string {
	return "custom error"
}

func returnError() error {
	var err error
	return err
}

func returnCustomError() error {
	var err CustomError
	return err
}

func main() {
	fmt.Println(
		"returnError()",
		returnError() == nil,
	)
	fmt.Println(
		"returnCustomError()",
		returnCustomError() == nil,
	)
}
```
&nbsp;

What would be the output? True in both cases, obviously. Otherwise, it will break our beloved `err != nil` flow.

Surprise, it's actually not:
``` bash {linenos=table}
returnError() true
returnCustomError() false
```
&nbsp;

[Go FAQ](https://go.dev/doc/faq#nil_error) says:
> Just keep in mind that if any concrete value has been stored in the interface, the interface will not be nil

So once we know the type, the interface is not nil anymore, as it holds type value internally.

References:
* [GO: FAQ / Why is my nil error value not equal to nil?](https://go.dev/doc/faq#nil_error)