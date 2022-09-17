package poker

import (
	"sync"
)

type InMemoryPlayerStore struct {
	store map[string]int
	mu    sync.Mutex
}

func NewInMemoryPlayerStore() *InMemoryPlayerStore {
	return &InMemoryPlayerStore{store: map[string]int{}}
}

func (s *InMemoryPlayerStore) GetPlayerScore(name string) int {
	return s.store[name]
}

func (s *InMemoryPlayerStore) RecordWin(name string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.store[name]++
}

func (s *InMemoryPlayerStore) GetLeague() League {
	var league []Player
	for player, score := range s.store {
		league = append(league, Player{Name: player, Wins: score})
	}
	return league
}
