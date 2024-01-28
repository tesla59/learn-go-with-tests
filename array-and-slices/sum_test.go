package arrayandslices_test

import (
	arrayandslices "go-with-tests/array-and-slices"
	"slices"
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

func TestSumAll(t *testing.T) {
	got := arrayandslices.SumAll([]int{1, 2}, []int{4, 7})
	want := []int{3, 11}

	if slices.Equal(got, want) != true {
		t.Errorf("got %v want %v", got, want)
	}
}
