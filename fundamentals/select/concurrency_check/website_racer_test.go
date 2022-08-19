package website_racer

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestWebsiteRacer(t *testing.T) {
	t.Run("both servers responded for 10 sec", func(t *testing.T) {
		slowServer := makeDelayedServer(1 * time.Millisecond)
		fastServer := makeDelayedServer(0 * time.Millisecond)

		defer fastServer.Close()
		defer slowServer.Close()

		fastURL := fastServer.URL
		slowURL := slowServer.URL

		want := fastURL
		got, err := WebsiteRacer(fastURL, slowURL)

		if err != nil {
			t.Fatalf("got error %v but didn't expect one", err)
		}

		if want != got {
			t.Errorf("wanted %q, got %q", want, got)
		}
	})
	t.Run("returns an error if a server doesn't respond within 10s", func(t *testing.T) {
		slowServer := makeDelayedServer(12 * time.Millisecond)
		fastServer := makeDelayedServer(11 * time.Millisecond)

		defer fastServer.Close()
		defer slowServer.Close()

		fastURL := fastServer.URL
		slowURL := slowServer.URL

		_, err := ConfigurableWebsiteRacer(fastURL, slowURL, 10*time.Millisecond)

		if err == nil {
			t.Errorf("expected to get error")
		}
	})
}

func makeDelayedServer(delay time.Duration) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(delay)
		w.WriteHeader(http.StatusOK)
	}))
}
