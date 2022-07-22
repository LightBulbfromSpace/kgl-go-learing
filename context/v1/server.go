package server

import (
	"fmt"
	"net/http"
)

type Store interface {
	Fetch() string
	Cancel()
}

type StubStore struct {
	response string
}

func (s *StubStore) Fetch() string {
	return s.response
}

func (s *StubStore) Cancel() {}

func Server(store Store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		data := make(chan string, 1)

		go func() {
			data <- store.Fetch()
		}()

		select {
		case d := <-data:
			fmt.Fprint(w, d)

		case <-ctx.Done():
			store.Cancel()
		}
	}
}
