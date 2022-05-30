package perimeter

import (
	"math"
)

type Shape interface {
	Perimeter() float64
	Area() float64
}

type Square struct {
	side float64
}

func (s Square) Perimeter() float64 {
	return s.side * 4
}

func (s Square) Area() float64 {
	return math.Pow(s.side, 2)
}

type Rectangle struct {
	width  float64
	height float64
}

func (rect Rectangle) Perimeter() float64 {
	return (rect.width + rect.height) * 2
}

func (rect Rectangle) Area() float64 {
	return rect.width * rect.height
}

type RightTriangle struct {
	a float64
	b float64
}

func (t RightTriangle) Perimeter() float64 {
	return t.a + t.b + math.Sqrt(t.a*t.a+t.b*t.b)
}

func (t RightTriangle) Area() float64 {
	return (t.a * t.b) / 2
}

type Circle struct {
	rad float64
}

func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.rad
}

func (c Circle) Area() float64 {
	return math.Pow(c.rad, 2) * math.Pi
}

/*func Perimeter(nameOfFigure string, l ...float64) float64 {
=======
	"fmt"
	"math"
)

func Perimeter(nameOfFigure string, l ...float64) float64 {
>>>>>>> f59ce94d2e9982d8d48a01c2906e2296dad90863
	switch nameOfFigure {
	case "square":
		return l[0] * 4
	case "rectangle":
		return (l[0] + l[1]) * 2
	case "right triangle":
		return l[0] + l[1] + math.Sqrt(l[0]*l[0]+l[1]*l[1])
	default:
		fmt.Print("This shape is not available or one hasn't been specified\n")
		return 0
	}
<<<<<<< HEAD
}*/

/*func Area(nameOfFigure string, l ...float64) float64 {
=======
}

func Area(nameOfFigure string, l ...float64) float64 {
>>>>>>> f59ce94d2e9982d8d48a01c2906e2296dad90863
	switch nameOfFigure {
	case "square":
		return l[0] * l[0]
	case "rectangle":
		return l[0] * l[1]
	case "right triangle":
		return (l[0] * l[1]) / 2
	default:
		fmt.Print("This shape is not available or one hasn't been specified\n")
		return 0
	}
<<<<<<< HEAD
}*/
