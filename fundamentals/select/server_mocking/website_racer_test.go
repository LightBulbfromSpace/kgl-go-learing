package website_racer

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestWebsiteRacer(t *testing.T) {
	slowServer := makeDelayedServer(1 * time.Millisecond)
	fastServer := makeDelayedServer(0 * time.Millisecond)

	defer fastServer.Close()
	defer slowServer.Close()

	fastURL := fastServer.URL
	slowURL := slowServer.URL

	want := fastURL
	got := WebsiteRacer(fastURL, slowURL)

	if want != got {
		t.Errorf("wanted %q, got %q", want, got)
	}
}

func makeDelayedServer(delay time.Duration) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(delay)
		w.WriteHeader(http.StatusOK)
	}))
}
