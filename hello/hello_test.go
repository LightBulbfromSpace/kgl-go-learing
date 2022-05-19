package main

import "testing"

func TestHello(t *testing.T) {
	assertCorrectMessage := func(t testing.TB, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}
	t.Run("saying hello to people in eng", func(t *testing.T) {
		got := Hello("Chris", "English")
		want := "Hello, Chris"
		assertCorrectMessage(t, got, want)
	})
	t.Run("empty string defaults to 'World' in eng", func(t *testing.T) {
		got := Hello("", "English")
		want := "Hello, World"
		assertCorrectMessage(t, got, want)
	})
	t.Run("saying hello to people in rus", func(t *testing.T) {
		got := Hello("Артем", "Russian")
		want := "Привет, Артем"
		assertCorrectMessage(t, got, want)
	})
	t.Run("empty string defaults to 'World' in rus", func(t *testing.T) {
		got := Hello("", "Russian")
		want := "Привет, Мир"
		assertCorrectMessage(t, got, want)
	})
	t.Run("saying hello to people in ger", func(t *testing.T) {
		got := Hello("Tom", "German")
		want := "Hallo, Tom"
		assertCorrectMessage(t, got, want)
	})
	t.Run("empty string defaults to 'World' in ger", func(t *testing.T) {
		got := Hello("", "German")
		want := "Hallo, Welt"
		assertCorrectMessage(t, got, want)
	})
}
