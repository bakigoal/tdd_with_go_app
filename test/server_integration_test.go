package test

import (
	"github.com/bakigoal/tdd_with_go_app/model"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/bakigoal/tdd_with_go_app/server"
	"github.com/bakigoal/tdd_with_go_app/store"
)

func TestRecordingWinsAndGettingThem(t *testing.T) {
	db, cleanDb := createTempFile(t, "")
	defer cleanDb()
	playerStore := &store.FileSystemPlayerStore{Database: db}
	playerServer := server.NewPlayerServer(playerStore)
	player := "Bobby"

	playerServer.ServeHTTP(httptest.NewRecorder(), newPostWinRequest(player))
	playerServer.ServeHTTP(httptest.NewRecorder(), newPostWinRequest(player))
	playerServer.ServeHTTP(httptest.NewRecorder(), newPostWinRequest(player))

	t.Run("get score", func(t *testing.T) {
		response := httptest.NewRecorder()
		playerServer.ServeHTTP(response, newGetScoreRequest(player))
		assert.Equal(t, http.StatusOK, response.Code)
		assert.Equal(t, "3", response.Body.String())
	})

	t.Run("get league", func(t *testing.T) {
		response := httptest.NewRecorder()
		playerServer.ServeHTTP(response, newLeagueRequest())
		assert.Equal(t, http.StatusOK, response.Code)
		got := getLeagueResponse(t, response)
		want := model.League{
			{"Bobby", 3},
		}
		assert.Equal(t, want, got)
	})
}

func newPostWinRequest(player string) *http.Request {
	return httptest.NewRequest(http.MethodPost, "/players/"+player, nil)
}

func newGetScoreRequest(player string) *http.Request {
	return httptest.NewRequest(http.MethodGet, "/players/"+player, nil)
}

func newLeagueRequest() *http.Request {
	return httptest.NewRequest(http.MethodGet, "/league", nil)
}
