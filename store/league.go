package store

import (
	"encoding/json"
	"github.com/bakigoal/tdd_with_go_app/server"
	"io"
)

type League []server.Player

func NewLeague(rdb io.Reader) (League, error) {
	var league League
	err := json.NewDecoder(rdb).Decode(&league)
	if err != nil {
		return nil, err
	}
	return league, nil
}

func (l League) Find(name string) *server.Player {
	for i, p := range l {
		if p.Name == name {
			return &l[i]
		}
	}
	return nil
}
