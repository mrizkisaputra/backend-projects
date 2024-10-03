package repositories

import (
	"gorm.io/gorm"
	"todolist-api/internal/models/dto"
	. "todolist-api/internal/models/entities"
)

type UserRepository struct {
	Repository[User]
}

func NewUserRepository() *UserRepository {
	return &UserRepository{
		Repository: &RepositoryImpl[User]{},
	}
}

func (u *UserRepository) FindAlreadyExistByEmail(db *gorm.DB, email string) (int64, error) {
	// SQL: select count(email) from users where email = ?
	var total int64
	if err := db.Model(User{}).Where(User{Email: email}).Count(&total).Error; err != nil {
		return 0, err
	}
	return total, nil
}

func (u *UserRepository) FindByEmail(db *gorm.DB, request *dto.LoginUserRequestBody) (*User, error) {
	var user User
	err := db.Where(User{Email: request.Email}).Take(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (u *UserRepository) FindByToken(db *gorm.DB, token string) (*User, error) {
	var user User
	if err := db.Model(User{}).Where(User{Token: token}).Take(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}
