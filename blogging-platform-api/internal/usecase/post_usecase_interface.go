package usecase

import (
	"blogging-platform-api/internal/models/dto"
	"context"
)

type PostUseCaseInterface interface {
	CreateNewBlogPost(ctx context.Context, request *dto.CreatePostRequestBody) (*dto.PostResponse, error)

	GetSingleBlogPost(ctx context.Context, request *dto.GetPostRequestParam) (*dto.PostResponse, error)

	UpdateExistingBlogPost(ctx context.Context, request *dto.UpdatePostRequestBody) (*dto.PostResponse, error)

	DeleteExistingBlogPost(ctx context.Context, request *dto.DeletePostRequestParam) error

	GetAllBlogPost(ctx context.Context, request *dto.SearchPostRequestQueryParam) (*[]dto.PostResponse, int64, error)
}
