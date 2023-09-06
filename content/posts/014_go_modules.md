---
title: "Local modules in Go"
date: 2023-09-06T21:05:14+02:00
draft: false
tags:
  - go
  - programming
---

Most of the Go code/project examples are organized as publicly available modules on GitHub. Historically, Go played nicely with Git-enabled repositories. But it's not that obvious how to play around on a local machine without pushing any code to a remote repository. At least it wasn't for me when I started playing with Go. Let's review how it's done in [Go 1.13+](https://go.dev/blog/using-go-modules).

First of all, we need to create a directory and initialize a module. We will name our project `playground`, why not. The `go mod init ...` command will create a `go.mod` file with module name and Go version used:
``` bash {linenos=table}
mkdir playground
cd playground
go mod init playground
```
&nbsp;

Write some code. Create an internal package (ideally we should write a test first sure, will pretend it's done already). For example `pkg/greeter`:
``` bash {linenos=table}
mkdir -p  pkg/greeter
```
And create a file `pkg/greeter/greeter.go`:
``` go {linenos=table}
package greeter

func Hello() string {
	return "Hello!"
}
```
&nbsp;

Time to use our 'internal' package, by `main.go`:
``` go {linenos=table}
package main

import (
	"fmt"
	"playground/pkg/greeter"
)

func main() {
	fmt.Println(greeter.Hello())
}
```
Notice the import path for the `pkg/greeter`. The trick is to include the root package name in relative imports, which is different from how it's done in other languages.

Our files tree looks like this now, and Go compiler is happy.
``` bash {linenos=table}
$ tree
.
├── go.mod
├── main.go
└── pkg
    └── greeter
        └── greeter.go
```
``` bash {linenos=table}
$ go run main.go
Hello!
```
&nbsp;

Just don't forget to open your package's root directory in VS Code ;)

References:
* [GO: Using Go Modules](https://go.dev/blog/using-go-modules)
* [DigitalOcean: How to Use Go Modules](https://www.digitalocean.com/community/tutorials/how-to-use-go-modules)