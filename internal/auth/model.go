package auth

import "time"

type UserModel struct {
	UserId string `json:"user_id"`
	Username string `json:"username"`
	Email string `json:"email"`
	DisplayName string `json:"display_name"`
	Password string `json:"password"`
	CreatedAt time.Time `json:"created_at"`
}