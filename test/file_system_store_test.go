package test

import (
	"github.com/bakigoal/tdd_with_go_app/src/model"
	"github.com/bakigoal/tdd_with_go_app/src/store"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestFileSystemStore(t *testing.T) {
	t.Run("league from a reader", func(t *testing.T) {
		db, cleanDb := createTempFile(t, `[
{"Name": "Cleo", "Wins": 10},
{"Name": "Chris", "Wins": 33}]`)
		defer cleanDb()
		playerStore := store.NewFileSystemPlayerStore(db)

		got := playerStore.GetLeague()
		want := model.League{
			{Name: "Cleo", Wins: 10},
			{Name: "Chris", Wins: 33},
		}
		assert.Equal(t, want, got)

		got = playerStore.GetLeague()
		assert.Equal(t, got, want)
	})
	t.Run("get player score", func(t *testing.T) {
		db, cleanDb := createTempFile(t, `[
{"Name": "Cleo", "Wins": 10},
{"Name": "Chris", "Wins": 33}]`)
		defer cleanDb()
		playerStore := store.NewFileSystemPlayerStore(db)

		assert.Equal(t, 33, playerStore.GetPlayerScore("Chris"))
		assert.Equal(t, 10, playerStore.GetPlayerScore("Cleo"))
	})

	t.Run("store wins for existing players", func(t *testing.T) {
		db, cleanDb := createTempFile(t, `[
{"Name": "Cleo", "Wins": 10},
{"Name": "Chris", "Wins": 33}]`)
		defer cleanDb()
		playerStore := store.NewFileSystemPlayerStore(db)

		playerStore.RecordWin("Chris")

		assert.Equal(t, 34, playerStore.GetPlayerScore("Chris"))
	})

	t.Run("store wins for new players", func(t *testing.T) {
		db, cleanDb := createTempFile(t, `[
{"Name": "Cleo", "Wins": 10},
{"Name": "Chris", "Wins": 33}]`)
		defer cleanDb()
		playerStore := store.NewFileSystemPlayerStore(db)

		playerStore.RecordWin("Pepper")

		assert.Equal(t, 1, playerStore.GetPlayerScore("Pepper"))
	})
}

func createTempFile(t testing.TB, initialData string) (*os.File, func()) {
	t.Helper()
	tmpfile, err := os.CreateTemp("", "db")

	if err != nil {
		t.Fatalf("could not create temp file %v", err)
	}

	tmpfile.Write([]byte(initialData))

	removeFile := func() {
		tmpfile.Close()
		os.Remove(tmpfile.Name())
	}

	return tmpfile, removeFile
}
