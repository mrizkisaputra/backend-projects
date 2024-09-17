package usecase

import (
	"github.com/go-playground/validator/v10"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"personal-blog/internal/models/dto"
	"personal-blog/internal/models/entities"
	"testing"
	"time"
)

var repository = MockArticleRepository{Mock: mock.Mock{}}
var usecase = NewArticleUseCase(
	"../../articles/test",
	logrus.New(),
	&repository,
	validator.New(),
)

func TestArticleUseCase_Create_Successfully(t *testing.T) {
	repository.Mock.On("Create", mock.Anything, "../../articles/test").Return(nil)

	publish := time.Now().Local().Format("2006-01-02")
	err := usecase.Create(&dto.ArticlePayloadBodyRequest{
		Title:       "Tutorial unit testing",
		Content:     "this is content it",
		Tags:        "unittesting,programming",
		Category:    "tutorial",
		PublishDate: publish,
	})
	assert.Nil(t, err)
}

func TestArticleUseCase_Create_ValidationError(t *testing.T) {
	//repository.Mock.On("Create", mock.Anything, "../../articles/test").Return(nil)

	publish := time.Now().Local().Format("2006-01-02")
	err := usecase.Create(&dto.ArticlePayloadBodyRequest{
		Title:       "", //title must be required
		Content:     "this is content it",
		Tags:        "unittesting,programming",
		Category:    "123", //category mus be alpha
		PublishDate: publish,
	})
	assert.Error(t, err)
}

func TestArticleUseCase_FindById_Successfully(t *testing.T) {
	repository.Mock.On("FindById", int64(1), "../../articles/test").Return(entities.Article{
		Id:      1,
		Title:   "Example",
		Content: "Content example",
		Tags:    []string{"example"},
	}, nil)

	article, err := usecase.FindById(&dto.ArticlePayloadParamIdRequest{
		Id: 1,
	})
	assert.Nil(t, err)
	assert.NotNil(t, article)
}
