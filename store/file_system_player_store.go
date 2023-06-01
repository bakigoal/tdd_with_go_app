package store

import (
	"github.com/bakigoal/tdd_with_go_app/server"
	"strings"
)

type FileSystemPlayerStore struct {
	Database *strings.Reader
}

func (s *FileSystemPlayerStore) GetPlayerScore(player string) int {
	return 0
}

func (s *FileSystemPlayerStore) RecordWin(player string) {
}

func (s *FileSystemPlayerStore) GetLeague() []server.Player {
	league, err := NewLeague(s.Database)
	if err != nil {
		return nil
	}
	return league
}
