---
title: "Hello TDD"
date: 2023-07-30T16:54:37+02:00
draft: false
tags:
  - golang
  - programming
  - methodology
---

Sooner or later any developer faces a nasty bug, which is too hard to find. In most cases such bugs can be covered with a simple unit test, but who write tests before release?

Let's check how it can be approached in a different way, in a world of test-driven development. How "Hello World" will look like in TDD? Code is availble on github [repo](https://github.com/ataraskov/blog/tree/main/hello-tdd).

#### Exampe Project

##### Step 0: First things first, let's write a README
Why not sprinkle a bit of [RDD]({{< ref "005_TIL_RDD_exists.md" >}}) on our example project.

Content of `README.md`:
```md {linenos=table}
Hello prints a hello message for the given name
or world if no name provided.

Usage:

    hello [<name>]
```
&nbsp;

##### Step 1: Write a test first
Back to TDD, write a test before actual code. And here we are facing one of main benifits of the approach. How would one test and reuse code. In the "Hello World" example, for unit testing we should not print on the standard output, but rather return a value as string. Otherwise testing will became a bit too much envolved.

Content of `hello_test.go`:
```go {linenos=table}
package main

import "testing"

func TestHello(t *testing.T) {
        actual := Hello()
        expected := "Hello, World"

        if actual != expected {
                t.Fatalf("got: %s, want: %s", actual, expected)
        }
}
```
&nbsp;

Just writing tests is not enough, let's use them now:
``` bash {linenos=table}
$ go test
go: go.mod file not found in current directory or any parent directory
```
&nbsp;

Don't worry it's expected, we are in empty project right now.

##### Step 2: Make tests pass
We have a clear error message, `go.mod` is missing. Create one and rerun tests:
``` bash {linenos=table}
$ go mod init hello-tdd
$ go test
# hello-tdd [hello-tdd.test]
./hello_test.go:6:12: undefined: Hello
FAIL    hello-tdd [build failed]
```
&nbsp;

Test fails again, but we have some progress.

##### Step 3: Make tests pass (again)
Error message says us that there are no implementation of `Hello` function. This is time to create one. 

Content of `hello.go`:
``` go {linenos=table}
package main

func Hello() string {
        return "Hello, World"
}
```
&nbsp;

Run tests:
``` bash {linenos=table}
$ go test
PASS
ok      hello-tdd       0.001s
```
&nbsp;
Better now. At this point it would be a good idea to commit your changes to a repository, so we have a checkpoint. 
But we are not at the point defined by our `README.md` yet.

##### Step 4: Add a test for next piece of logic
We have covered one of possible use cases for our fancy program. But there is another one. We need to pass an argument to the function. Will start with test first (how else).

Content of `hello_test.go`:
``` go {linenos=table}
package main

import "testing"

func TestHello(t *testing.T) {
	tcs := []struct {
		input    string
		expected string
	}{
		{"", "Hello, World"},
		{"Bob", "Hello, Bob"},
	}

	for _, tc := range tcs {
		actual := Hello(tc.input)
		if actual != tc.expected {
			t.Fatalf(
				"got: %s, want: %s",
				actual,
				tc.expected,
			)
		}
	}
}
```
&nbsp;

Run tests:
``` bash {linenos=table}
$ go test
# hello-tdd [hello-tdd.test]
./hello_test.go:15:19: too many arguments in call to Hello
        have (string)
        want ()
FAIL    hello-tdd [build failed]
```
&nbsp;

##### Step 5: Make tests pass
Again error message is concise and to the point. Update `Hello` implementation to get an argument and change return value accordingly.

Content of `hello.go`:
``` go {linenos=table}
package main

func Hello(name string) string {
        if name == "" {
                name = "World"
        }
        return "Hello, " + name
}
```
&nbsp;

Run tests:
``` bash {linenos=table}
$ go test -cover
PASS
        hello-tdd       coverage: 100.0% of statements
ok      hello-tdd       0.001s
```
Nice!

##### Step 6: Build and Release
Reterning to our `README.md`, we have covered all what was promissed. It time to build now. Wait a moment, dont we need a `main` function as entrypoint? Yeap, will have one in a decicated `main.go` file.

Content of `main.go`:
``` go {linenos=table}
package main

import (
        "fmt"
        "os"
)

func main() {
        name := ""
        if len(os.Args) > 1 {
                name = os.Args[1]
        }
        fmt.Println(Hello(name))
}
```
&nbsp;

Main is printing to standard output. We'll test it manually for now, were are not that religious about TDD afterall. 
``` bash {linenos=table}
$ go build -o hello
$ ./hello
Hello, World
$ ./hello John
Hello, John
```
This is a success :).


#### Take aways
One can notice a pattern above, and there is one:
1) Write the test for specific functionality
2) Do minimal changes to the code to satisfy the test
3) Improve and/or refactor code
4) Rinse and Repeat...

Interested in test-driven development and in Go? I would recommend to check [Learn Go with Tests](https://quii.gitbook.io/). This project is an example of beautiful сonvergence of tech, writing, and gophers ;)

References:
* [Wikipedia:Test-driven development](https://en.wikipedia.org/wiki/Test-driven_development)
* [Learn Go with Tests](https://quii.gitbook.io/)
