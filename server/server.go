package server

import (
	"bla/configuration"
	"crypto/tls"
	"database/sql"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"strconv"
	"time"
)

type Server struct {
	Config      *configuration.Configuration
	Router      *httprouter.Router
	Db          *sql.DB
	HttpsServer *http.Server
	HttpServer  *http.Server
}

func (s *Server) Run() error {
	httpUrl := s.Config.ServerUrl + ":" + strconv.Itoa(s.Config.HttpPort)
	httpsUrl := s.Config.ServerUrl + ":" + strconv.Itoa(s.Config.HttpsPort)

	s.setRoutes()
	s.newServer(httpsUrl, httpUrl)

	go s.HttpServer.ListenAndServe()
	return s.HttpsServer.ListenAndServeTLS(s.Config.CertFile, s.Config.KeyFile)
}

func (s *Server) newServer(tlsServerAddress string, redirectServerAddress string) {
	s.HttpsServer = &http.Server{
		Addr:         tlsServerAddress,
		Handler:      s.Router,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
		TLSConfig: &tls.Config{
			PreferServerCipherSuites: true,
			CurvePreferences:         []tls.CurveID{tls.CurveP256, tls.X25519},
			MinVersion:               tls.VersionTLS12,
			CipherSuites: []uint16{
				tls.TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384,
				tls.TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384,
				tls.TLS_ECDHE_ECDSA_WITH_CHACHA20_POLY1305,
				tls.TLS_ECDHE_RSA_WITH_CHACHA20_POLY1305,
				tls.TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256,
				tls.TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256,
			},
		},
	}

	s.HttpServer = &http.Server{
		Addr:         redirectServerAddress,
		Handler:      redirectToHttps(tlsServerAddress),
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
	}
}

func (s *Server) setRoutes() {
	// HTML pages
	s.Router.GET("/", handleIndex())
	s.Router.GET("/page/:tag", handlePage())
	s.Router.GET("/favicon.ico", handleFavicon())
	s.Router.ServeFiles("/p/*filepath", http.Dir("./content/pages"))
	s.Router.ServeFiles("/css/*filepath", http.Dir("./content/static/css"))
	s.Router.ServeFiles("/js/*filepath", http.Dir("./content/static/js"))
	s.Router.ServeFiles("/img/*filepath", http.Dir("./content/static/img"))
	s.Router.ServeFiles("/image/*filepath", http.Dir("./content/image"))

	// Posts
	s.Router.GET("/post", handleGetAllPosts(s.Db))
	s.Router.GET("/favorites", handleGetAllFavoritePosts(s.Db))
	s.Router.GET("/post/id/:id", handleGetPostById(s.Db))
	s.Router.GET("/post/tag/:tag", handleGetPostByTag(s.Db))
	s.Router.POST("/post", handleCreatePost(s.Db))
	s.Router.PUT("/post/id/:id", handleUpdatePost(s.Db))

	// Users
	s.Router.GET("/user", handleGetAllUsers(s.Db))
	s.Router.GET("/user/id/:id", handleGetUserById(s.Db))
	s.Router.PUT("/user", handleLogin(s.Db))

	// Images

	// Admin
	s.Router.GET("/admin", handleAdminPage())
	s.Router.GET("/page/:tag/admin", handleAdminPage())
}

func handleIndex() httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		http.ServeFile(w, r, "./content/static/index.html")
	}
}

func handlePage() httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		http.ServeFile(w, r, "./content/static/page.html")
	}
}

func handleFavicon() httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		http.ServeFile(w, r, "./content/static/img/favicon.ico")
	}
}

func handleAdminPage() httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		http.ServeFile(w, r, "./content/static/admin.html")
	}
}

func redirectToHttps(url string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, url+r.RequestURI, http.StatusMovedPermanently)
	}
}
