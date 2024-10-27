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
	Token string `json:"token"`
}