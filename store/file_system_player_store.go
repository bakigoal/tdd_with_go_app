package store

import (
	"encoding/json"
	"github.com/bakigoal/tdd_with_go_app/model"
	"io"
)

type FileSystemPlayerStore struct {
	Database io.ReadWriteSeeker
	league   model.League
}

func NewFileSystemPlayerStore(database io.ReadWriteSeeker) *FileSystemPlayerStore {
	database.Seek(0, 0)
	league, _ := model.NewLeague(database)
	return &FileSystemPlayerStore{
		Database: database,
		league:   league,
	}
}

func (s *FileSystemPlayerStore) GetPlayerScore(player string) int {
	found := s.league.Find(player)
	if found != nil {
		return found.Wins
	}
	return 0
}

func (s *FileSystemPlayerStore) RecordWin(player string) {
	found := s.league.Find(player)
	if found != nil {
		found.Wins++
	} else {
		s.league = append(s.league, model.Player{Name: player, Wins: 1})
	}

	s.Database.Seek(0, 0)
	json.NewEncoder(s.Database).Encode(s.league)
}

func (s *FileSystemPlayerStore) GetLeague() model.League {
	return s.league
}
