package usecase

import (
	"context"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"todolist-api/internal/models/dto"
	"todolist-api/internal/models/entities"
	"todolist-api/internal/repositories"
)

type taskUseCase struct {
	DB             *gorm.DB
	Log            *logrus.Logger
	Validate       *validator.Validate
	TaskRepository repositories.TaskRepository
}

func NewTaskUseCase(
	log *logrus.Logger,
	db *gorm.DB,
	repository repositories.TaskRepository,
	validate *validator.Validate,
) TaskUseCaseInterface {
	return &taskUseCase{
		DB:             db,
		Log:            log,
		Validate:       validate,
		TaskRepository: repository,
	}
}

func (t *taskUseCase) Create(ctx context.Context, request *dto.CreateTaskRequest) (*dto.TaskResponse, error) {
	tx := t.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	if err := t.Validate.Struct(request); err != nil {
		t.Log.WithError(err).Error("error validate task request body")
		return nil, err.(validator.ValidationErrors)
	}

	entity := entities.Task{
		Id:          uuid.New().String(),
		IdUser:      request.IdUser,
		Title:       request.Title,
		Description: request.Description,
		Status:      "in-progress",
	}
	task, err := t.TaskRepository.Create(tx, &entity)
	if err != nil {
		t.Log.WithError(err).Error("error create task")
		return nil, fiber.ErrInternalServerError
	}

	if err := tx.Commit().Error; err != nil {
		t.Log.WithError(err).Error("error commit task")
		return nil, fiber.ErrInternalServerError
	}

	return t.toTaskResponse(task), nil
}

func (t *taskUseCase) Update(ctx context.Context, request *dto.UpdateTaskRequest) (*dto.TaskResponse, error) {
	tx := t.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	if err := t.Validate.Struct(request); err != nil {
		t.Log.WithError(err).Error("error validated request")
		return nil, err.(validator.ValidationErrors)
	}

	entity := new(entities.Task)
	if err := t.TaskRepository.FindByIdAndIdUser(tx, entity, request.Id, request.IdUser); err != nil {
		t.Log.WithError(err).Error("error find task by id")
		return nil, fiber.ErrNotFound
	}
	entity.Title = request.Title
	entity.Description = request.Description
	task, err := t.TaskRepository.Update(tx, entity)
	if err != nil {
		t.Log.WithError(err).Error("error update task")
		return nil, fiber.ErrInternalServerError
	}

	if err := tx.Commit().Error; err != nil {
		t.Log.WithError(err).Error("error commit update task")
		return nil, fiber.ErrInternalServerError
	}

	return t.toTaskResponse(task), nil
}

func (t *taskUseCase) UpdateStatus(ctx context.Context, request *dto.UpdateStatusTaskRequest) (*dto.TaskResponse, error) {
	tx := t.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	if err := t.Validate.Struct(request); err != nil {
		t.Log.WithError(err).Error("error validated request")
		return nil, err.(validator.ValidationErrors)
	}

	entity := new(entities.Task)
	if err := t.TaskRepository.FindByIdAndIdUser(tx, entity, request.Id, request.IdUser); err != nil {
		t.Log.WithError(err).Error("error find task by id")
		return nil, fiber.ErrNotFound
	}
	entity.Status = request.Mark
	if err := t.TaskRepository.UpdateStatus(tx, entity); err != nil {
		t.Log.WithError(err).Error("error update status task")
		return nil, fiber.ErrInternalServerError
	}

	if err := tx.Commit().Error; err != nil {
		t.Log.WithError(err).Error("error commit update task")
		return nil, fiber.ErrInternalServerError
	}

	return t.toTaskResponse(entity), nil
}

func (t *taskUseCase) Delete(ctx context.Context, request *dto.DeleteTaskRequest) error {
	tx := t.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	if err := t.Validate.Struct(request); err != nil {
		return err.(validator.ValidationErrors)
	}

	entity := new(entities.Task)
	if err := t.TaskRepository.FindByIdAndIdUser(tx, entity, request.Id, request.IdUser); err != nil {
		t.Log.WithError(err).Error("error find task by id")
		return fiber.ErrNotFound
	}

	if err := t.TaskRepository.Delete(tx, entity, entity.Id); err != nil {
		t.Log.WithError(err).Error("error delete task")
		return fiber.ErrInternalServerError
	}

	if err := tx.Commit().Error; err != nil {
		t.Log.WithError(err).Error("error commit delete task")
		return fiber.ErrInternalServerError
	}
	return nil
}

func (t *taskUseCase) toTaskResponse(task *entities.Task) *dto.TaskResponse {
	return &dto.TaskResponse{
		Id:          task.Id,
		IdUser:      task.IdUser,
		Title:       task.Title,
		Description: task.Description,
		Status:      task.Status,
		CreatedAt:   task.CreatedAt,
		UpdatedAt:   task.UpdatedAt,
	}
}

func (t *taskUseCase) FindAll(ctx context.Context, request *dto.SearchTaskRequest) ([]dto.TaskResponse, int64, error) {
	tx := t.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	if err := t.Validate.Struct(request); err != nil {
		t.Log.WithError(err).Error("error validated request")
		return nil, 0, err.(validator.ValidationErrors)
	}

	tasks, total, err := t.TaskRepository.FindAll(tx, request)
	if err != nil {
		t.Log.WithError(err).Error("error find all tasks")
		return nil, 0, fiber.ErrInternalServerError
	}

	if err := tx.Commit().Error; err != nil {
		t.Log.WithError(err).Error("error commit delete task")
		return nil, 0, fiber.ErrInternalServerError
	}

	response := make([]dto.TaskResponse, len(*tasks))
	for i, task := range *tasks {
		response[i] = *t.toTaskResponse(&task)
	}

	return response, total, nil
}
