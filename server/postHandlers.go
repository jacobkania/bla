package server

import (
	"bla/storage"
	"encoding/json"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func handleGetAllPosts(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	posts, err := storage.GetAllPosts()
	if Check(err, w) {
		return
	}

	jsonPosts, err := json.Marshal(posts)
	if Check(err, w) {
		return
	}

	fmt.Fprintf(w, string(jsonPosts))
}

func handleGetAllFavoritePosts(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	posts, err := storage.GetAllFavoritePosts()
	if Check(err, w) {
		return
	}

	jsonPosts, err := json.Marshal(posts)
	if Check(err, w) {
		return
	}

	fmt.Fprintf(w, string(jsonPosts))
}

func handleGetPostById(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	post, err := storage.GetPostById(p.ByName("id"))
	if Check(err, w) {
		return
	}

	jsonPost, err := json.Marshal(post)

	fmt.Fprintf(w, string(jsonPost))
}

func handleGetPostByTag(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	post, err := storage.GetPostByTag(p.ByName("tag"))
	if Check(err, w) {
		return
	}

	jsonPost, err := json.Marshal(post)

	fmt.Fprintf(w, string(jsonPost))
}

func handleCreatePost(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

}

func handleUpdatePost(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

}
