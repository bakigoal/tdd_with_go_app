package server

import (
	"fmt"
	"net/http"
)

func PlayerServer(w http.ResponseWriter, req *http.Request) {
	fmt.Fprint(w, 20)
}
