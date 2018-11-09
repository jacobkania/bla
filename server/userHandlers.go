package server

import (
	"bla/storage"
	"database/sql"
	"encoding/json"
	"github.com/gofrs/uuid"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func handleGetAllUsers(db *sql.DB) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		users, err := storage.GetAllUsers(db)
		if check(err, 404, "Users not found", w) {
			return
		}

		jsonUsers, err := json.Marshal(&users)
		if check(err, 500, "Server error loading users", w) {
			return
		}

		writeResponse(200, &jsonUsers, w)
	}
}

func handleGetUserById(db *sql.DB) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		user, err := storage.GetUserById(db, uuid.FromStringOrNil(p.ByName("id")))
		if check(err, 404, "Users not found", w) {
			return
		}

		jsonUser, err := json.Marshal(&user)
		if check(err, 500, "Server error loading user", w) {
			return
		}

		writeResponse(200, &jsonUser, w)
	}
}
