package store

import (
	"encoding/json"
	"io"
)

type FileSystemPlayerStore struct {
	Database io.ReadWriteSeeker
}

func (s *FileSystemPlayerStore) GetPlayerScore(player string) int {
	found := s.GetLeague().Find(player)
	if found != nil {
		return found.Wins
	}
	return 0
}

func (s *FileSystemPlayerStore) RecordWin(player string) {
	league := s.GetLeague()
	found := league.Find(player)
	if found != nil {
		found.Wins++
	} else {

	}

	s.Database.Seek(0, 0)
	json.NewEncoder(s.Database).Encode(league)
}

func (s *FileSystemPlayerStore) GetLeague() League {
	s.Database.Seek(0, 0)
	league, _ := NewLeague(s.Database)
	return league
}
