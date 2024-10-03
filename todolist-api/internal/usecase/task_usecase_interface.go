package usecase

import (
	"context"
	"todolist-api/internal/models/dto"
)

type TaskUseCaseInterface interface {
	Create(ctx context.Context, request *dto.CreateTaskRequest) (*dto.TaskResponse, error)

	Update(ctx context.Context, request *dto.UpdateTaskRequest) (*dto.TaskResponse, error)

	UpdateStatus(ctx context.Context, request *dto.UpdateStatusTaskRequest) (*dto.TaskResponse, error)

	Delete(ctx context.Context, request *dto.DeleteTaskRequest) error

	FindAll(ctx context.Context, request *dto.SearchTaskRequest) ([]dto.TaskResponse, int64, error)
}
