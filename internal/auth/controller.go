package auth

import (
	"log/slog"
	"net/http"
	"github.com/md-muzzammil-rashid/blog-app-backend-go/pkg/utils"
)

func RegisterUser(authRepo *AuthRepository) http.HandlerFunc {
	return func (w http.ResponseWriter, r *http.Request){
		var userDetails RegisterUserDTO
		err:=utils.ReadJSON(*r, &userDetails); if err != nil {
			slog.Info(err.Error())
			return
		}

		statusCode, err := RegisterUserService(&userDetails, authRepo); if err != nil {
			utils.WriteError(w, statusCode, err.Error())
			return
		}

		utils.WriteJSON(w, statusCode, "User has been successfully register", userDetails)
	}
}