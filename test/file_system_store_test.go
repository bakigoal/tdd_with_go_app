package test

import (
	"github.com/bakigoal/tdd_with_go_app/server"
	"github.com/bakigoal/tdd_with_go_app/store"
	"github.com/stretchr/testify/assert"
	"io"
	"os"
	"testing"
)

func TestFileSystemStore(t *testing.T) {
	t.Run("league from a reader", func(t *testing.T) {
		db, cleanDb := createTempFile(t, `[
{"Name": "Cleo", "Wins": 10},
{"Name": "Chris", "Wins": 33}]`)
		defer cleanDb()
		store := store.FileSystemPlayerStore{Database: db}

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
		db, cleanDb := createTempFile(t, `[
{"Name": "Cleo", "Wins": 10},
{"Name": "Chris", "Wins": 33}]`)
		defer cleanDb()
		store := store.FileSystemPlayerStore{Database: db}

		assert.Equal(t, 33, store.GetPlayerScore("Chris"))
		assert.Equal(t, 10, store.GetPlayerScore("Cleo"))
	})
}

func createTempFile(t testing.TB, initialData string) (io.ReadWriteSeeker, func()) {
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
