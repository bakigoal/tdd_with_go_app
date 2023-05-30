package main

import (
	"log"
	"net/http"

	"github.com/bakigoal/tdd_with_go_app/server"
)

func main() {
	handler := http.HandlerFunc(server.PlayerServer)
	log.Fatal(http.ListenAndServe(":8888", handler))
}
