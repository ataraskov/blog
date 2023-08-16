---
title: "Pipes and passwords"
date: 2023-08-16T14:27:50+02:00
draft: false
tags:
  - go
  - programming
  - cli
---

How hard can it be to read a password in your next CLI project written in Go? Should be as straightforward as:
``` go {linenos=table}
func getPass(prompt string) (string, error) {
	stdin := int(os.Stdin.Fd())
	fmt.Print(prompt)
	pass, err := term.ReadPassword(stdin)
	if err != nil {
		return "", err
	}
    return pass, nil
}
```
&nbsp;

Nice, issue closed, commit is ready, we are done for today! But wait, what if we do not like to pass password as an argument (any one will see it as argument in top/ps/... output). Easy, just pass it via standard input, shouldn't we? Let's test quickly:
``` bash {linesos=table}
$ echo $secret|go run main.go
Password: panic: inappropriate ioctl for device

goroutine 1 [running]:
main.main()
        main.go:15 +0x154
exit status 2
```
&nbsp;

Oops. ```term.ReadPassword``` is not able to handle pipes. Pipes are just special files, so we can check for pipes via ```fs.ModeNamedPipe``` bit mask. And just read from it:
``` go {linenos=table}
func getPass() (string, error) {
	fi, err := os.Stdin.Stat()
    ...

    // check if we are dealing with a pipe
	if fi.Mode()&fs.ModeNamedPipe != 0 {
			reader := bufio.NewReader(os.Stdin)
	        pass, err := reader.ReadString('\n')
    }
    ...
}
```
&nbsp;

And smashing everything together. Plus a bonus of handling termination and interruption of password prompt (thanks to an old google groups [thread](https://groups.google.com/forum/#!topic/golang-nuts/kTVAbtee9UA)):
``` go {linenos=table}
// readFile reads password from a file descriptor
func readFile(fd *os.File) (string, error) {
	reader := bufio.NewReader(os.Stdin)
	pass, err := reader.ReadString('\n')
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(pass), nil
}

// readTerm reads password from a terminal
func readTerm(prompt string) (string, error) {
	stdin := int(os.Stdin.Fd())
	// Get the initial state of the terminal.
	initialTermState, err := term.GetState(stdin)
	if err != nil {
		return "", err
	}
	// Restore it in the event of an interrupt.
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGTERM, syscall.SIGINT)
	go func() {
		<-c
		_ = term.Restore(stdin, initialTermState)
		os.Exit(1)
	}()
	defer signal.Stop(c)

	fmt.Print(prompt)
	pass, err := term.ReadPassword(stdin)
	fmt.Println("")
	if err != nil {
		return "", err
	}

	return string(pass), nil
}

// getPassword reads password from terminal or pipe
func getPassword(prompt string) (string, error) {
	fi, err := os.Stdin.Stat()
	if err != nil {
		return "", err
	}
	fmt.Printf(">> Stdin mode: %v\n", fi.Mode())

	switch {
	case fi.Mode()&fs.ModeNamedPipe != 0:
		fmt.Println(">> Reading from a pipe")
		return readFile(os.Stdin)
	case fi.Mode()&fs.ModeDevice != 0:
		fmt.Println(">> Reading from a term")
		return readTerm(prompt)
	default:
        // this will break *nix Heredoc, but this is intentional
		return "", fmt.Errorf("unsupported input")
	}
}

func main() {
	pass, err := getPassword("Password: ")
	if err != nil {
		panic(err)
	}
	fmt.Printf("Password is %s\n", pass)
}
```
&nbsp;

This is not a production ready code, and it has some debug/demonstration prints in it, but something consice and functional enough to start with ;)

References:
* [Groups: Handle ^C interrupts in ReadPassword](https://groups.google.com/forum/#!topic/golang-nuts/kTVAbtee9UA)
* [GO: os.FileMode](https://pkg.go.dev/os#FileMode)
* [GO: golang.org/x/term.ReadPassword](https://pkg.go.dev/golang.org/x/term#Terminal.ReadPassword)