package repositories

import (
	"fmt"
	"gorm.io/gorm"
	"todolist-api/internal/models/dto"
	. "todolist-api/internal/models/entities"
)

type TaskRepository struct {
	Repository[Task]
}

func NewTaskRepository() *TaskRepository {
	return &TaskRepository{
		Repository: &RepositoryImpl[Task]{},
	}
}

func (t *TaskRepository) FindByIdAndIdUser(db *gorm.DB, entity *Task, id string, idUser string) error {
	return db.Where("id = ? AND id_user = ?", id, idUser).Take(entity).Error
}

func (t *TaskRepository) UpdateStatus(db *gorm.DB, entity *Task) error {
	/**
	SQL: update tasks set status = 'done' where id = ?
	*/
	return db.Select("status").Save(entity).Error
}

func (t *TaskRepository) FindAll(db *gorm.DB, request *dto.SearchTaskRequest) (*[]Task, int64, error) {
	var data []Task
	if err := db.Model(new(Task)).
		Scopes(t.sorting(request)).
		Scopes(t.filter(request)).
		Offset((request.Page - 1) * request.Limit).
		Limit(request.Limit).
		Find(&data).Error; err != nil {

		return nil, 0, err
	}

	var total int64
	if err := db.Model(new(Task)).Scopes(t.filter(request)).Scopes(t.sorting(request)).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	return &data, total, nil
}

func (t *TaskRepository) filter(request *dto.SearchTaskRequest) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB { // filtering
		db.Where("id_user = ?", request.IdUser)
		if request.Status != "" {
			db.Where("status = ?", request.Status)
		}
		return db
	}
}

func (t *TaskRepository) sorting(request *dto.SearchTaskRequest) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB { // sorting
		if request.Sort != "" {
			orderByCreatedAt := fmt.Sprintf("%s %s", request.Sort, request.Order)
			db.Order(orderByCreatedAt)
		}
		return db
	}
}
