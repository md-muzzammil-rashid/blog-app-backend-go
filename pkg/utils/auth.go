package utils

import (
	"errors"
	"os"

	"github.com/golang-jwt/jwt/v5"
)

func GetUserIdFromToken(tokenString string)(string, error) {
	token, err:= jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC); if !ok {
			return nil, errors.New("Invalid method signature")
		}
		return []byte(os.Getenv("JWT_SECRET")), nil
	})
	if err!= nil {
        return "", err
    }
	if !token.Valid {
		return "" , errors.New("invalid token")
	}
	
	claims := token.Claims.(jwt.MapClaims)
	userId := claims["user_id"].(string) 
	return userId, nil
}

