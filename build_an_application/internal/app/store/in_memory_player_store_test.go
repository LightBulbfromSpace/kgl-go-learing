package store

import (
	"sync"
	"testing"
)

func TestInMemoryStore(t *testing.T) {
	player := "Herb"
	initialScore := 0
	store := InMemoryPlayerStore{
		store: map[string]int{
			player: initialScore,
		},
	}
	t.Run("get value", func(t *testing.T) {
		score := store.GetPlayerScore(player)
		if score != initialScore {
			t.Errorf("got %d but want %d", score, initialScore)
		}
	})
	t.Run("records 100 wins", func(t *testing.T) {
		wins := 100
		var wg sync.WaitGroup
		wg.Add(wins)

		for i := 0; i < 100; i++ {
			go func() {
				store.RecordWin(player)
				wg.Done()
			}()
		}

		wg.Wait()

		got := store.GetPlayerScore(player)
		if got != wins {
			t.Errorf("got %d but want %d", got, wins)
		}
	})
}
