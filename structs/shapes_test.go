package structs_test

import (
	"go-with-tests/structs"
	"testing"
)

func TestPerimeter(t *testing.T) {
	checkPerimeter := []struct {
		shape structs.Shape
		want  float64
	}{
		{structs.Rectangle{Width: 10, Height: 10}, 40.0},
		{structs.Circle{Radius: 10}, 62.83185307179586},
	}
	for _, v := range checkPerimeter {
		got := v.shape.Perimeter()
		if got != v.want {
			t.Errorf("got %g want %g", got, v.want)
		}
	}
}

func TestArea(t *testing.T) {
	checkArea := []struct {
		shape structs.Shape
		want  float64
	}{
		{structs.Rectangle{Width: 12, Height: 6}, 72.0},
		{structs.Circle{Radius: 10}, 314.1592653589793},
	}
	for _, v := range checkArea {
		got := v.shape.Area()
		if got != v.want {
			t.Errorf("got %g want %g", got, v.want)
		}
	}
}
