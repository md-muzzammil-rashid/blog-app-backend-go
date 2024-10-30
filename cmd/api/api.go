package api

import (
	// "encoding/json"
	"database/sql"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/md-muzzammil-rashid/blog-app-backend-go/internal/auth"
	"github.com/md-muzzammil-rashid/blog-app-backend-go/pkg/utils"
	// "github.com/md-muzzammil-rashid/blog-app-backend-go/pkg/utils"
	// "github.com/md-muzzammil-rashid/blog-app-backend-go/pkg/utils"
)

func ApiHandler (router *mux.Router, db *sql.DB) {

	authRepo := auth.NewAuthRepository(db)

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        utils.WriteJSON(w, http.StatusOK, "API is working", "Welcome to the blog API")
	}).Methods("GET")

	v1 := router.PathPrefix("/api/v1").Subrouter()

	authRoute := v1.PathPrefix("/auth").Subrouter()

	authRoute.HandleFunc("/register", auth.RegisterUser(authRepo)).Methods(http.MethodPost)

	authRoute.HandleFunc("/login", auth.LoginUser(authRepo)).Methods(http.MethodGet)
}