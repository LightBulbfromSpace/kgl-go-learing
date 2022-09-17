package poker

import (
	"io"
	"time"
)

type TexasHoldem struct {
	alerter BlindAlerter
	store   PlayerStore
}

func NewTexasHoldem(store PlayerStore, alerter BlindAlerter) *TexasHoldem {
	return &TexasHoldem{
		store:   store,
		alerter: alerter,
	}
}

func (g *TexasHoldem) Start(players int, alertsDestination io.Writer) {
	blindIncrement := time.Duration(5+players) * time.Minute // time.Minute for real game, time.Second for test

	blinds := []int{100, 200, 300, 400, 500, 600, 800, 1000, 2000, 4000, 8000}
	blindTime := 0 * time.Minute // time.Minute for real game, time.Second for test
	for _, blind := range blinds {
		g.alerter.SheduleAlertAt(blindTime, blind, alertsDestination)
		blindTime += blindIncrement
	}
}

func (g *TexasHoldem) Finish(winner string) {
	g.store.RecordWin(winner)
}
