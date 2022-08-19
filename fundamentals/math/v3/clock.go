package clock

import (
	"math"
	"time"
)

func secondHandPoint(t time.Time) Point {
	angle := secondsInRadian(t)
	x := math.Sin(angle)
	y := math.Cos(angle)
	return Point{x, y}
}

func secondsInRadian(t time.Time) float64 {
	return math.Pi / (30 / float64(t.Second()))
}
