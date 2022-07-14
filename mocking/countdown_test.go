package main

import (
	"bytes"
	"testing"
)

func TestCountdown(t *testing.T) {
	buffer := &bytes.Buffer{}
	spySleeper := &SpySleeper{}

	Countdown(buffer, spySleeper)

	got := buffer.String()
	want := "3\n" +
		"2\n" +
		"1\n" +
		"Go!\n"

	if got != want {
		t.Errorf("got %q but want %q", got, want)
	}
	if spySleeper.Calls != 3 {
		t.Errorf("wrong number of calls to sleeper, want 3 got %d", spySleeper.Calls)
	}
}
