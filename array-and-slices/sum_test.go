package arrayandslices_test

import (
	arrayandslices "go-with-tests/array-and-slices"
	"testing"
)

func TestSum(t *testing.T) {
	num := []int{1, 2, 3, 4, 5}
	got := arrayandslices.Sum(num)
	want := 15
	if got != want {
		t.Errorf("got %d want %d given, %v", got, want, num)
	}
}
