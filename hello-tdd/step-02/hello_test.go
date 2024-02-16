package main

import "testing"

func TestHello(t *testing.T) {
	actual := Hello()
	expected := "Hello, World"

	if actual != expected {
		t.Fatalf("got: %s, want: %s", actual, expected)
	}
}
