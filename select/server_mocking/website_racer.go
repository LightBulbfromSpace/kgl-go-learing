package website_racer

import (
	"net/http"
	"time"
)

func WebsiteRacer(a, b string) string {

	durationA := measureResponseTime(a)
	durationB := measureResponseTime(b)

	if durationA < durationB {
		return a
	}
	return b
}

func measureResponseTime(URL string) time.Duration {
	start := time.Now()
	http.Get(URL)
	return time.Since(start)
}
