package server

import (
	"fmt"
	"log"
	"net/http"
)

func Check(err error, w http.ResponseWriter) bool {
	if err != nil {
		log.Print(err)
		w.WriteHeader(404)
		fmt.Fprint(w, "404: page not found")
		return true
	}
	return false
}
