---
title: Get the most out of programming exercises
date: 2023-05-31T14:26:23+02:00
draft: false
tabs:
  - programming
  - learn
  - go
---
Lets take a simple exercise to check if a word/string consists of uniq characters (`-` and ` ` can repeat though). This is actually one of tasks from exercism.org Go learning track.

Here how I would approach it:

##### Step 1: Make sure you have tests coverage.
Fortunately exercism.org has me covered here, by provided tests for all exercises.

##### Step 2: Implement good-enough solution of the problem.
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
Now it's time to make it a bit better. We have a working code already, lets make it a bit better. Lets do few adjustments:
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
Try to solve the same problem using another approach. Lets switch from if/else to switch, just to check how it goes:
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
Hm. Not much changed in previous iteration. Time to check standard library for inspiration. For example we can use `strings.ContainsRune`. In any case the most important is to check available tools once some time is spend on the task.

&nbsp;

This may not work well for production development, but perfect for my learning expirience at least ;)
