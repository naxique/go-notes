package main

import (
	"log"
	"net/http"
	Storage "notes/db"
	handlers "notes/handlers"
	middleware "notes/middleware"

	"github.com/gofor-little/env"
	"github.com/gorilla/mux"
)

func ServerSetup() *http.Server {
	router := mux.NewRouter()

	var (
		db      Storage.Database
		handler handlers.Handlers
	)

	db.InitStorage()
	handler.SetStorage(&db)

	router.Use(middleware.LoggerMiddleware)

	secret, err := env.MustGet("JWT_SECRET")
	if err != nil {
		log.Fatalln("JWT_SECRET doesn't exist", err)
	}

	jwt := new(middleware.JWT)
	jwt.InitJWT(secret)
	handler.SetJWT(jwt)

	unprotected := router.NewRoute().Subrouter()

	unprotected.HandleFunc("/api/status", handlers.StatusHandler).Methods("GET")

	unprotected.HandleFunc("/api/user/signup", handler.UserSignupHandler).Methods("POST")
	unprotected.HandleFunc("/api/user/login", handler.UserLoginHandler).Methods("POST")
	unprotected.HandleFunc("/api/user/{userId}", handler.UserGetHandler).Methods("GET")
	unprotected.HandleFunc("/api/user/delete/{userId}", handler.UserDeleteHandler).Methods("DELETE")

	protected := router.NewRoute().Subrouter()
	protected.Use(jwt.AuthMiddleware)
	protected.HandleFunc("/api/user/logout/{userId}", handler.UserLogoutHandler).Methods("POST")

	protected.HandleFunc("/api/note/addnote/{userId}", handler.NoteAddHandler).Methods("POST")
	protected.HandleFunc("/api/note/getallnotes/{userId}", handler.NoteGetAllHandler).Methods("GET")
	protected.HandleFunc("/api/note/getnote/{noteId}", handler.NoteGetHandler).Methods("GET")
	protected.HandleFunc("/api/note/editnote/{noteId}", handler.NoteEditHandler).Methods("PATCH")
	protected.HandleFunc("/api/note/deletenote/{noteId}", handler.NoteDeleteHandler).Methods("DELETE")

	return &http.Server{
		Addr:    ":8080",
		Handler: router,
	}
}
