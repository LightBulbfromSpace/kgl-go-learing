package poker_test

import (
	"fmt"
	poker "github.com/LightBulbfromSpace/build_an_application"
	"io"
	"testing"
	"time"
)

func TestGame_Start(t *testing.T) {
	t.Run("sheduled alerts for 5 players", func(t *testing.T) {
		blindAlerter := &SpyBlindAlerter{}
		game := poker.NewTexasHoldem(dummyPlayerStore, blindAlerter)
		game.Start(5, io.Discard)
		cases := []sheduledAlert{
			{0 * time.Minute, 100},
			{10 * time.Minute, 200},
			{20 * time.Minute, 300},
			{30 * time.Minute, 400},
			{40 * time.Minute, 500},
			{50 * time.Minute, 600},
			{60 * time.Minute, 800},
			{70 * time.Minute, 1000},
			{80 * time.Minute, 2000},
			{90 * time.Minute, 4000},
			{100 * time.Minute, 8000},
		}
		checkSchedulingCases(cases, t, blindAlerter)
	})
	t.Run("sheduled alerts for 7 players", func(t *testing.T) {
		blindAlerter := &SpyBlindAlerter{}
		game := poker.NewTexasHoldem(dummyPlayerStore, blindAlerter)
		game.Start(7, io.Discard)
		cases := []sheduledAlert{
			{0 * time.Minute, 100},
			{12 * time.Minute, 200},
			{24 * time.Minute, 300},
			{36 * time.Minute, 400},
		}
		checkSchedulingCases(cases, t, blindAlerter)
	})
}

func TestGame_Finish(t *testing.T) {
	store := &poker.StubPlayerStore{}
	winner := "Alex\n"

	game := poker.NewTexasHoldem(store, dummyBlindAlerter)

	game.Finish(winner)

	poker.AssertPlayerWin(t, store, winner)

}

func checkSchedulingCases(cases []sheduledAlert, t *testing.T, blindAlerter *SpyBlindAlerter) {
	for i, want := range cases {
		t.Run(fmt.Sprintf("%d scheduled for %v", want.amount, want.sheduledAt),
			func(t *testing.T) {

				if len(blindAlerter.alerts) <= i {
					t.Fatalf("alert %d was not scheduled %v", i, blindAlerter.alerts)
				}

				got := blindAlerter.alerts[i]
				assertAlert(t, got, want)
			})
	}
}
