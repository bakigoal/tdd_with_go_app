package store

import "github.com/bakigoal/tdd_with_go_app/server"

func NewInMemoryPlayerStore() *InMemoryPlayerStore {
	return &InMemoryPlayerStore{map[string]int{}}
}

type InMemoryPlayerStore struct {
	store map[string]int
}

func (ps *InMemoryPlayerStore) GetPlayerScore(player string) int {
	return ps.store[player]
}

func (ps *InMemoryPlayerStore) RecordWin(player string) {
	ps.store[player]++
}

func (ps InMemoryPlayerStore) GetLeague() []server.Player {
	return []server.Player{}
}
