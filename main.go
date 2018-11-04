package main

import (
	"bla/initialize"
	"bla/server"
	"github.com/julienschmidt/httprouter"
)

func main() {
	router := httprouter.New()

	initialize.Initialize()

	server.SetupBlogRouter(*router)
}
