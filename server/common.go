package server

import (
	"fmt"
	"log"
	"net/http"
)

// check is used to determine whether or not an error was encountered. If an error was
// encountered, it writes the given HTTP status code and message to the http.ResponseWriter
// and returns true. If no error occurred, returns false.
func check(err error, statusCode int, message string, w http.ResponseWriter) bool {
	if err != nil {
		log.Printf("Error: %v", err)
		w.WriteHeader(statusCode)
		fmt.Fprint(w, message)
		return true
	}
	return false
}

// writeResponse will write an HTTP status code and a response to the http.ResponseWriter.
func writeResponse(statusCode int, response *[]byte, w http.ResponseWriter) {
	w.WriteHeader(statusCode)
	fmt.Fprint(w, string(*response))
}
