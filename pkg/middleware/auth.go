package middleware

import (
	"context"
	"net/http"

	"github.com/md-muzzammil-rashid/blog-app-backend-go/internal/auth"
	"github.com/md-muzzammil-rashid/blog-app-backend-go/pkg/utils"
)

func VerifyJWT (next http.HandlerFunc, authRepo auth.AuthRepository) http.HandlerFunc{
	return func (w http.ResponseWriter, r *http.Request){
		cookie, err := r.Cookie("access_token")
		authToken := r.Header.Get("Authorization")
		if len(authToken)>7 {
			authToken = authToken[7:]
		}
		if err != nil && r.Header.Get("Authorization") == ""{

            utils.WriteError(w, http.StatusUnauthorized, "Access token is not provided")
			return
		}

		var userId string
		if err == nil  {
			userId, err = utils.GetUserIdFromToken(cookie.Value ); if err != nil {
				utils.WriteError(w, http.StatusUnauthorized, "Invalid access token")
                return
			}
		} else {
			userId, err = utils.GetUserIdFromToken(authToken); if err != nil {
				utils.WriteError(w, http.StatusUnauthorized, "Invalid access token")
                return
			}
		}
		if len(userId) == 0 {
			utils.WriteError(w, http.StatusUnauthorized, "User not found")
            return
		}
		_,err = authRepo.GetUserByID(userId); if err != nil {
			utils.WriteError(w, http.StatusUnauthorized, "Invalid access token")
            return
		}
		ctx := context.WithValue(r.Context(), "user_id", userId)
		next.ServeHTTP(w, r.WithContext(ctx))
		

	}
}