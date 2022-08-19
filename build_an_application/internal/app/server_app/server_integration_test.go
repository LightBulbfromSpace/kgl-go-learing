package server

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRecordingWinsAndRetrievingThem(t *testing.T) {
	server := NewServerForInMemoryStore()
	player := "Fred"

	for i := 0; i < 3; i++ {
		server.router.ServeHTTP(httptest.NewRecorder(), newPostWinRequest(player))
	}

	response := httptest.NewRecorder()
	server.router.ServeHTTP(response, newGetScoreRequest(player))

	AssertStatus(t, response.Code, http.StatusOK)
	AssertResponseBody(t, response.Body.String(), "3")
}
