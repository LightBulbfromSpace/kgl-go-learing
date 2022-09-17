package poker_test

import (
	poker "github.com/LightBulbfromSpace/build_an_application"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func TestRecordingWinsAndRetrievingThem(t *testing.T) {
	database, cleanDatabase := createTempFile(t, `[]`)
	defer cleanDatabase()
	store, err := poker.NewFileSystemPlayerStore(database)
	assert.NoError(t, err)
	server, err := poker.NewServer(store, dummyGame)
	assert.NoError(t, err)
	player := "Fred"

	for i := 0; i < 3; i++ {
		server.ServeHTTP(httptest.NewRecorder(), newPostWinRequest(player))
	}

	t.Run("get score", func(t *testing.T) {
		response := httptest.NewRecorder()
		server.ServeHTTP(response, newGetScoreRequest(player))

		poker.AssertStatus(t, response, http.StatusOK)
		poker.AssertResponseBody(t, response.Body.String(), "3")
	})
	t.Run("get league", func(t *testing.T) {
		response := httptest.NewRecorder()
		server.ServeHTTP(response, newLeagueRequest())

		poker.AssertStatus(t, response, http.StatusOK)
		got := getLeagueFromResponse(t, response.Body)
		want := poker.League{
			{"Fred", 3},
		}
		poker.AssertLeague(t, want, got)
	})
}

func createTempFile(t testing.TB, initialData string) (*os.File, func()) {
	t.Helper()

	tmpfile, err := os.CreateTemp("", "db")
	if err != nil {
		t.Fatalf("could not create temp File %v", err)
	}

	if _, err := tmpfile.Write([]byte(initialData)); err != nil {
		t.Fatalf("could not write in temp File %v", err)
	}

	removeFile := func() {
		tmpfile.Close()
		os.Remove(tmpfile.Name())
	}

	return tmpfile, removeFile
}
