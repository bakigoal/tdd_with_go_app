package test

import (
	"github.com/bakigoal/tdd_with_go_app/server"
	"github.com/bakigoal/tdd_with_go_app/store"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestFileSystemStore(t *testing.T) {
	db := strings.NewReader(`[
{"Name": "Cleo", "Wins": 10},
{"Name": "Chris", "Wins": 33}]`)
	store := store.FileSystemPlayerStore{Database: db}
	t.Run("league from a reader", func(t *testing.T) {
		got := store.GetLeague()
		want := []server.Player{
			{Name: "Cleo", Wins: 10},
			{Name: "Chris", Wins: 33},
		}
		assert.Equal(t, want, got)

		got = store.GetLeague()
		assert.Equal(t, got, want)
	})
	t.Run("get player score", func(t *testing.T) {
		assert.Equal(t, 33, store.GetPlayerScore("Chris"))
		assert.Equal(t, 10, store.GetPlayerScore("Cleo"))
	})
}
