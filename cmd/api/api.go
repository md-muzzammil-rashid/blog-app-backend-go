package api

import (
	// "encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/md-muzzammil-rashid/blog-app-backend-go/pkg/utils"
	// "github.com/md-muzzammil-rashid/blog-app-backend-go/pkg/utils"
	// "github.com/md-muzzammil-rashid/blog-app-backend-go/pkg/utils"
)

func ApiHandler (router *mux.Router) {

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        utils.WriteJSON(w, http.StatusOK, "Welcome to the blog API")
	}).Methods("GET")

	// v1 := router.PathPrefix("/api/v1").Subrouter()

	// authRoute := v1.PathPrefix("/auth").Subrouter()

	// auth

}