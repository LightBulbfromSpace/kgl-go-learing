package website_racer

import "testing"

func TestWebsiteRacer(t *testing.T) {

	slowURL := "https://go.dev/"
	fastURL := "https://web.telegram.org/"

	want := fastURL
	got := WebsiteRacer(slowURL, fastURL)

	if got != want {
		t.Errorf("got %q but want %q", got, want)
	}
}
