package poker_test

import (
	poker "github.com/LightBulbfromSpace/build_an_application"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFileSystemStore(t *testing.T) {
	t.Run("league sorted", func(t *testing.T) {
		database, cleanDatabase := createTempFile(t, `[
            {"Name": "Cleo", "Wins": 10},
            {"Name": "Chris", "Wins": 33}]`)

		defer cleanDatabase()

		store, err := poker.NewFileSystemPlayerStore(database)
		assert.NoError(t, err)

		got := store.GetLeague()
		want := []poker.Player{
			{"Chris", 33},
			{"Cleo", 10},
		}
		poker.AssertLeague(t, want, got)

		got = store.GetLeague()
		poker.AssertLeague(t, got, want)
	})
	t.Run("get player score", func(t *testing.T) {
		database, cleanDatabase := createTempFile(t, `[
            {"Name": "Cleo", "Wins": 10},
            {"Name": "Chris", "Wins": 33}]`)

		defer cleanDatabase()

		store, err := poker.NewFileSystemPlayerStore(database)
		assert.NoError(t, err)
		got := store.GetPlayerScore("Chris")
		want := 33
		assertScoreEquals(t, got, want)
	})
	t.Run("store wins for existing players", func(t *testing.T) {
		database, cleanDatabase := createTempFile(t, `[
            {"Name": "Cleo", "Wins": 10},
            {"Name": "Chris", "Wins": 33}]`)

		defer cleanDatabase()

		store, err := poker.NewFileSystemPlayerStore(database)
		assert.NoError(t, err)
		store.RecordWin("Chris")

		got := store.GetPlayerScore("Chris")
		want := 34
		assertScoreEquals(t, got, want)
	})
	t.Run("store wins for new players", func(t *testing.T) {
		database, cleanDatabase := createTempFile(t, `[
            {"Name": "Cleo", "Wins": 10},
            {"Name": "Chris", "Wins": 33}]`)

		defer cleanDatabase()

		store, err := poker.NewFileSystemPlayerStore(database)
		assert.NoError(t, err)
		store.RecordWin("Alex")

		got := store.GetPlayerScore("Alex")
		want := 1
		assertScoreEquals(t, got, want)
	})
	t.Run("works with an empty File", func(t *testing.T) {
		database, cleanDatabase := createTempFile(t, ``)

		defer cleanDatabase()

		_, err := poker.NewFileSystemPlayerStore(database)
		assert.NoError(t, err)
	})
}

func assertScoreEquals(t testing.TB, got, want int) {
	t.Helper()
	if got != want {
		t.Errorf("got score %d but want %d", got, want)
	}
}
