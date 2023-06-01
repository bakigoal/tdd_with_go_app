package store

import (
	"github.com/bakigoal/tdd_with_go_app/server"
	"io"
)

type FileSystemPlayerStore struct {
	Database io.ReadSeeker
}

func (s *FileSystemPlayerStore) GetPlayerScore(player string) int {
	return 0
}

func (s *FileSystemPlayerStore) RecordWin(player string) {
}

func (s *FileSystemPlayerStore) GetLeague() []server.Player {
	s.Database.Seek(0, 0)
	league, _ := NewLeague(s.Database)
	return league
}
