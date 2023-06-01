package store

import (
	"encoding/json"
	"github.com/bakigoal/tdd_with_go_app/server"
	"io"
)

type FileSystemPlayerStore struct {
	Database io.ReadWriteSeeker
}

func (s *FileSystemPlayerStore) GetPlayerScore(player string) int {
	for _, p := range s.GetLeague() {
		if p.Name == player {
			return p.Wins
		}
	}
	return 0
}

func (s *FileSystemPlayerStore) RecordWin(player string) {
	league := s.GetLeague()
	for i, p := range s.GetLeague() {
		if p.Name == player {
			league[i].Wins++
		}
	}
	s.Database.Seek(0, 0)
	json.NewEncoder(s.Database).Encode(league)
}

func (s *FileSystemPlayerStore) GetLeague() []server.Player {
	s.Database.Seek(0, 0)
	league, _ := NewLeague(s.Database)
	return league
}
