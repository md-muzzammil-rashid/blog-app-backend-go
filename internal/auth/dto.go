package auth

type RegisterUserDTO struct {
	Username string `json:"username" validate:"required"`
	Email string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
	ConfirmPassword string `json:"confirm_password" validate:"required"`
	DisplayName string `json:"display_name" validate:"required"`
}

type LoginUserDTO struct {
    Email string `json:"email"`
    Password string `json:"password"`
}


type LoginUserResponseDTO struct {
	UserId string `json:"user_id"`
	AccessToken string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}