package auth

import (
	"log/slog"
	"net/http"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/md-muzzammil-rashid/blog-app-backend-go/pkg/utils"
	"golang.org/x/crypto/bcrypt"
)

func RegisterUser() http.HandlerFunc {
	return func (w http.ResponseWriter, r *http.Request){
		var userDetails RegisterUserDTO
		err:=utils.ReadJSON(*r, &userDetails); if err != nil {
			slog.Info(err.Error())
			// return
		}
		encryptPassword, err := bcrypt.GenerateFromPassword([]byte(userDetails.Password), 10); if err != nil {
			utils.WriteError(w, http.StatusInternalServerError, "failed to hashed password before registering")
		}
		slog.Info(string(encryptPassword))
		
		if strings.Compare(userDetails.Password, userDetails.ConfirmPassword) != 0 {
			utils.WriteError(w, http.StatusBadRequest, "password not matched" );
			return
		}

		err = validator.New().Struct(userDetails); if err != nil {
			utils.WriteJSON(w, http.StatusBadRequest, utils.ValidationsError((err.(validator.ValidationErrors))))
			return
		}

		

		slog.Info(userDetails.Username, userDetails.Password, userDetails.ConfirmPassword, userDetails.Email)
		utils.WriteJSON(w, http.StatusCreated, userDetails)
	}
}