package auth

import (
	"database/sql"
	"log/slog"
	"net/http"

	"github.com/md-muzzammil-rashid/blog-app-backend-go/pkg/utils"
)

func RegisterUser(authRepo *AuthRepository) http.HandlerFunc {
	return func (w http.ResponseWriter, r *http.Request){
		var userDetails RegisterUserDTO
		err:=utils.ReadJSON(*r, &userDetails); if err != nil {
			slog.Info(err.Error())
			utils.WriteError(w, http.StatusBadRequest, "Invalid JSON : "+ err.Error() )
			return
		}

		statusCode, err := RegisterUserService(&userDetails, authRepo); if err != nil {
			utils.WriteError(w, statusCode, err.Error())
			return
		}

		utils.WriteJSON(w, statusCode, "User has been successfully register", userDetails)
	}
}

func LoginUser(authRepo *AuthRepository) http.HandlerFunc {
	return func (w http.ResponseWriter, r *http.Request) {
		var loginUser LoginUserDTO
		err := utils.ReadJSON(*r, &loginUser); if err != nil {
			slog.Info(err.Error())
			utils.WriteError(w, http.StatusBadRequest, "Invalid JSON : "+ err.Error() )
			return
		}

		user, err := LoginUserService(authRepo, loginUser.Email, loginUser.Password); if err != nil {
			if err.Error() == "Invalid credentials" {
				utils.WriteError(w, http.StatusBadRequest, err.Error())
				return
			}else if err==sql.ErrNoRows {
				utils.WriteError(w, http.StatusNotFound, "User not found")
                return
			}
			utils.WriteError(w, http.StatusInternalServerError, err.Error())
			return
		}
		http.SetCookie(w, &http.Cookie{Name: "access_token", Value: user.AccessToken})
		utils.WriteJSON(w, http.StatusOK, "User has been successfully logged in", user)

	}
}