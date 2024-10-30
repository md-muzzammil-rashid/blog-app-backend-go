package auth

import (
	"database/sql"
)

type UserModel struct {
	UserId string `json:"user_id"`
	Username string `json:"username"`
	Email string `json:"email"`
	DisplayName string `json:"display_name"`
	Password string `json:"password"`
	CreatedAt sql.NullTime `json:"created_at"`
}