package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Chris", "")
		want := "Hello, Chris"
		assertCorrectHelper(t, got, want)
	})
	t.Run("say 'Hello, World' when empty string is passed to Hello", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World"
		assertCorrectHelper(t, got, want)
	})
	t.Run("say hello to people in Spanish", func(t *testing.T) {
		got := Hello("Elodie", "Spanish")
		want := "Hola, Elodie"
		assertCorrectHelper(t, got, want)
	})
	t.Run("say hellp to people in French", func(t *testing.T) {
		got := Hello("Elodie", "French")
		want := "Bonjour, Elodie"
		assertCorrectHelper(t, got, want)
	})
}

func assertCorrectHelper(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
