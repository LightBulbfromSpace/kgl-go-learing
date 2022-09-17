package poker

import (
	"net/http/httptest"
	"reflect"
	"testing"
)

type StubPlayerStore struct {
	Scores   map[string]int
	WinCalls []string
	League   []Player
}

func (s *StubPlayerStore) GetPlayerScore(name string) int {
	return s.Scores[name]
}

func (s *StubPlayerStore) RecordWin(name string) {
	s.WinCalls = append(s.WinCalls, name)
}

func (s *StubPlayerStore) GetLeague() League {
	return s.League
}

func AssertResponseBody(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("Got %s but want %s", got, want)
	}
}

func AssertStatus(t testing.TB, got *httptest.ResponseRecorder, want int) {
	t.Helper()
	if got.Code != want {
		t.Errorf("did not get correct status, got %d, want %d", got.Code, want)
	}
}

func AssertLeague(t testing.TB, wantedLeague, got []Player) {
	t.Helper()
	if !reflect.DeepEqual(wantedLeague, got) {
		t.Errorf("got %v want %v", got, wantedLeague)
	}
}

func AssertContentType(t testing.TB, response *httptest.ResponseRecorder, expectedContentType string) {
	t.Helper()
	if response.Result().Header.Get("Content-Type") != expectedContentType {
		t.Errorf("response did not have content-type of %s, got %v", expectedContentType, response.Result().Header)
	}
}

func AssertPlayerWin(t testing.TB, store *StubPlayerStore, winner string) {
	t.Helper()

	if len(store.WinCalls) != 1 {
		t.Fatalf("got %d calls to RecordWin want %d", len(store.WinCalls), 1)
	}

	if store.WinCalls[0] != winner {
		t.Errorf("did not store correct winner got %q want %q", store.WinCalls[0], winner)
	}
}
