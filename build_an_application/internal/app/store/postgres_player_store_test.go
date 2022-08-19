package store

import (
	"testing"
)

func TestGetPlayerScoreAndRecordWin(t *testing.T) {
	config := NewConfigForTest()
	store := NewPostgresPlayerStore(config)
	if err := store.Open(); err != nil {
		t.Error("failed to open store", err)
	}
	player := "test_player"
	initialScore := 7
	t.Run("get player's score", func(t *testing.T) {
		got := store.GetPlayerScore(player)

		if got != initialScore {
			t.Errorf("got %d but wnat %d", got, initialScore)
		}
	})
	t.Run("record win", func(t *testing.T) {
		store.RecordWin(player)
		got := store.GetPlayerScore(player)
		if got != initialScore+1 {
			t.Errorf("got %d but wnat %d", got, initialScore+1)
		}
	})
	store.db.QueryRow("UPDATE players_scores SET score=$1 WHERE name=$2", initialScore, player)
}
