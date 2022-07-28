package clock

import (
	"math"
	"testing"
	"time"
)

func TestSecondsInRadian(t *testing.T) {
	cases := []struct {
		time  time.Time
		angle float64
	}{
		{simpleTime(0, 0, 30), math.Pi},
		{simpleTime(0, 0, 45), 3 * math.Pi / 2},
		{simpleTime(0, 0, 15), math.Pi / 2},
		{simpleTime(0, 0, 17), math.Pi / 30 * 17},
	}

	for _, tc := range cases {
		t.Run(testName(tc.time), func(t *testing.T) {
			got := secondsInRadian(tc.time)
			if tc.angle != got {
				t.Fatalf("got %.16f but wanted radians %.16f", tc.angle, got)
			}
		})
	}
}

func testName(t time.Time) string {
	return t.Format("15:01:05")
}

func TestSecondHandEnd(t *testing.T) {
	cases := []struct {
		time  time.Time
		point Point
	}{
		{simpleTime(0, 0, 30), Point{0, -1}},
		{simpleTime(0, 0, 45), Point{-1, 0}},
	}
	for _, tc := range cases {
		t.Run(testName(tc.time), func(t *testing.T) {
			got := secondHandPoint(tc.time)
			if !roughlyEqualPoints(tc.point, got) {
				t.Fatalf("got %v but wanted point %v", got, tc.point)
			}
		})
	}
}

func roughlyEqualFloat64(a, b float64) bool {
	const equalityThreshold = 1e-7
	return math.Abs(a-b) < equalityThreshold
}

func roughlyEqualPoints(a, b Point) bool {
	return roughlyEqualFloat64(a.X, b.X) && roughlyEqualFloat64(a.Y, b.Y)
}

func simpleTime(hours, minutes, secs int) time.Time {
	return time.Date(1970, time.September, 31, hours, minutes, secs, 0, time.UTC)
}
