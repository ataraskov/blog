---
title: "Git commit in version output"
date: 2023-10-18T10:13:31+02:00
draft: false
tags:
  - go
  - programming
---

I don't know why, but version output is really important and touchy subject for me. It is really nice to have extra metadata, not just version number. See version output of the docker CLI for example:
``` bash {linenos=table}
$ docker version
Client: Docker Engine - Community
 Version:           24.0.6
 API version:       1.43
 Go version:        go1.20.7
 Git commit:        ed223bc
 Built:             Mon Sep  4 12:31:44 2023
 OS/Arch:           linux/amd64
 Context:           default
...
```
&nbsp;

Easy enough we can just hard code these values, and update them as part of commit/release:
``` go {linenos=table}
var (
        Version = "0.0.1"
        GoVer   = "go1.21.2"
        Commit  = "aabbccd"
        Built   = "2023-10-18 00:00:00"
        OsArch  = "linux/amd64"
)

func main() {
        fmt.Printf("Version:\t%s\n", Version)
        fmt.Printf("Go version:\t%s\n", GoVer)
        fmt.Printf("Commit:\t\t%s\n", Commit)
        fmt.Printf("Built:\t\t%s\n", Built)
        fmt.Printf("OS/Arch:\t%s\n", OsArch)
}
```
&nbsp;

But wait, Git commit is always changing, how can we capture current one without any changes to code? No problem, by using build flags and providing values for our variables during build time. Go's build tool can do that via `-ldflags` which passes extra arguments to the linker. Linker has [many](https://pkg.go.dev/cmd/link) arguments to play with, one of them is `-X`:
```
-X importpath.name=value
	Set the value of the string variable in importpath named name to value.
	This is only effective if the variable is declared in the source code either uninitialized
	or initialized to a constant string expression. -X will not work if the initializer makes
	a function call or refers to other variables.
	Note that before Go 1.5 this option took two separate arguments.
```
Let's use it:
``` bash {linenos=table}
$ go build \
    -o example.bin \
    -ldflags "
     -X 'main.Version=0.0.2'
     -X 'main.GoVer=go1.21.3'
     -X 'main.Commit=bbccdde'
     -X 'main.Built=2023-10-18 10:49:06'
     -X 'main.OsArch=windows/amd64'
    " main.go 
    
$ ./example.bin 
Version:	0.0.2
Go version:	go1.21.3
Commit:		bbccdde
Built:		2023-10-18 10:47:08
OS/Arch:	windows/amd64

```
&nbsp;

And some of that info is provided by Go's `runtime` already:
``` go {linenos=table}
var (
	Version = "0.0.1"
	Commit  = "aabbccd"
	Built   = "2023-10-18 00:00:00"
)

func main() {
	fmt.Printf("Version:\t%s\n", Version)
	fmt.Printf("Go version:\t%s\n", runtime.Version())
	fmt.Printf("Commit:\t\t%s\n", Commit)
	fmt.Printf("Built:\t\t%s\n", Built)
	fmt.Printf("OS/Arch:\t%s/%s\n", 
        runtime.GOOS, runtime.GOARCH,
    )
}
```
&nbsp;

And with the rise of CI/CD and DevOps practices, we are building our software inside pipelines, usually. So let's rely on default variables provided (by Bitbucket in the below example):
``` bash {linenos=table}
tag=$BITBUCKET_TAG       # combination of $GITHUB_REF_TYPE 
                         # and $GITHUB_REF_NAME in GitHub
commit=$BITBUCKET_COMMIT # $GITHUB_SHA in GitHub
date=$(date +%F\ %T)

$ go build \
    -o example.bin \
    -ldflags "
     -X 'main.Version=$tag'
     -X 'main.Commit=$commit'
     -X 'main.Built=$date'
    " main.go 
```
&nbsp;

The full list of predefined variables can be found [here](https://docs.github.com/en/actions/learn-github-actions/variables) for GitHub, and [here](https://support.atlassian.com/bitbucket-cloud/docs/variables-and-secrets/) for Bitbucket (or somewhere else for your CI/CD solution).

But what if you need to set variable in some nested package, naming can be tricky, so better to check symbols:
``` bash {linenos=table}
# compile binary
run build -o example.bin main.go
# search symbols
go tool nm example.bin | grep Version
```
&nbsp;

References:
* [Go Packages: cmd/go](https://pkg.go.dev/cmd/go#hdr-Compile_packages_and_dependencies)
* [Go Packages: cmd/link](https://pkg.go.dev/cmd/link)
* [Go Packages: cmd/nm](https://pkg.go.dev/cmd/nm)
