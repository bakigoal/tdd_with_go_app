package store

type InMemoryPlayerStore struct {
}

func (ps *InMemoryPlayerStore) GetPlayerScore(player string) int {
	return 42
}

func (ps *InMemoryPlayerStore) RecordWin(player string) {

}
