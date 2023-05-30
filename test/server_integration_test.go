package test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/bakigoal/tdd_with_go_app/server"
	"github.com/bakigoal/tdd_with_go_app/store"
	"github.com/stretchr/testify/assert"
)

func TestRecordingWinsAndGettingThem(t *testing.T) {
	store := store.NewInMemoryPlayerStore()
	server := server.PlayerServer{Store: store}
	player := "Bobby"

	server.ServeHTTP(httptest.NewRecorder(), newPostWinRequest(player))
	server.ServeHTTP(httptest.NewRecorder(), newPostWinRequest(player))
	server.ServeHTTP(httptest.NewRecorder(), newPostWinRequest(player))

	response := httptest.NewRecorder()
	server.ServeHTTP(response, newGetScoreRequest(player))

	assert.Equal(t, http.StatusOK, response.Code)
	assert.Equal(t, "3", response.Body.String())
}

func newPostWinRequest(player string) *http.Request {
	return httptest.NewRequest(http.MethodPost, "/players/"+player, nil)
}

func newGetScoreRequest(player string) *http.Request {
	return httptest.NewRequest(http.MethodGet, "/players/"+player, nil)
}
