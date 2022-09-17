package poker_test

import (
	"encoding/json"
	"fmt"
	poker "github.com/LightBulbfromSpace/build_an_application"
	"github.com/gorilla/websocket"
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"
)

func TestGETPlayers(t *testing.T) {
	stubStore := poker.StubPlayerStore{
		Scores: map[string]int{
			"Steve": 20,
			"David": 10,
		},
	}

	server, err := poker.NewServer(&stubStore, dummyGame)
	assert.NoError(t, err)

	t.Run("returns Steve's score", func(t *testing.T) {
		request := newGetScoreRequest("Steve")
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		poker.AssertStatus(t, response, http.StatusOK)
		poker.AssertResponseBody(t, response.Body.String(), "20")
	})
	t.Run("returns David's score", func(t *testing.T) {
		request := newGetScoreRequest("David")
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		poker.AssertStatus(t, response, http.StatusOK)
		poker.AssertResponseBody(t, response.Body.String(), "10")

	})
	t.Run("returns 404 on missing players", func(t *testing.T) {
		request := newGetScoreRequest("Fred")
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		poker.AssertStatus(t, response, http.StatusNotFound)
	})
}

func TestStoreWins(t *testing.T) {
	stubStore := &poker.StubPlayerStore{
		Scores:   map[string]int{},
		WinCalls: nil,
	}
	server, err := poker.NewServer(stubStore, dummyGame)
	assert.NoError(t, err)

	t.Run("it returns accepted on POST", func(t *testing.T) {
		player := "Fred"
		request := newPostWinRequest(player)
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		poker.AssertStatus(t, response, http.StatusAccepted)
		poker.AssertPlayerWin(t, stubStore, player)
	})
}

func TestLeague(t *testing.T) {
	wantedLeague := []poker.Player{
		{"Cleo", 32},
		{"Chris", 20},
		{"Tiest", 14},
	}
	server, err := poker.NewServer(&poker.StubPlayerStore{League: wantedLeague}, dummyGame)
	assert.NoError(t, err)

	t.Run("it returns 200 on /league", func(t *testing.T) {
		request := newLeagueRequest()
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		got := getLeagueFromResponse(t, response.Body)

		poker.AssertStatus(t, response, http.StatusOK)
		poker.AssertLeague(t, wantedLeague, got)
		poker.AssertContentType(t, response, poker.JSONContentType)
	})
}

func TestGame(t *testing.T) {
	t.Run("it returns 200 on /game", func(t *testing.T) {
		server, err := poker.NewServer(&poker.StubPlayerStore{}, dummyGame)
		assert.NoError(t, err)
		request := newGameRequest()
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		poker.AssertStatus(t, response, http.StatusOK)
	})
	t.Run("start a game with 3 players and declare Ruth the winner", func(t *testing.T) {
		wantedBlindAlert := "Blind is 100"
		winner := "Ruth"

		game := &GameSpy{BlindAlert: []byte(wantedBlindAlert)}
		server := httptest.NewServer(mustMakePlayerServer(t, dummyPlayerStore, game))
		ws := mustDialWS(t, "ws"+strings.TrimPrefix(server.URL, "http")+"/ws")

		defer server.Close()
		defer ws.Close()

		writeWSMessage(t, ws, "3")
		writeWSMessage(t, ws, winner)

		time.Sleep(10 * time.Millisecond)

		assertGameStartedWith(t, game, 3)
		assertFinishCalledWith(t, game, winner)

		within(t, 10*time.Millisecond, func() { assertWebSocketSGotMsg(t, ws, wantedBlindAlert) })
	})
}

func newLeagueRequest() *http.Request {
	request, _ := http.NewRequest(http.MethodGet, "/league", nil)
	return request
}

func newPostWinRequest(name string) *http.Request {
	request, _ := http.NewRequest(http.MethodPost, fmt.Sprintf("/players/%s", name), nil)
	return request
}

func newGetScoreRequest(name string) *http.Request {
	request, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("/players/%s", name), nil)
	return request
}

func newGameRequest() *http.Request {
	request, _ := http.NewRequest(http.MethodGet, "/game", nil)
	return request
}

func getLeagueFromResponse(t *testing.T, body io.Reader) []poker.Player {
	var league []poker.Player
	err := json.NewDecoder(body).Decode(&league)
	if err != nil {
		t.Fatalf("Unable to parse response from server %q into slice of Player, '%v'", body, err)
	}
	return league
}

func mustMakePlayerServer(t *testing.T, store poker.PlayerStore, game *GameSpy) *poker.PlayerServer {
	server, err := poker.NewServer(store, game)
	if err != nil {
		t.Fatal("problem creating player server", err)
	}
	return server
}

func mustDialWS(t *testing.T, url string) *websocket.Conn {
	ws, _, err := websocket.DefaultDialer.Dial(url, nil)
	if err != nil {
		t.Fatalf("could not open a ws connection on %s %v", url, err)
	}
	return ws
}

func writeWSMessage(t testing.TB, ws *websocket.Conn, msg string) {
	t.Helper()
	if err := ws.WriteMessage(websocket.TextMessage, []byte(msg)); err != nil {
		t.Fatalf("could not send a message over ws connection on %v", err)
	}
}

func within(t testing.TB, d time.Duration, assert func()) {
	t.Helper()

	done := make(chan struct{}, 1)

	go func() {
		assert()
		done <- struct{}{}
	}()

	select {
	case <-time.After(d):
		t.Error("timed out")
	case <-done:
	}
}

func assertWebSocketSGotMsg(t *testing.T, ws *websocket.Conn, expected string) {
	_, msg, err := ws.ReadMessage()
	assert.NoError(t, err)
	assert.Equal(t, expected, string(msg))
}
