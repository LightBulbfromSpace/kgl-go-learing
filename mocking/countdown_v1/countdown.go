package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

const (
	finalWord      = "Go!"
	countdownStart = 3
)

type Sleeper interface {
	Sleep()
}
type SpySleeper struct {
	Calls int
}

type RealSleeper struct{}

func (s *SpySleeper) Sleep() {
	s.Calls++
}

func (rs *RealSleeper) Sleep() {
	time.Sleep(1 * time.Second)
}

func Countdown(out io.Writer, slp Sleeper) {
	for i := countdownStart; i > 0; i-- {
		fmt.Fprintln(out, i)
		slp.Sleep()
	}
	fmt.Fprintln(out, finalWord)
}

func main() {
	sleeper := &RealSleeper{}
	Countdown(os.Stdout, sleeper)
}
