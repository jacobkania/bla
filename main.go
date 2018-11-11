package main

import (
	"bla/initialize"
	"bla/server"
	"database/sql"
	"github.com/julienschmidt/httprouter"
	"log"
)

func main() {
	router := httprouter.New()

	db, err := sql.Open("sqlite3", "./content/data/bla.db")
	if err != nil {
		log.Fatalf("Database failed to open")
	}

	err = initialize.Initialize(db)
	if err != nil {
		log.Fatalf("Failed to run initialization")
	}

	srv := server.Server{
		Router: router,
		Db:     db,
	}

	srv.SetRoutes()
	srv.NewServer(":8081", ":8080")

	log.Fatal(srv.Run())
}
