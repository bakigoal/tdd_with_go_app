package store

import (
	"encoding/json"
	"github.com/bakigoal/tdd_with_go_app/src/model"
	"github.com/bakigoal/tdd_with_go_app/src/utils"
	"os"
)

type FileSystemPlayerStore struct {
	Database *json.Encoder
	league   model.League
}

func NewFileSystemPlayerStore(database *os.File) *FileSystemPlayerStore {
	database.Seek(0, 0)
	league, _ := model.NewLeague(database)
	return &FileSystemPlayerStore{
		Database: json.NewEncoder(&utils.Tape{File: database}),
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

	s.Database.Encode(s.league)
}

func (s *FileSystemPlayerStore) GetLeague() model.League {
	return s.league
}
