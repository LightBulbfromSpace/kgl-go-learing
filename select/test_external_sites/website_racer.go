package website_racer

import (
	"net/http"
	"time"
)

func WebsiteRacer(a, b string) (winner string) {
	startA := time.Now()
	http.Get(a)
	durationA := time.Since(startA)

	startB := time.Now()
	http.Get(b)
	durationB := time.Since(startB)

	if durationA < durationB {
		return a
	}
	return b
}
