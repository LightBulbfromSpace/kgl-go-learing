package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello("User")
	want := "Hello, User"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
