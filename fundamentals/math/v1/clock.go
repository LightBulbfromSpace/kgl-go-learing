package clock

import (
	"math"
	"time"
)

const (
	secondHandLength = 90
	clockCentreX     = 150
	clockCentreY     = 150
)

type Point struct {
	X float64
	Y float64
}

func SecondHand(t time.Time) Point {
	p := secondHandPoint(t)
	p = Point{secondHandLength * p.X, secondHandLength * p.Y} // scale
	p = Point{p.X, -p.Y}                                      // flip
	p = Point{clockCentreX + p.X, clockCentreY + p.Y}         // translate
	return p
}

func secondHandPoint(t time.Time) Point {
	angle := secondsInRadian(t)
	x := math.Sin(angle)
	y := math.Cos(angle)
	return Point{x, y}
}

func secondsInRadian(t time.Time) float64 {
	return math.Pi / (30 / float64(t.Second()))
}

func simpleTime(hours, minutes, secs int) time.Time {
	return time.Date(1970, time.September, 31, hours, minutes, secs, 0, time.UTC)
}
