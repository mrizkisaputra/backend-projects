package usecase

import (
	"personal-blog/internal/models/dto"
	. "personal-blog/internal/models/entities"
)

type ArticleUseCaseInterface interface {
	Create(request *dto.ArticlePayloadBodyRequest) error
	FindById(request *dto.ArticlePayloadParamIdRequest) (*Article, error)
	FindAll() ([]Article, error)
	Delete(request *dto.ArticlePayloadParamIdRequest) error
	Update(requestId *dto.ArticlePayloadParamIdRequest, request *dto.ArticlePayloadBodyRequest) error
}
