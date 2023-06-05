package main

import (
	"log"
	"net/http"
	"os"

	"github.com/bakigoal/tdd_with_go_app/src/server"
	"github.com/bakigoal/tdd_with_go_app/src/store"
)

const dbFileName = "game.db.json"

func main() {
	db, err := os.OpenFile(dbFileName, os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		log.Fatalf("problem opening %s %v", dbFileName, err)
	}

	playerStore := store.NewFileSystemPlayerStore(db)
	server := server.NewPlayerServer(playerStore)

	log.Fatal(http.ListenAndServe(":8888", server))
}
