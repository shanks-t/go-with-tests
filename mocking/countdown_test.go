package mocking

import (
	"bytes"
	"testing"
)

type SpySleeper struct {
	Calls int
}

func (s *SpySleeper) Sleep() {
	s.Calls++
}

func TestCountdown(t *testing.T) {
	buffer := &bytes.Buffer{}
	SpySleeper := &SpySleeper{}

	Countdown(buffer, SpySleeper)

	got := buffer.String()
	want := `3
2
1
Go!`

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}

	if SpySleeper.Calls != 3 {
		t.Errorf("not enough calls to sleeper, want 3, got %d", SpySleeper.Calls)
	}
}
