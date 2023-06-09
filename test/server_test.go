package test

import (
	"github.com/bakigoal/tdd_with_go_app/src/model"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/bakigoal/tdd_with_go_app/src/server"
	"github.com/stretchr/testify/assert"
)

type StubPlayerStore struct {
	scores   map[string]int
	winCalls []string
	league   model.League
}

func (ps *StubPlayerStore) GetPlayerScore(player string) int {
	return ps.scores[player]
}

func (ps *StubPlayerStore) RecordWin(player string) {
	ps.winCalls = append(ps.winCalls, player)
}

func (ps *StubPlayerStore) GetLeague() model.League {
	return ps.league
}

func TestGETPlayers(t *testing.T) {
	store := StubPlayerStore{
		scores: map[string]int{
			"Pepper": 20,
			"Floyd":  10,
		},
		winCalls: []string{},
	}
	server := server.NewPlayerServer(&store)
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
		scores:   map[string]int{},
		winCalls: []string{},
	}
	server := server.NewPlayerServer(&store)
	t.Run("returns Accepted on POST", func(t *testing.T) {
		player := "John Snow"
		req, _ := http.NewRequest(http.MethodPost, "/players/"+player, nil)
		res := httptest.NewRecorder()

		server.ServeHTTP(res, req)

		assert.Equal(t, http.StatusAccepted, res.Code)
		assert.Equal(t, player, store.winCalls[0])
	})

}

func TestLeague(t *testing.T) {
	wantedLeague := model.League{
		{"Johny", 22},
		{"Brad", 32},
		{"Baki", 42},
	}
	store := StubPlayerStore{
		league: wantedLeague,
	}
	playerServer := server.NewPlayerServer(&store)
	t.Run("returns 200 on /league", func(t *testing.T) {
		req, _ := http.NewRequest(http.MethodGet, "/league", nil)
		res := httptest.NewRecorder()

		playerServer.ServeHTTP(res, req)

		got := getLeagueResponse(t, res)
		assert.Equal(t, http.StatusOK, res.Code)
		assert.Equal(t, wantedLeague, got)
		assertHeader(t, res, "content-type", server.JsonContentType)
	})

}

func assertHeader(t *testing.T, res *httptest.ResponseRecorder, header string, expectedValue string) {
	t.Helper()
	assert.Equal(t, expectedValue, res.Result().Header.Get(header))
}

func getLeagueResponse(t *testing.T, res *httptest.ResponseRecorder) model.League {
	t.Helper()
	league, err := model.NewLeague(res.Body)
	assert.NoError(t, err)
	return league
}
