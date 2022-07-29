package main

import (
	"fmt"
	"testing"
)

var (
	cases = []string{
		"lumberjacks",
		"pump",
	}
)

func TestFindIsogram(t *testing.T) {
	for _, test := range cases {

		t.Run(fmt.Sprintf("is %s an isogram", test), func(t *testing.T) {
			got, _ := FindIsogram(string(test))
			want := false

			if got != want {
				t.Errorf("got %v, want %v", got, want)
			}
		})
	}
}
