package repositories

import (
	"gorm.io/gorm"
)

type RepositoryImpl[T any] struct {
}

func (r *RepositoryImpl[T]) Create(db *gorm.DB, entity *T) (*T, error) {
	if err := db.Create(entity).Error; err != nil {
		return nil, err
	}
	return entity, nil
}

func (r *RepositoryImpl[T]) Update(db *gorm.DB, entity *T) (*T, error) {
	if err := db.Save(entity).Error; err != nil {
		return nil, err
	}
	return entity, nil
}

func (r *RepositoryImpl[T]) Delete(db *gorm.DB, entity *T, id string) error {
	return db.Delete(entity).Error
}

func (r *RepositoryImpl[T]) FindById(db *gorm.DB, entity *T, id any) error {
	return db.Where("id = ?", id).Take(entity).Error
}
