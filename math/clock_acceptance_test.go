package clock_test

import (
	"github.com/LightBulbfromSpace/kld-go-learning/math/clock"
	"testing"
	"time"
)

func TestSecondHandAtMidnight(t *testing.T) {
	tm := time.Date(1970, time.January, 1, 0, 0, 0, 0, time.UTC)

	want := clock.SecondHand(tm)
	got := clock.Point{X: 150, Y: 150 - 90}

	if got != want {
		t.Errorf("got %v but want %v", want, got)
	}
}

func TestSecondHandAt30Seconds(t *testing.T) {
	tm := time.Date(1970, time.January, 1, 0, 0, 30, 0, time.UTC)

	want := clock.SecondHand(tm)
	got := clock.Point{X: 150, Y: 150 + 90}

	if got != want {
		t.Errorf("got %v but want %v", want, got)
	}
}
