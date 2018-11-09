package server

import (
	"bla/models"
	"bla/storage"
	"database/sql"
	"encoding/json"
	"github.com/gofrs/uuid"
	"github.com/julienschmidt/httprouter"
	"gopkg.in/russross/blackfriday.v2"
	"io/ioutil"
	"net/http"
)

func handleGetAllPosts(db *sql.DB) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		posts, err := storage.GetAllPosts(db)
		if check(err, 404, "Posts not found", w) {
			return
		}

		jsonPosts, err := json.Marshal(&posts)
		if check(err, 500, "Server error loading posts", w) {
			return
		}

		writeResponse(200, &jsonPosts, w)
	}
}

func handleGetAllFavoritePosts(db *sql.DB) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		posts, err := storage.GetAllFavoritePosts(db)
		if check(err, 404, "Posts not found", w) {
			return
		}

		jsonPosts, err := json.Marshal(&posts)
		if check(err, 500, "Server error loading posts", w) {
			return
		}

		writeResponse(200, &jsonPosts, w)
	}
}

func handleGetPostById(db *sql.DB) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		post, err := storage.GetPostById(db, uuid.FromStringOrNil(p.ByName("id")))
		if check(err, 404, "Couldn't get post", w) {
			return
		}

		jsonPost, err := json.Marshal(&post)
		if check(err, 500, "Server error loading post", w) {
			return
		}

		writeResponse(200, &jsonPost, w)
	}
}

func handleGetPostByTag(db *sql.DB) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		post, err := storage.GetPostByTag(db, p.ByName("tag"))
		if check(err, 404, "Couldn't get post", w) {
			return
		}

		jsonPost, err := json.Marshal(&post)
		if check(err, 500, "Server error loading post", w) {
			return
		}

		writeResponse(200, &jsonPost, w)
	}
}

func handleCreatePost(db *sql.DB) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		body, err := ioutil.ReadAll(r.Body)
		if check(err, 400, "Request body failed to be read", w) {
			return
		}

		post := models.Post{}

		err = json.Unmarshal(body, &post)
		if check(err, 422, "Request body could not be parsed", w) {
			return
		}

		post.ContentHTML = string(blackfriday.Run([]byte(post.ContentMD)))

		completedPost, err := storage.CreatePost(db, post.Tag, post.Title, post.ContentMD, post.ContentHTML, post.Published, post.Edited, post.IsFavorite, post.Author)
		if check(err, 500, "Internal server error", w) {
			return
		}

		completedPostJson, err := json.Marshal(completedPost)
		if check(err, 500, "JSON parsing error, data may be corrupted", w) {
			return
		}

		writeResponse(201, &completedPostJson, w)
	}
}

func handleUpdatePost(db *sql.DB) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		body, err := ioutil.ReadAll(r.Body)
		if check(err, 400, "Request body failed to be read", w) {
			return
		}

		post := models.Post{}

		err = json.Unmarshal(body, &post)
		if check(err, 422, "Request body could not be parsed", w) {
			return
		}

		post.ContentHTML = string(blackfriday.Run([]byte(post.ContentMD)))

		completedPost, err := storage.UpdatePost(db, uuid.FromStringOrNil(p.ByName("id")), post.Tag, post.Title, post.ContentMD, post.ContentHTML, post.Published, post.Edited, post.IsFavorite, post.Author)
		if check(err, 500, "Internal server error", w) {
			return
		}

		completedPostJson, err := json.Marshal(completedPost)
		if check(err, 500, "JSON parsing error, data may be corrupted", w) {
			return
		}

		writeResponse(200, &completedPostJson, w)
	}
}
