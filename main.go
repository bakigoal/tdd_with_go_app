package main

import (
	"log"
	"net/http"

	"github.com/bakigoal/tdd_with_go_app/server"
	"github.com/bakigoal/tdd_with_go_app/service"
)

func main() {
	server := &server.PlayerServer{
		Store: &service.PlayerService{},
	}
	log.Fatal(http.ListenAndServe(":8888", server))
}
