package http

import (
	"blogging-platform-api/internal/models/dto"
	. "blogging-platform-api/internal/usecase"
	"context"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
	"math"
)

type postController struct {
	UseCase PostUseCaseInterface
	Log     *logrus.Logger
}

func NewPostController(usecase PostUseCaseInterface, logger *logrus.Logger) PostControllerInterface {
	return &postController{
		UseCase: usecase,
		Log:     logger,
	}
}

func (p *postController) Create(ctx *fiber.Ctx) error {
	contentType := ctx.Get("Content-Type")
	if contentType != "application/json" {
		p.Log.WithFields(logrus.Fields{"layer": "delivery/http"}).Error("Content-Type is not application/json")
		return fiber.NewError(fiber.StatusBadRequest, "REQUIRED_CONTENT_TYPE_APPLICATION/JSON")
	}

	var requestBody = new(dto.CreatePostRequestBody)
	if err := ctx.BodyParser(requestBody); err != nil {
		p.Log.WithError(err).WithFields(logrus.Fields{"layer": "delivery/http"}).Error("Error parsing request body")
		return fiber.NewError(fiber.StatusBadRequest, "REQUIRED_REQUEST_BODY")
	}

	response, err := p.UseCase.CreateNewBlogPost(context.Background(), requestBody)
	if err != nil {
		return err
	}

	apiResponse := &dto.APIResponse[dto.PostResponse]{
		Status: fiber.StatusCreated,
		Data:   *response,
		Paging: nil,
	}
	return ctx.Status(fiber.StatusCreated).JSON(apiResponse, "application/json")
}

func (p *postController) Get(ctx *fiber.Ctx) error {
	if contentType := ctx.Get("Content-Type"); contentType != "application/json" {
		p.Log.WithFields(logrus.Fields{"layer": "delivery/http"}).Error("Content-Type is not application/json")
		return fiber.NewError(fiber.StatusBadRequest, "REQUIRED_CONTENT_TYPE_APPLICATION/JSON")
	}

	id, err1 := ctx.ParamsInt("id", 0)
	if err1 != nil {
		return fiber.ErrInternalServerError
	}

	requestParam := dto.GetPostRequestParam{
		Id: int64(id),
	}
	postResponse, err := p.UseCase.GetSingleBlogPost(context.Background(), &requestParam)
	if err != nil {
		return err
	}
	apiResponse := dto.APIResponse[dto.PostResponse]{
		Status: fiber.StatusOK,
		Data:   *postResponse,
		Paging: nil,
	}
	return ctx.Status(fiber.StatusOK).JSON(apiResponse, "application/json")
}

func (p *postController) Delete(ctx *fiber.Ctx) error {
	id, err1 := ctx.ParamsInt("id", 0)
	if err1 != nil {
		return fiber.ErrInternalServerError
	}

	requestParam := dto.DeletePostRequestParam{
		Id: int64(id),
	}

	if err := p.UseCase.DeleteExistingBlogPost(context.Background(), &requestParam); err != nil {
		return err
	}

	return ctx.SendStatus(fiber.StatusNoContent)
}

func (p *postController) Update(ctx *fiber.Ctx) error {
	if contentType := ctx.Get("Content-Type"); contentType != "application/json" {
		p.Log.WithFields(logrus.Fields{"layer": "delivery/http"}).Error("Content-Type is not application/json")
		return fiber.NewError(fiber.StatusBadRequest, "REQUIRED_CONTENT_TYPE_APPLICATION/JSON")
	}

	requestBody := new(dto.UpdatePostRequestBody)
	if err := ctx.BodyParser(requestBody); err != nil {
		p.Log.WithError(err).WithFields(logrus.Fields{"layer": "delivery/http"}).Error("error parsing request body")
		return fiber.NewError(fiber.StatusBadRequest, "REQUIRED_REQUEST_BODY")
	}
	id, errGetParam := ctx.ParamsInt("id", 0)
	if errGetParam != nil {
		return fiber.ErrInternalServerError
	}
	requestBody.Id = int64(id)

	postResponse, err := p.UseCase.UpdateExistingBlogPost(context.Background(), requestBody)
	if err != nil {
		return err
	}

	apiResponse := dto.APIResponse[dto.PostResponse]{
		Status: fiber.StatusOK,
		Data:   *postResponse,
		Paging: nil,
	}
	return ctx.Status(fiber.StatusOK).JSON(apiResponse, "application/json")
}

func (p *postController) List(ctx *fiber.Ctx) error {
	if contentType := ctx.Get("Content-Type"); contentType != "application/json" {
		return fiber.NewError(fiber.StatusBadRequest, "REQUIRED_CONTENT_TYPE_APPLICATION/JSON")
	}

	requestQueryParam := dto.SearchPostRequestQueryParam{
		Title:    ctx.Query("title", ""),
		Category: ctx.Query("category", ""),
		Page:     ctx.QueryInt("page", 1),
		Limit:    ctx.QueryInt("limit", 10),
	}

	postsResponse, total, err := p.UseCase.GetAllBlogPost(context.Background(), &requestQueryParam)
	if err != nil {
		return err
	}

	paging := dto.PageMetadata{
		Page:      requestQueryParam.Page,
		Limit:     requestQueryParam.Limit,
		TotalData: total,
		TotalPage: int64(math.Ceil(float64(total) / float64(requestQueryParam.Limit))),
	}

	apiResponse := dto.APIResponse[[]dto.PostResponse]{
		Status: fiber.StatusOK,
		Data:   *postsResponse,
		Paging: &paging,
	}
	return ctx.Status(fiber.StatusOK).JSON(apiResponse, "application/json")
}
