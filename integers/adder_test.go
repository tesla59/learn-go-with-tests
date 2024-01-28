package integers_test

import (
	"fmt"
	"go-with-tests/integers"
	"testing"
)

func TestAdder(t *testing.T) {
	sum := integers.Add(2, 2)
	expected := 4

	if sum != expected {
		t.Errorf("expected %d but got %d", expected, sum)
	}
}

func ExampleAdd() {
	sum := integers.Add(5, 5)
	fmt.Println(sum)
	// Output: 10
}
