package store

import (
	"encoding/json"
	"github.com/bakigoal/tdd_with_go_app/model"
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
		league = append(league, model.Player{Name: player, Wins: 1})
	}

	s.Database.Seek(0, 0)
	json.NewEncoder(s.Database).Encode(league)
}

func (s *FileSystemPlayerStore) GetLeague() model.League {
	s.Database.Seek(0, 0)
	league, _ := model.NewLeague(s.Database)
	return league
}
