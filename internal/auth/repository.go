package auth

import (
	"database/sql"
	"strings"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type Authenticator interface {
	RegisterUser(user UserModel) error
	LoginUser(email, password string) (LoginUserResponseDTO, error)
	// GetUserByID(id string) (*UserModel, error)
	// UpdateUser(user UserModel) error
	// DeleteUser() error


}
type AuthRepository struct {
	Db *sql.DB
}

func NewAuthRepository(db *sql.DB ) *AuthRepository {
	
    return &AuthRepository{Db: db}
}

func (m *AuthRepository) RegisterUser(userData RegisterUserDTO) error {
	userId := uuid.NewString()
	stmt,err := m.Db.Prepare("INSERT INTO users (user_id, username, email, password, display_name) VALUES (?, ?, ?, ?, ?)"); if err != nil {
		return err
	}
	_ , err =stmt.Exec(userId, strings.ToLower(userData.Username), strings.ToLower(userData.Email), userData.Password, userData.DisplayName); if err != nil {
		return err
	}

	return nil

}

func (m *AuthRepository) LoginUser(email, password string) (LoginUserResponseDTO, error) {
	query := "SELECT * FROM users where email = ?"
	var user UserModel

	err := m.Db.QueryRow(query, strings.ToLower(email)).Scan(&user); if err != nil {
		return LoginUserResponseDTO{}, err
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); if err != nil {
		return LoginUserResponseDTO{}, err
	}

	return LoginUserResponseDTO{UserId: user.UserId, AccessToken: "this is accesstoken", RefreshToken: "This is refresh token" }, nil
	
}