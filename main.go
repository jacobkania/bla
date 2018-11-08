package main

import (
	"bla/initialize"
	"bla/server"
	"github.com/julienschmidt/httprouter"
	"log"
)

func main() {
	err := initialize.Initialize()
	if err != nil {
		log.Panicf("Aww, fail.")
	}

	router := httprouter.New()
	server.SetupBlogRouter(router)

	srv := server.NewServer(":8081", router)
	log.Fatal(srv.ListenAndServe())
}
