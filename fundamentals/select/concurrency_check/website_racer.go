package website_racer

import (
	"fmt"
	"net/http"
	"time"
)

const tenSecTimeout = 10 * time.Second

func WebsiteRacer(a, b string) (winner string, err error) {
	return ConfigurableWebsiteRacer(a, b, tenSecTimeout)
}

func ConfigurableWebsiteRacer(a, b string, timeout time.Duration) (winner string, err error) {
	select {
	case <-ping(a):
		return a, nil
	case <-ping(b):
		return b, nil
	case <-time.After(timeout):
		return "", fmt.Errorf("timed out waiting for %s and %s", a, b)
	}
}

func ping(URL string) chan struct{} {
	ch := make(chan struct{})
	go func() {
		http.Get(URL)
		close(ch)
	}()
	return ch
}
