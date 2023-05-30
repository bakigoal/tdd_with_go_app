package test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/bakigoal/tdd_with_go_app/server"
	"github.com/stretchr/testify/assert"
)

type StubPlayerStore struct {
	scores map[string]int
}

func (ps *StubPlayerStore) GetPlayerScore(player string) int {
	return ps.scores[player]
}

func TestGETPlayers(t *testing.T) {
	store := StubPlayerStore{
		map[string]int{
			"Pepper": 20,
			"Floyd":  10,
		},
	}
	server := &server.PlayerServer{Store: &store}
	t.Run("returns Pepper's score", func(t *testing.T) {
		req, _ := http.NewRequest(http.MethodGet, "/players/Pepper", nil)
		res := httptest.NewRecorder()

		server.ServeHTTP(res, req)

		assert.Equal(t, "20", res.Body.String())
	})
	t.Run("returns Floyd's score", func(t *testing.T) {
		req, _ := http.NewRequest(http.MethodGet, "/players/Floyd", nil)
		res := httptest.NewRecorder()

		server.ServeHTTP(res, req)

		assert.Equal(t, "10", res.Body.String())
	})
}
