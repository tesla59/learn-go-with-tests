package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Mike")
		want := "Hello, Mike"

		assertCorrectMessage(t, got, want)
	})
	t.Run("saying hello to a place", func(t *testing.T) {
		got := Hello("Paris")
		want := "Hello, Paris"

		assertCorrectMessage(t, got, want)
	})
	t.Run("saying hello to an empty string", func(t *testing.T) {
		got := Hello("")
		want := "Hello, World"

		assertCorrectMessage(t, got, want)
	})
}

func assertCorrectMessage(t *testing.T, got, want string) {
	// t.Helper() is needed to tell the test suite that this method is a helper.
	// By doing this when it fails the line number reported will be in our function call rather than inside our test helper.
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
