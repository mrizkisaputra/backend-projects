package usecase

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
	"personal-blog/internal/models/dto"
	. "personal-blog/internal/models/entities"
	. "personal-blog/internal/repositories"
	"strings"
	"time"
)

type articleUseCase struct {
	Dir        string
	Log        *logrus.Logger
	Repository ArticleRepositoryInterface
	Validate   *validator.Validate
}

func NewArticleUseCase(
	dir string,
	log *logrus.Logger,
	articleRepository ArticleRepositoryInterface,
	validate *validator.Validate,
) ArticleUseCaseInterface {
	return &articleUseCase{
		Dir:        dir,
		Log:        log,
		Repository: articleRepository,
		Validate:   validate,
	}
}

func (self *articleUseCase) Create(request *dto.ArticlePayloadBodyRequest) error {
	if err := self.Validate.Struct(request); err != nil {
		self.Log.WithError(err).Error("error validation")
		return fiber.ErrBadRequest
	}
	publish, err := time.Parse("2006-01-02", request.PublishDate)
	if err != nil {
		self.Log.WithError(err).Error("error parse time")
		return fiber.ErrInternalServerError
	}
	article := &Article{
		Id:          time.Now().UnixNano(),
		Title:       request.Title,
		Content:     request.Content,
		PublishDate: publish,
		Category:    request.Category,
		Tags:        strings.Split(request.Tags, ","),
		UpdatedAt:   time.Now().Local(),
	}
	if err := self.Repository.Create(article, self.Dir); err != nil {
		self.Log.WithError(err).Error("error create article")
		return fiber.ErrInternalServerError
	}
	return nil
}

func (self *articleUseCase) FindById(request *dto.ArticlePayloadParamIdRequest) (*Article, error) {
	if err := self.Validate.Struct(request); err != nil {
		self.Log.WithError(err).Error("error validation")
		return nil, fiber.ErrBadRequest
	}
	article, err := self.Repository.FindById(request.Id, self.Dir)
	if err != nil {
		self.Log.WithError(err).Errorf("error find article with id %d", request.Id)
		return nil, err
	}
	return article, nil
}

func (self *articleUseCase) FindAll() ([]Article, error) {
	articles, err := self.Repository.FindAll(self.Dir)
	if err != nil {
		self.Log.WithError(err).Error("error find all articles")
		return nil, err
	}
	return articles, nil
}

func (self *articleUseCase) Delete(request *dto.ArticlePayloadParamIdRequest) error {
	if err := self.Validate.Struct(request); err != nil {
		self.Log.WithError(err).Error("error validation")
		return fiber.ErrBadRequest
	}
	_, err := self.Repository.Delete(request.Id, self.Dir)
	if err != nil {
		self.Log.WithError(err).Error("error delete article")
		return err
	}
	return nil
}

func (self *articleUseCase) Update(requestId *dto.ArticlePayloadParamIdRequest, request *dto.ArticlePayloadBodyRequest) error {
	if err := self.Validate.Struct(request); err != nil {
		self.Log.WithError(err).Error("error validation ArticlePayloadBodyRequest")
		return fiber.ErrBadRequest
	}

	if err := self.Validate.Struct(requestId); err != nil {
		self.Log.WithError(err).Error("error validation ArticlePayloadParamIdRequest")
		return fiber.ErrBadRequest
	}

	article := &Article{
		Id:          requestId.Id,
		Title:       request.Title,
		Content:     request.Content,
		PublishDate: time.Now().Local(),
		Category:    request.Category,
		Tags:        strings.Split(request.Tags, ","),
		UpdatedAt:   time.Now().Local(),
	}
	_, err := self.Repository.Update(requestId.Id, article, self.Dir)
	if err != nil {
		return err
	}
	return nil
}
