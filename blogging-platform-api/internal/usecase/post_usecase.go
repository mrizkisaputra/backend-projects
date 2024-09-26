package usecase

import (
	"blogging-platform-api/internal/models/converters"
	"blogging-platform-api/internal/models/dto"
	"blogging-platform-api/internal/models/entities"
	"blogging-platform-api/internal/repositories"
	"context"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"strings"
	"time"
)

type postUseCase struct {
	Log        *logrus.Logger
	Validate   *validator.Validate
	Repository repositories.PostRepositoryInterface
	DB         *gorm.DB
}

func NewPostUseCase(
	logger *logrus.Logger,
	validation *validator.Validate,
	postRepository repositories.PostRepositoryInterface,
	database *gorm.DB,
) PostUseCaseInterface {
	return &postUseCase{
		Log:        logger,
		Validate:   validation,
		Repository: postRepository,
		DB:         database,
	}
}

func (p *postUseCase) CreateNewBlogPost(ctx context.Context, request *dto.CreatePostRequestBody) (*dto.PostResponse, error) {
	tx := p.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	if err := p.Validate.Struct(request); err != nil {
		p.Log.WithError(err).WithFields(logrus.Fields{"layer": "usecase"}).Error("error validated request body")
		errValidated := err.(validator.ValidationErrors)
		return nil, errValidated
	}

	entity := &entities.Post{
		Id:       time.Now().UnixNano(),
		Title:    request.Title,
		Content:  request.Content,
		Category: request.Category,
		Tags:     strings.Join(request.Tags, ","),
	}
	post, err := p.Repository.Create(tx, entity)
	if err != nil {
		p.Log.WithError(err).WithFields(logrus.Fields{"layer": "usecase"}).Error("error created blog posts")
		return nil, fiber.ErrInternalServerError
	}

	if err := tx.Commit().Error; err != nil {
		p.Log.WithError(err).WithFields(logrus.Fields{"layer": "usecase"}).Error("error commit transaction to created blog posts")
		return nil, fiber.ErrInternalServerError
	}

	return converters.PostToPostResponse(post), nil
}

func (p *postUseCase) GetSingleBlogPost(ctx context.Context, request *dto.GetPostRequestParam) (*dto.PostResponse, error) {
	tx := p.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	if err := p.Validate.Struct(request); err != nil {
		p.Log.WithError(err).WithFields(logrus.Fields{"layer": "usecase"}).Error("error validated request path param")
		errValidated := err.(validator.ValidationErrors)
		return nil, errValidated
	}

	post, err := p.Repository.FindById(tx, request.Id)
	if err != nil {
		p.Log.WithError(err).WithFields(logrus.Fields{"layer": "usecase"}).Error("error getting blog post")
		return nil, fiber.ErrNotFound
	}

	if err := tx.Commit().Error; err != nil {
		p.Log.WithError(err).WithFields(logrus.Fields{"layer": "usecase"}).Error("error commit transaction to find single blog posts")
		return nil, fiber.ErrInternalServerError
	}

	return converters.PostToPostResponse(post), nil
}

func (p *postUseCase) UpdateExistingBlogPost(ctx context.Context, request *dto.UpdatePostRequestBody) (*dto.PostResponse, error) {
	tx := p.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	if err := p.Validate.Struct(request); err != nil {
		p.Log.WithError(err).WithFields(logrus.Fields{"layer": "usecase"}).Error("error validated request body update blog post")
		errValidated := err.(validator.ValidationErrors)
		return nil, errValidated
	}

	post, errFindById := p.Repository.FindById(tx, request.Id)
	if errFindById != nil {
		p.Log.WithError(errFindById).WithFields(logrus.Fields{"layer": "usecase"}).Error("error getting blog post")
		return nil, fiber.ErrNotFound
	}

	post.Title = request.Title
	post.Content = request.Content
	post.Category = request.Category
	post.Tags = strings.Join(request.Tags, ",")

	result, err := p.Repository.Update(tx, post)
	if err != nil {
		p.Log.WithError(err).WithFields(logrus.Fields{"layer": "usecase"}).Error("error updated blog post")
		return nil, fiber.ErrInternalServerError
	}

	if err := tx.Commit().Error; err != nil {
		p.Log.WithError(err).WithFields(logrus.Fields{"layer": "usecase"}).Error("error commit transaction to find single blog posts")
		return nil, fiber.ErrInternalServerError
	}
	return converters.PostToPostResponse(result), nil
}

func (p *postUseCase) DeleteExistingBlogPost(ctx context.Context, request *dto.DeletePostRequestParam) error {
	tx := p.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	if err := p.Validate.Struct(request); err != nil {
		p.Log.WithError(err).WithFields(logrus.Fields{"layer": "usecase"}).Error("error validated request path param")
		errValidated := err.(validator.ValidationErrors)
		return errValidated
	}

	post, err := p.Repository.FindById(tx, request.Id)
	if err != nil {
		return fiber.ErrNotFound
	}

	if err := p.Repository.Remove(tx, post.Id); err != nil {
		return fiber.ErrInternalServerError
	}

	if err := tx.Commit().Error; err != nil {
		p.Log.WithError(err).WithFields(logrus.Fields{"layer": "usecase"}).Error("error commit transaction to delete blog posts")
		return fiber.ErrInternalServerError
	}

	return nil
}

func (p *postUseCase) GetAllBlogPost(ctx context.Context, request *dto.SearchPostRequestQueryParam) (*[]dto.PostResponse, int64, error) {
	tx := p.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	if err := p.Validate.Struct(request); err != nil {
		p.Log.WithError(err).WithFields(logrus.Fields{"layer": "usecase"}).Error("error validated query param")
		errValidated := err.(validator.ValidationErrors)
		return nil, 0, errValidated
	}

	posts, total, err := p.Repository.FindAll(tx, request)
	if err != nil {
		p.Log.WithError(err).WithFields(logrus.Fields{"layer": "usecase"}).Error("error getting blog posts")
		return nil, 0, fiber.ErrInternalServerError
	}

	if err := tx.Commit().Error; err != nil {
		p.Log.WithError(err).WithFields(logrus.Fields{"layer": "usecase"}).Error("error commit transaction to find blog posts")
		return nil, 0, fiber.ErrInternalServerError
	}

	responses := make([]dto.PostResponse, len(*posts))
	for i, post := range *posts {
		responses[i] = *converters.PostToPostResponse(&post)
	}

	return &responses, total, nil
}
