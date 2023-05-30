package store

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
