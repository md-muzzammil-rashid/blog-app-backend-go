package auth

import (
	"errors"
	"log/slog"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v5"
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


func LoginUserService(authRepo *AuthRepository, email string, password string) (LoginUserResponseDTO, error) {
	var user UserModel
	user, err := authRepo.GetUserByEmail(email); if err != nil {
		return LoginUserResponseDTO{}, err
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); if err != nil {
		return LoginUserResponseDTO{} , errors.New("Invalid credentials")
	}

	userResponseWithToken := LoginUserResponseDTO{UserId: user.UserId}
	GenerateAccessAndRefreshToken(user, &userResponseWithToken)
	return userResponseWithToken, nil
}

func GenerateAccessAndRefreshToken(userDetails UserModel, user *LoginUserResponseDTO) error {
	// expiry, err := os.LookupEnv("JWT_EXPIRY"); if err != nil {
	// 	return errors.New("Expiry not found in environment variable")
	// }
	secret, ok := os.LookupEnv("JWT_SECRET"); if !ok {

		return errors.New("SECRET not found in environment variable")
	}

	claims := jwt.MapClaims{
		"user_id": userDetails.UserId,
		"email" : userDetails.Email,
		"exp": time.Now().Add(time.Hour * 24 * 7 ).Unix(),
	}
	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(secret))
	if err != nil {
		return errors.New("failed to sign token")
	}
    user.AccessToken = string(token)
	return nil
}
