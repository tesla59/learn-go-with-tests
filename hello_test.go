package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Mike")
		want := "Hello, Mike"

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})
	t.Run("saying hello to a place", func(t *testing.T) {
		got := Hello("Paris")
		want := "Hello, Paris"

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})
	t.Run("saying hello to an empty string", func(t *testing.T) {
		got := Hello("")
		want := "Hello, World"

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})
}
