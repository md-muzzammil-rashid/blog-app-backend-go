package auth

import (
	"errors"
	"log/slog"
	"net/http"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/md-muzzammil-rashid/blog-app-backend-go/pkg/utils"
	"golang.org/x/crypto/bcrypt"
)

func RegisterUserService(userDetails *RegisterUserDTO, authRepo *AuthRepository) (int, error) {
	
	if strings.Compare(userDetails.Password, userDetails.ConfirmPassword) != 0 {
		return http.StatusBadRequest, errors.New("Password not matched")
	}

	encryptPassword, err := bcrypt.GenerateFromPassword([]byte(userDetails.Password), 10); if err != nil {
		return http.StatusInternalServerError, errors.New("failed to hashed password before registering")
	}

	userDetails.Password = string(encryptPassword)
	
	err = validator.New().Struct(userDetails); if err != nil {
		return http.StatusBadRequest, utils.ValidationsError(err.(validator.ValidationErrors))
	}

	err = authRepo.RegisterUser(*userDetails); if err != nil {
		return http.StatusInternalServerError, errors.New("failed to register user" + ": " + err.Error())
	}
	slog.Info(userDetails.Username, userDetails.Password, userDetails.ConfirmPassword, userDetails.Email)
	return http.StatusCreated, nil
}

func LoginUserService(emain, password string) {
	
}