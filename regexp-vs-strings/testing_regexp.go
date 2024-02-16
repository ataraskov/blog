package main

import (
	"regexp"
	"strings"
)

func WhatIsItRegexp(input string) string {
	isQuestion := regexp.MustCompile(`\?$`)
	isEmpty := regexp.MustCompile(`^\s*$`)
	switch {
	case isQuestion.MatchString(input):
		return question
	case isEmpty.MatchString(input):
		return empty
	default:
		return other
	}
}

func WhatIsItStrings(input string) string {
	switch {
	case strings.HasSuffix(input, "?"):
		return question
	case strings.TrimSpace(input) == "":
		return empty
	default:
		return other
	}
}
