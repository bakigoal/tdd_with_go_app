package test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/bakigoal/tdd_with_go_app/server"
	"github.com/stretchr/testify/assert"
)

type StubPlayerStore struct {
	scores   map[string]int
	winCalls []string
}

func (ps *StubPlayerStore) GetPlayerScore(player string) int {
	return ps.scores[player]
}

func (ps *StubPlayerStore) RecordWin(player string) {
	ps.winCalls = append(ps.winCalls, player)
}

func TestGETPlayers(t *testing.T) {
	store := StubPlayerStore{
		map[string]int{
			"Pepper": 20,
			"Floyd":  10,
		},
		[]string{},
	}
	server := &server.PlayerServer{Store: &store}
	t.Run("returns Pepper's score", func(t *testing.T) {
		req, _ := http.NewRequest(http.MethodGet, "/players/Pepper", nil)
		res := httptest.NewRecorder()

		server.ServeHTTP(res, req)

		assert.Equal(t, http.StatusOK, res.Code)
		assert.Equal(t, "20", res.Body.String())
	})
	t.Run("returns Floyd's score", func(t *testing.T) {
		req, _ := http.NewRequest(http.MethodGet, "/players/Floyd", nil)
		res := httptest.NewRecorder()

		server.ServeHTTP(res, req)

		assert.Equal(t, http.StatusOK, res.Code)
		assert.Equal(t, "10", res.Body.String())
	})
	t.Run("returns 404 on missing players", func(t *testing.T) {
		req, _ := http.NewRequest(http.MethodGet, "/players/Apollo", nil)
		res := httptest.NewRecorder()

		server.ServeHTTP(res, req)

		assert.Equal(t, http.StatusNotFound, res.Code)
	})
}

func TestStoreWins(t *testing.T) {
	store := StubPlayerStore{
		map[string]int{},
		[]string{},
	}
	server := &server.PlayerServer{Store: &store}
	t.Run("returns Accepted on POST", func(t *testing.T) {
		player := "John Snow"
		req, _ := http.NewRequest(http.MethodPost, "/players/"+player, nil)
		res := httptest.NewRecorder()

		server.ServeHTTP(res, req)

		assert.Equal(t, http.StatusAccepted, res.Code)
		assert.Equal(t, player, store.winCalls[0])
	})

}
