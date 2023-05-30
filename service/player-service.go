package service

type PlayerService struct {
}

func (ps *PlayerService) GetPlayerScore(player string) int {
	if player == "Pepper" {
		return 20
	}
	if player == "Floyd" {
		return 10
	}
	return 0
}
