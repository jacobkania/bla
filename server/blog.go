package server

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func SetupBlogRouter(router *httprouter.Router) {
	router.GET("/", handleIndex)
}

func handleIndex(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	http.ServeFile(w, r, "./content/static/index.html")
}
