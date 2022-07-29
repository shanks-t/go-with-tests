// Calculate the hamming distance between two DNA strands
package main

import "testing"

func TestGetHammingDistance(t *testing.T) {

	s1 := "GTAC"
	s2 := "GGAC"

	got, err := GetHammingDistance(s1, s2)
	want := 1

	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
	if err != nil {
		t.Fatal(err)
	}

}
