package server

import (
	"fmt"
	"log"
	"net/http"
)

func check(err error, statusCode int, message string, w http.ResponseWriter) bool {
	if err != nil {
		log.Printf("Error: %v", err)
		w.WriteHeader(statusCode)
		fmt.Fprint(w, message)
		return true
	}
	return false
}
