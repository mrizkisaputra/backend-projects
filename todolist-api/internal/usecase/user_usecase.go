package usecase

import (
	"context"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"strings"
	. "todolist-api/internal/models/dto"
	"todolist-api/internal/models/entities"
	. "todolist-api/internal/repositories"
)

type userUseCase struct {
	Log            *logrus.Logger
	UserRepository UserRepository
	DB             *gorm.DB
	Validate       *validator.Validate
}

func NewUserUseCase(
	log *logrus.Logger,
	db *gorm.DB,
	repository UserRepository,
	validate *validator.Validate,
) UserUseCaseInterface {
	return &userUseCase{
		Log:            log,
		DB:             db,
		UserRepository: repository,
		Validate:       validate,
	}
}

func (u *userUseCase) Register(ctx context.Context, request *RegisterUserRequestBody) (*UserResponse, error) {
	tx := u.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	if err := u.Validate.Struct(request); err != nil {
		u.Log.WithError(err).Error("validated user register failed")
		return nil, err.(validator.ValidationErrors)
	}

	total, err := u.UserRepository.FindAlreadyExistByEmail(tx, request.Email)
	if err != nil {
		u.Log.WithError(err).Error("find user by email failed")
		return nil, fiber.ErrInternalServerError
	}

	if total > 0 {
		return nil, fiber.NewError(fiber.StatusConflict, "email already registered")
	}

	encryptedPassword, errEncrypted := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
	if errEncrypted != nil {
		u.Log.WithError(errEncrypted).Error("failed to encrypt password")
		return nil, fiber.ErrInternalServerError
	}
	entity := entities.User{
		Id:       uuid.New().String(),
		Name:     request.Name,
		Email:    request.Email,
		Password: string(encryptedPassword),
	}
	user, errCreate := u.UserRepository.Create(tx, &entity)
	if errCreate != nil {
		u.Log.WithError(errCreate).Error("failed to create registered user")
		return nil, fiber.ErrInternalServerError
	}

	if err := tx.Commit().Error; err != nil {
		return nil, fiber.ErrInternalServerError
	}

	return u.toUserResponse(user), nil
}

func (u *userUseCase) Login(ctx context.Context, request *LoginUserRequestBody) (*TokenUserLoginResponse, error) {
	tx := u.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	if err := u.Validate.Struct(request); err != nil {
		u.Log.WithError(err).Error("validated user login failed")
		return nil, err.(validator.ValidationErrors)
	}

	user, errFindByEmail := u.UserRepository.FindByEmail(tx, request)
	if errFindByEmail != nil {
		u.Log.WithError(errFindByEmail).Error("find user by email failed")
		return nil, fiber.NewError(fiber.StatusUnauthorized, "email doesn't register")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.Password)); err != nil {
		u.Log.WithError(err).Error("password compare failed")
		return nil, fiber.NewError(fiber.StatusUnauthorized, "invalid password")
	}

	user.Token = uuid.New().String()
	user, err := u.UserRepository.Update(tx, user)
	if err != nil {
		return nil, fiber.ErrInternalServerError
	}

	if err := tx.Commit().Error; err != nil {
		u.Log.WithError(err).Error("error commit user login")
		return nil, fiber.ErrInternalServerError
	}

	return u.toTokenUserLoginResponse(user), nil
}

func (u *userUseCase) toUserResponse(user *entities.User) *UserResponse {
	return &UserResponse{
		Id:        user.Id,
		Name:      user.Name,
		Email:     user.Email,
		Password:  user.Password,
		CreatedAt: user.CreatedAt,
	}
}

func (u *userUseCase) toTokenUserLoginResponse(user *entities.User) *TokenUserLoginResponse {
	return &TokenUserLoginResponse{
		Token: user.Token,
	}
}

func (u *userUseCase) VerifyAccessToken(ctx context.Context, request *VerifyUserRequest) (*UserResponse, error) {
	tx := u.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	if err := u.Validate.Struct(request); err != nil {
		u.Log.WithError(err).Error("error validated access token")
		return nil, err.(validator.ValidationErrors)
	}

	accessToken, _ := strings.CutPrefix(request.AccessToken, "Bearer ")
	user, err := u.UserRepository.FindByToken(tx, accessToken)
	if err != nil {
		u.Log.WithError(err).Error("find user by token failed")
		return nil, fiber.NewError(fiber.StatusUnauthorized, "Access token invalid")
	}

	if err := tx.Commit().Error; err != nil {
		u.Log.WithError(err).Error("error commit verify access token")
		return nil, fiber.ErrInternalServerError
	}
	return u.toUserResponse(user), nil
}
