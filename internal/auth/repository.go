package auth

import (
	"database/sql"
	"strings"

	"github.com/google/uuid"
)

type Authenticator interface {
	RegisterUser(user UserModel) error
	GetUserByEmail(email string) (UserModel, error)
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

func (a *AuthRepository) RegisterUser(userData RegisterUserDTO) error {
	userId := uuid.NewString()
	stmt,err := a.Db.Prepare("INSERT INTO users (user_id, username, email, password, display_name) VALUES (?, ?, ?, ?, ?)"); if err != nil {
		return err
	}
	_ , err =stmt.Exec(userId, strings.ToLower(userData.Username), strings.ToLower(userData.Email), userData.Password, userData.DisplayName); if err != nil {
		return err
	}

	return nil

}

func (a *AuthRepository) GetUserByEmail(email string) (UserModel, error) {
	var user UserModel
	query := "SELECT * FROM users WHERE email = ?"

	stmt, err := a.Db.Prepare(query)
	if err != nil {
		return UserModel{}, err
	}
	defer stmt.Close() // Ensure the statement is closed after use

	row := stmt.QueryRow(email)
	user, err = ScanRowIntoUser(row)
	if err != nil {
		return UserModel{}, err
	}
	return user, nil
}

func ScanRowIntoUser(row *sql.Row) (UserModel, error) {
	var user UserModel
	err := row.Scan(
		&user.UserId,
		&user.Username,
		&user.Email,
		&user.Password,
		&user.DisplayName,
		&user.CreatedAt, // sql.NullTime should handle NULL values gracefully
	)
	if err != nil {
		return UserModel{}, err
	}
	return user, nil
}