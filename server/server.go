package server

import (
	"fmt"
	"net/http"
	"strings"
)

func PlayerServer(w http.ResponseWriter, req *http.Request) {
	player := strings.TrimPrefix(req.URL.Path, "/players/")
	if player == "Pepper" {
		fmt.Fprint(w, 20)
		return
	}
	if player == "Floyd" {
		fmt.Fprint(w, 10)
		return
	}
}
