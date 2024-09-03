package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Yrll", "")
		want := "Hello, Yrll"

		assertCorrectMessage(t, got, want)
	})

	t.Run("say 'Hello, World' when an empty string is supplied", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World"

		assertCorrectMessage(t, got, want)
	})

	t.Run("Say Hello in spanish when a second language parameter is included", func(t *testing.T) {
		got := Hello("Yrll", "Spanish")
		want := "Hola, Yrll"

		assertCorrectMessage(t, got, want)
	})

	t.Run("Say Hello in French", func(t *testing.T) {
		got := Hello("Yrll", "French")
		want := "Bonjour, Yrll"
		assertCorrectMessage(t, got, want)
	})
}

func assertCorrectMessage(t *testing.T, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
