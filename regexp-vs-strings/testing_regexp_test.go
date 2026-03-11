package main

import "testing"

var cases = []struct {
	desc  string
	input string
	want  string
}{
	{
		desc:  "A question",
		input: "Some text?",
		want:  question,
	},
	{
		desc:  "Empty string",
		input: "",
		want:  empty,
	},
	{
		desc:  "Something else",
		input: "Whatever here is not a question nor empty string",
		want:  other,
	},
}

func TestRegexp(t *testing.T) {
	for _, c := range cases {
		t.Run(c.desc, func(t *testing.T) {
			got := WhatIsItRegexp(c.input)
			if got != c.want {
				t.Errorf("got %q, want %q", got, c.want)
			}
		})
	}

}

func TestStrings(t *testing.T) {
	for _, c := range cases {
		t.Run(c.desc, func(t *testing.T) {
			got := WhatIsItStrings(c.input)
			if got != c.want {
				t.Errorf("got %q, want %q", got, c.want)
			}
		})
	}

}

func BenchmarkRegexp(b *testing.B) {
	if testing.Short() {
		b.Skip("skipping benchmark in short mode.")
	}
	for i := 0; i < b.N; i++ {
		for _, c := range cases {
			WhatIsItRegexp(c.input)
		}
	}
}

func BenchmarkStrings(b *testing.B) {
	if testing.Short() {
		b.Skip("skipping benchmark in short mode.")
	}
	for i := 0; i < b.N; i++ {
		for _, c := range cases {
			WhatIsItStrings(c.input)
		}
	}
}
