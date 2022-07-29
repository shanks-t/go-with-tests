package main

import (
	"fmt"
	"testing"
)

var (
	testCases = []struct {
		phrase string
		ans    bool
	}{
		{phrase: "lumberjacks", ans: true},
		{phrase: "pump", ans: false},
		{phrase: "ten-year-old", ans: false},
		{phrase: "six-year-old", ans: true},
		{phrase: "Capital-c", ans: false},
		{phrase: "Coldguy-c", ans: false},
		{phrase: "Coldguy c", ans: false},
		{phrase: "1-Coldguy 1", ans: false},
	}
)

func TestFindIsogram(t *testing.T) {
	for _, tc := range testCases {

		t.Run(fmt.Sprintf("is %s an isogram", tc.phrase), func(t *testing.T) {
			got := FindIsogram(string(tc.phrase))
			want := tc.ans

			if got != want {
				t.Errorf("for %s, got %v, want %v", tc.phrase, got, want)
			}
		})
	}
}
