package perimeter

import (
	"fmt"
	"math"
)

func Perimeter(nameOfFigure string, l ...float64) float64 {
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
}

func Area(nameOfFigure string, l ...float64) float64 {
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
}
