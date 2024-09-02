package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello("Yrll")
	want := "Hello, Yrll"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
