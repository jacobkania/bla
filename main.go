package main

import (
	"github.com/jacobkania/bla/configuration"
	"github.com/jacobkania/bla/server"

	"database/sql"
	"github.com/julienschmidt/httprouter"
	"log"
)

func main() {
	config := configuration.Load()

	router := httprouter.New()

	db, err := sql.Open("sqlite3", "./content/data/bla.db")
	if err != nil {
		log.Fatalf("Database failed to open")
	}

	err = Initialize(db)
	if err != nil {
		log.Fatalf("Failed to run initialization")
	}

	srv := server.Server{
		Config: config,
		Router: router,
		Db:     db,
	}

	log.Fatal(srv.Run())
}
