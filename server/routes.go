package server

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func SetupBlogRouter(router *httprouter.Router) {
	// HTML pages
	router.GET("/", handleIndex)

	// Posts
	router.GET("/post", handleGetAllPosts)
	router.GET("/favorites", handleGetAllFavoritePosts)
	router.GET("/post/id/:id", handleGetPostById)
	router.GET("/post/tag/:tag", handleGetPostByTag)
	router.POST("/post", handleCreatePost)
	router.PUT("/post/id/:id", handleUpdatePost)

	// Users

	// Images
}

func handleIndex(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	http.ServeFile(w, r, "./content/static/index.html")
}
