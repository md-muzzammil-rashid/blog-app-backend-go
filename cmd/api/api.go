package api

import "github.com/gorilla/mux"

func ApiHandler (router *mux.Router) {
	v1 := router.PathPrefix("/api/v1").Subrouter()

	authRoute := v1.PathPrefix("/auth").Subrouter()

	authRoute.HandleFunc().Methods("POST")
}