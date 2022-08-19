package clock

import (
	"math"
	"testing"
	"time"
)

func TestHoursInRadian(t *testing.T) {
	cases := []struct {
		time  time.Time
		angle float64
	}{
		{simpleTime(0, 0, 0), math.Pi * 0},
		{simpleTime(3, 0, 0), math.Pi / 2},
		{simpleTime(7, 30, 0), math.Pi / 12 * 15},
		{simpleTime(21, 0, 0), math.Pi * 1.5},
		{simpleTime(0, 1, 30), math.Pi / ((6 * 60 * 60) / 90)},
	}

	for _, tc := range cases {
		t.Run(testName(tc.time), func(t *testing.T) {
			got := hoursInRadian(tc.time)
			if !roughlyEqualFloat64(tc.angle, got) {
				t.Fatalf("got %.6f but wanted radians %.6f", got, tc.angle)
			}
		})
	}
}

func TestMinutesInRadian(t *testing.T) {
	cases := []struct {
		time  time.Time
		angle float64
	}{
		{simpleTime(0, 0, 0), math.Pi * 0},
		{simpleTime(0, 0, 7), math.Pi / (30 * 60) * 7},
		{simpleTime(0, 15, 0), math.Pi / 2},
		{simpleTime(0, 17, 0), math.Pi / 30 * 17},
		{simpleTime(0, 30, 0), math.Pi},
		{simpleTime(0, 45, 0), 3 * math.Pi / 2},
	}

	for _, tc := range cases {
		t.Run(testName(tc.time), func(t *testing.T) {
			got := minutesInRadian(tc.time)
			if !roughlyEqualFloat64(tc.angle, got) {
				t.Fatalf("got %.16f but wanted radians %.16f", got, tc.angle)
			}
		})
	}
}

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
			if !roughlyEqualFloat64(tc.angle, got) {
				t.Fatalf("got %.16f but wanted radians %.16f", got, tc.angle)
			}
		})
	}
}

func testName(t time.Time) string {
	return t.Format("15:01:05")
}

func TestHourHandEnd(t *testing.T) {
	cases := []struct {
		time  time.Time
		point Point
	}{
		{simpleTime(6, 0, 0), Point{0, -1}},
		{simpleTime(21, 0, 0), Point{-1, 0}},
	}
	for _, tc := range cases {
		t.Run(testName(tc.time), func(t *testing.T) {
			got := hourHandPoint(tc.time)
			if !roughlyEqualPoints(tc.point, got) {
				t.Fatalf("got %v but wanted point %v", got, tc.point)
			}
		})
	}
}

func TestMinuteHandEnd(t *testing.T) {
	cases := []struct {
		time  time.Time
		point Point
	}{
		{simpleTime(0, 0, 30), Point{math.Sin(math.Pi / 60), math.Cos(math.Pi / 60)}},
		{simpleTime(0, 30, 0), Point{0, -1}},
		{simpleTime(0, 45, 0), Point{-1, 0}},
	}
	for _, tc := range cases {
		t.Run(testName(tc.time), func(t *testing.T) {
			got := minuteHandPoint(tc.time)
			if !roughlyEqualPoints(tc.point, got) {
				t.Fatalf("got %v but wanted point %v", got, tc.point)
			}
		})
	}
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
