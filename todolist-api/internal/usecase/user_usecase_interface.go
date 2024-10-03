package usecase

import (
	"context"
	. "todolist-api/internal/models/dto"
)

type UserUseCaseInterface interface {
	Register(ctx context.Context, request *RegisterUserRequestBody) (*UserResponse, error)

	Login(ctx context.Context, request *LoginUserRequestBody) (*TokenUserLoginResponse, error)

	VerifyAccessToken(ctx context.Context, accessToken *VerifyUserRequest) (*UserResponse, error)
}
