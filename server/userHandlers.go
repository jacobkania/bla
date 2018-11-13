package server

import (
	"bla/authentication"
	"bla/models"
	"bla/storage"
	"database/sql"
	"encoding/json"
	"github.com/gofrs/uuid"
	"github.com/julienschmidt/httprouter"
	"io/ioutil"
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

func handleLogin(db *sql.DB) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		body, err := ioutil.ReadAll(r.Body)
		if check(err, 400, "Request body failed to be read", w) {
			return
		}

		login := struct {
			Username string `json:"username"`
			Password string `json:"password"`
		}{}

		err = json.Unmarshal(body, &login)
		if check(err, 422, "Request body could not be parsed", w) {
			return
		}

		user, err := storage.GetUserByPersonalLogin(db, login.Username)
		if err != nil {
			user = &models.User{HashedPw: ""}
		}

		answer := authentication.CheckPassword(user.HashedPw, login.Password)

		var answerText []byte

		if answer {
			answerText = []byte("true")
		} else {
			answerText = []byte("false")
		}

		writeResponse(200, &answerText, w)
	}
}
