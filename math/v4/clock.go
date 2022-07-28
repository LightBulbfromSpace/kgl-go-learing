package clock

import (
	"math"
	"time"
)

const (
	secondsInHalfClock = 30
	secondsInClock     = 2 * secondsInHalfClock
	minutesInHalfClock = 30
	minutesInClock     = 2 * minutesInHalfClock
	hoursInHalfClock   = 6
	hoursInClock       = 2 * hoursInHalfClock
)

func hourHandPoint(t time.Time) Point {
	return angleToPoint(hoursInRadian(t))
}

func minuteHandPoint(t time.Time) Point {
	return angleToPoint(minutesInRadian(t))
}

func secondHandPoint(t time.Time) Point {
	return angleToPoint(secondsInRadian(t))
}

func angleToPoint(angle float64) Point {
	x := math.Sin(angle)
	y := math.Cos(angle)

	return Point{x, y}
}

func hoursInRadian(t time.Time) float64 {
	return minutesInRadian(t)/hoursInClock +
		math.Pi/(hoursInHalfClock/float64(t.Hour()%hoursInClock))
}

func minutesInRadian(t time.Time) float64 {
	return secondsInRadian(t)/minutesInClock +
		math.Pi/(minutesInHalfClock/float64(t.Minute()))
}

func secondsInRadian(t time.Time) float64 {
	return math.Pi / (secondsInHalfClock / float64(t.Second()))
}
