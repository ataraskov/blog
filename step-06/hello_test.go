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
