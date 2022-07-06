package stringmethods

import "testing"

func TestDoesContain(t *testing.T) {
	contains := doesContain("hello", "lo")
	expected := true

	if contains != expected {
		t.Errorf("expected '%t' but got '%t'", expected, contains)
	}
}
