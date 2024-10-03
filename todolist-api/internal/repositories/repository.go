package repositories

import "gorm.io/gorm"

type Repository[T any] interface {
	Create(db *gorm.DB, entity *T) (*T, error)

	Update(db *gorm.DB, entity *T) (*T, error)

	Delete(db *gorm.DB, entity *T, id string) error

	FindById(db *gorm.DB, entity *T, id any) error
}
