package server

import (
	"bla/models"
	"bla/storage"
	"encoding/json"
	"fmt"
	"github.com/gofrs/uuid"
	"github.com/julienschmidt/httprouter"
	"io/ioutil"
	"net/http"
)

func handleGetAllPosts(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	posts, err := storage.GetAllPosts()
	if check(err, 404, "Posts not found", w) {
		return
	}

	jsonPosts, err := json.Marshal(posts)
	if check(err, 500, "Server error loading posts", w) {
		return
	}

	fmt.Fprintf(w, string(jsonPosts))
}

func handleGetAllFavoritePosts(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	posts, err := storage.GetAllFavoritePosts()
	if check(err, 404, "Posts not found", w) {
		return
	}

	jsonPosts, err := json.Marshal(posts)
	if check(err, 500, "Server error loading post", w) {
		return
	}

	fmt.Fprintf(w, string(jsonPosts))
}

func handleGetPostById(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	post, err := storage.GetPostById(p.ByName("id"))
	if check(err, 404, "Couldn't get post", w) {
		return
	}

	jsonPost, err := json.Marshal(post)

	fmt.Fprintf(w, string(jsonPost))
}

func handleGetPostByTag(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	post, err := storage.GetPostByTag(p.ByName("tag"))
	if check(err, 404, "Couldn't get post", w) {
		return
	}

	jsonPost, err := json.Marshal(post)

	fmt.Fprintf(w, string(jsonPost))
}

func handleCreatePost(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	body, err := ioutil.ReadAll(r.Body)
	if check(err, 400, "Request body failed to be read", w) {
		return
	}

	post := models.Post{}

	err = json.Unmarshal(body, &post)
	if check(err, 422, "Request body could not be parsed", w) {
		return
	}

	id, err := uuid.NewV4()
	if check(err, 500, "Error generating unique post ID", w) {
		return
	}

	post.Id = id

	err = storage.CreatePost(&post)
	if check(err, 500, "Internal server error", w) {
		return
	}
}

func handleUpdatePost(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	body, err := ioutil.ReadAll(r.Body)
	if check(err, 400, "Request body failed to be read", w) {
		return
	}

	post := models.Post{}

	err = json.Unmarshal(body, &post)
	if check(err, 422, "Request body could not be parsed", w) {
		return
	}

	err = storage.UpdatePost(&post)
	if check(err, 500, "Internal server error", w) {
		return
	}
}
