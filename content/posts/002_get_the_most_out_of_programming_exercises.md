---
title: Get the most out of programming exercises
date: 2023-06-01T09:00:00+02:00
draft: false
tags:
  - learn
  - go
  - programming
---
Programming exercises are rarely a fun activity. In most cases we do them to sharpen skills. But how to get the most out of that activity?

Take a simple exercise to check if a word/string consists of uniq characters (dash '`-`' and space '` `' can repeat though). This is actually one of the tasks from [exercism.org](https://exercism.org) Go learning track.

Here is how I would approach it (code snippets omit imports and other Go source code boilerplate for brevity):

##### Step 1: Make sure you have good test coverage.
Fortunately exercism.org has me covered here, by provided tests for all exercises.

##### Step 2: Implement a good-enough solution of the problem.
The main goal is to pass unit tests and commit changes. Code should not be ideal, just like one below:
``` go {linenos=table}
func f(word string) bool {
	letters := make(map[rune]bool)
	for _, l := range strings.ToLower(word) {
		if l == ' ' || l == '-' {
			continue
		}
		if _, exists := letters[l]; exists {
			return false
		}
		letters[l] = true
	}
	return true
}
```
&nbsp;

##### Step 3: Refine your solution. 
Now it's time to make it a bit better. We have a working code already, let's make it a bit better. Let's do a few adjustments:
* Rename variables to reflect usage
* Rely on default value of non-existing map elements (line 7)
``` go {linenos=table}
func f(word string) bool {
	seen := make(map[rune]bool)
	for _, l := range strings.ToLower(word) {
		if l == ' ' || l == '-' {
			continue
		}
		if seen[l] {
			return false
		}
		seen[l] = true
	}
	return true
}
```
&nbsp;

##### Step 4: Search for an alternative approach.
Try to solve the same problem using another approach. Let's switch from if/else to switch, just to check how it goes:
``` go {linenos=table}
func f(word string) bool {
	seen := make(map[rune]bool)
	for _, l := range strings.ToLower(word) {
		switch {
		case l == ' ' || l == '-':
			continue
		case seen[l]:
			return false
		default:
			seen[l] = true
		}
	}
	return true
}
```
&nbsp;

##### Step 5: Check what is available via standard library.
Hmm. Not much changed in the previous iteration. Time to check standard library for an inspiration. For example we can use `strings.ContainsRune`. In any case the most important is to check available tools once some time is spent on the task.

&nbsp;

This may not work well for production development, but it's perfect for my learning experience at least ;)
