package repositories

import (
	"blogging-platform-api/internal/models/dto"
	"blogging-platform-api/internal/models/entities"
	"gorm.io/gorm"
)

type PostRepositoryInterface interface {
	Create(db *gorm.DB, entity *entities.Post) (*entities.Post, error)

	FindById(db *gorm.DB, id int64) (*entities.Post, error)

	Update(db *gorm.DB, entity *entities.Post) (*entities.Post, error)

	Remove(db *gorm.DB, id int64) error

	FindAll(db *gorm.DB, request *dto.SearchPostRequestQueryParam) (*[]entities.Post, int64, error)
}
