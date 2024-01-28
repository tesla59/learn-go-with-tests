package structs_test

import (
	"go-with-tests/structs"
	"testing"
)

func TestPerimeter(t *testing.T) {
	rectangle := structs.Rectangle{Width: 10.0, Height: 10.0}
	got := rectangle.Perimeter()
	want := 40.0

	if got != want {
		t.Errorf("got %.2f want %.2f", got, want)
	}
}

func TestArea(t *testing.T) {
	checkArea := func(t *testing.T, shape structs.Shape, want float64) {
		t.Helper()
		got := shape.Area()
		if got != want {
			t.Errorf("got %.2f want %.2f", got, want)
		}
	}
	t.Run("rectangle", func(t *testing.T) {
		rectangle := structs.Rectangle{Width: 12.0, Height: 6.0}
		want := 72.0
		checkArea(t, rectangle, want)
	})
	t.Run("circle", func(t *testing.T) {
		circle := structs.Circle{Radius: 10}
		want := 314.1592653589793
		checkArea(t, circle, want)
	})
}
