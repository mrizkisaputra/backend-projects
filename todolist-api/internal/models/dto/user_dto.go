package dto

type RegisterUserRequestBody struct {
	Name     string `json:"name" validate:"required,max=100"`
	Email    string `json:"email" validate:"required,email,max=100"`
	Password string `json:"password" validate:"required,min=6,max=100"`
}

type LoginUserRequestBody struct {
	Email    string `json:"email" validate:"required,email,max=100"`
	Password string `json:"password" validate:"required,min=6,max=100"`
}

type TokenUserLoginResponse struct {
	Token string `json:"token"`
}

type UserResponse struct {
	Id        string `json:"id"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	CreatedAt int64  `json:"created_at"`
}

type VerifyUserRequest struct {
	AccessToken string `validate:"required,max=100"`
}
