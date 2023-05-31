package main

import (
	"log"
	"net/http"

	"github.com/bakigoal/tdd_with_go_app/server"
	"github.com/bakigoal/tdd_with_go_app/store"
)

func main() {
	server := server.NewPlayerServer(store.NewInMemoryPlayerStore())
	log.Fatal(http.ListenAndServe(":8888", server))
}
