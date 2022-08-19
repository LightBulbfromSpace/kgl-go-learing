package perimeter

import (
	"testing"
)

func TestPerimeter(t *testing.T) {
	perimeterTests := []struct {
		name  string
		shape Shape
		want  float64
	}{
		{name: "Square", shape: Square{side: 10.0}, want: 40.0},
		{name: "Rectangle", shape: Rectangle{width: 30.0, height: 40.0}, want: 140.0},
		{name: "RightTriangle", shape: RightTriangle{a: 3.0, b: 4.0}, want: 12.0},
		{name: "Circle", shape: Circle{rad: 10.0}, want: 62.83185307179586},
	}

	for _, figt := range perimeterTests {
		t.Run(figt.name, func(t *testing.T) {
			got := figt.shape.Perimeter()
			if got != figt.want {
				t.Errorf("got %g want %g", got, figt.want)
			}
		})
	}
}

func TestArea(t *testing.T) {
	areaTests := []struct {
		name  string
		shape Shape
		want  float64
	}{
		{name: "Square", shape: Square{side: 10.0}, want: 100.0},
		{name: "Rectangle", shape: Rectangle{width: 10.0, height: 20.0}, want: 200.0},
		{name: "RightTriangle", shape: RightTriangle{3.0, 4.0}, want: 6.0},
		{name: "Circle", shape: Circle{rad: 10.0}, want: 314.1592653589793},
	}
	for _, tt := range areaTests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.shape.Area()
			if got != tt.want {
				t.Errorf("%#v got %g want %g", tt.shape, got, tt.want)
			}
		})
	}
}
