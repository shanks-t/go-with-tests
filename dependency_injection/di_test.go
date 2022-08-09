package main

import (
	"bytes"
	"testing"
)

func TestGreet(t *testing.T) {
	buffer := bytes.Buffer{}
	Greet(&buffer, "Trey")

	got := buffer.String()
	want := "Hello, Trey"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
