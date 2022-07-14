package main

import (
	"bytes"
	"reflect"
	"testing"
	"time"
)

func TestCountdown(t *testing.T) {
	t.Run("ensures about messages are printed out", func(t *testing.T) {
		buffer := &bytes.Buffer{}
		spySleeperPrinter := &SpyCountdownOperations{}

		Countdown(buffer, spySleeperPrinter)

		got := buffer.String()
		want := "3\n" +
			"2\n" +
			"1\n" +
			"Go!\n"

		if want != got {
			t.Errorf("got %q but want %q", got, want)
		}

	})
	t.Run("sleep before every print", func(t *testing.T) {
		spySleeperPrinter := &SpyCountdownOperations{}

		Countdown(spySleeperPrinter, spySleeperPrinter)

		want := []string{
			write,
			sleep,
			write,
			sleep,
			write,
			sleep,
			write,
		}

		if !reflect.DeepEqual(want, spySleeperPrinter.Calls) {
			t.Errorf("got %v but want %v", spySleeperPrinter.Calls, want)
		}
	})
}

func TestConfigurableSleeper(t *testing.T) {
	sleepTime := 5 * time.Second
	spyTime := &SpyTime{}

	sleeper := ConfigurableSleeper{sleepTime, spyTime.Sleep}
	sleeper.Sleep()

	if spyTime.durationSlept != sleepTime {
		t.Errorf("should have slept for %v but slept for %v", sleepTime, spyTime.durationSlept)
	}
}
