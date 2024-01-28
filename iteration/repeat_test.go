package iteration_test

import (
	"fmt"
	"go-with-tests/iteration"
	"testing"
)

func TestRepeat(t *testing.T) {
	t.Run("repeat 5 times", func(t *testing.T) {
		got := iteration.Repeat("a", 5)
		want := "aaaaa"

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})
	t.Run("repeat 10 times", func(t *testing.T) {
		got := iteration.Repeat("a", 10)
		want := "aaaaaaaaaa"

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		iteration.Repeat("a", 50)
	}
}

func ExampleRepeat() {
	repeated := iteration.Repeat("a", 5)
	fmt.Println(repeated)
	// Output: aaaaa
}
