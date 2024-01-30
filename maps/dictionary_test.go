package maps_test

import (
	"go-with-tests/maps"
	"testing"
)

func TestSearch(t *testing.T) {
	dict := maps.Dictionary{"test": "this is just a test"}

	t.Run("known word", func(t *testing.T) {
		got, _ := dict.Search("test")
		want := "this is just a test"

		assertStrings(t, got, want)
	})
	t.Run("unknown word", func(t *testing.T) {
		_, got := dict.Search("doest exist")
		want := maps.ErrNotFound
		assertError(t, got, want)
	})
}

func TestAdd(t *testing.T) {
	t.Run("add a new word", func(t *testing.T) {

		dict := maps.Dictionary{}
		dict.Add("test", "this is a test addition")

		got, _ := dict.Search("test")
		want := "this is a test addition"
		assertStrings(t, got, want)
	})
}

func assertStrings(t *testing.T, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func assertError(t *testing.T, got, want error) {
	t.Helper()
	if got != want {
		t.Errorf("got error %q want error %q", got, want)
	}
}
