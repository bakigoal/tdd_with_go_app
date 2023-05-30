package server

import (
	"fmt"
	"net/http"
	"strings"
)

type PlayerStore interface {
	GetPlayerScore(name string) int
	RecordWin(player string)
}

type PlayerServer struct {
	Store PlayerStore
}

func (p *PlayerServer) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case http.MethodPost:
		p.processWin(w, req)
	case http.MethodGet:
		p.showScore(w, req)
	}
}

func (p *PlayerServer) processWin(w http.ResponseWriter, req *http.Request) {
	player := strings.TrimPrefix(req.URL.Path, "/players/")
	p.Store.RecordWin(player)
	w.WriteHeader(http.StatusAccepted)
}

func (p *PlayerServer) showScore(w http.ResponseWriter, req *http.Request) {
	player := strings.TrimPrefix(req.URL.Path, "/players/")
	score := p.Store.GetPlayerScore(player)
	if score == 0 {
		w.WriteHeader(http.StatusNotFound)
	}

	fmt.Fprint(w, score)
}
