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
	var league []server.Player
	for name, wins := range ps.store {
		league = append(league, server.Player{Name: name, Wins: wins})
	}
	return league
}
