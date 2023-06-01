package store

import (
	"encoding/json"
	"github.com/bakigoal/tdd_with_go_app/server"
	"io"
)

func NewLeague(rdb io.Reader) ([]server.Player, error) {
	var league []server.Player
	err := json.NewDecoder(rdb).Decode(&league)
	if err != nil {
		return nil, err
	}
	return league, nil
}
