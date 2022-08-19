package server

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

type StubPlayerStore struct {
	scores   map[string]int
	winCalls []string
}

func (s *StubPlayerStore) GetPlayerScore(name string) int {
	return s.scores[name]
}

func (s *StubPlayerStore) RecordWin(name string) {
	s.winCalls = append(s.winCalls, name)
}

func TestGETPlayers(t *testing.T) {
	stubStore := StubPlayerStore{
		scores: map[string]int{
			"Steve": 20,
			"David": 10,
		},
	}

	server := &PlayerServer{store: &stubStore}
	server.configureRouter()

	t.Run("returns Steve's score", func(t *testing.T) {
		request := newGetScoreRequest("Steve")
		responce := httptest.NewRecorder()

		server.ServeHTTP(responce, request)

		AssertStatus(t, responce.Code, http.StatusOK)
		AssertResponseBody(t, responce.Body.String(), "20")
	})
	t.Run("returns David's score", func(t *testing.T) {
		request := newGetScoreRequest("David")
		responce := httptest.NewRecorder()

		server.ServeHTTP(responce, request)

		AssertStatus(t, responce.Code, http.StatusOK)
		AssertResponseBody(t, responce.Body.String(), "10")

	})
	t.Run("returns 404 on missing players", func(t *testing.T) {
		request := newGetScoreRequest("Fred")
		responce := httptest.NewRecorder()

		server.ServeHTTP(responce, request)
		AssertStatus(t, responce.Code, http.StatusNotFound)
	})
}

func TestStoreWins(t *testing.T) {
	stubStore := StubPlayerStore{
		scores:   map[string]int{},
		winCalls: nil,
	}
	server := PlayerServer{store: &stubStore}
	server.configureRouter()

	t.Run("it returns accepted on POST", func(t *testing.T) {
		player := "Fred"
		request := newPostWinRequest(player)
		responce := httptest.NewRecorder()

		server.ServeHTTP(responce, request)
		AssertStatus(t, responce.Code, http.StatusAccepted)

		if len(stubStore.winCalls) != 1 {
			t.Errorf("got %d calls to RecordWin want %d", len(stubStore.winCalls), 1)
		}

		if stubStore.winCalls[0] != player {
			t.Errorf("did not store correct winner got %q want %q", stubStore.winCalls[0], player)
		}
	})
}

func TestLeague(t *testing.T) {
	store := &StubPlayerStore{}
	server := &PlayerServer{store: store}
	server.configureRouter()

	t.Run("it returns 200 on /league", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/league", nil)
		responce := httptest.NewRecorder()

		server.ServeHTTP(responce, request)

		AssertStatus(t, responce.Code, http.StatusOK)
	})
}

func newPostWinRequest(name string) *http.Request {
	request, _ := http.NewRequest(http.MethodPost, fmt.Sprintf("/players/%s", name), nil)
	return request
}

func newGetScoreRequest(name string) *http.Request {
	request, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("/players/%s", name), nil)
	return request
}
func AssertResponseBody(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("Got %s but want %s", got, want)
	}
}

func AssertStatus(t testing.TB, got, want int) {
	t.Helper()
	if got != want {
		t.Errorf("did not get correct status, got %d, want %d", got, want)
	}
}
