package http

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
	"math"
	"todolist-api/internal/delivery/http/middleware"
	"todolist-api/internal/models/dto"
	. "todolist-api/internal/usecase"
)

type taskController struct {
	UseCase TaskUseCaseInterface
	Log     *logrus.Logger
}

func NewTaskController(
	usecase TaskUseCaseInterface,
	log *logrus.Logger,
) TaskControllerInterface {
	return &taskController{
		UseCase: usecase,
		Log:     log,
	}
}

func (t taskController) Create(ctx *fiber.Ctx) error {
	user := middleware.GetUser(ctx)

	if contentType := ctx.Get("Content-Type"); contentType != "application/json" {
		t.Log.Error("Content-Type not supported")
		return fiber.NewError(fiber.StatusBadRequest, "Content-Type must be application/json")
	}

	requestBody := new(dto.CreateTaskRequest)
	if err := ctx.BodyParser(requestBody); err != nil {
		t.Log.WithError(err).Error("error parsing request body")
		return fiber.ErrInternalServerError
	}
	requestBody.IdUser = user.Id

	responseEntity, err := t.UseCase.Create(ctx.UserContext(), requestBody)
	if err != nil {
		return err
	}

	apiResponse := dto.ApiResponse{
		Status: fiber.StatusCreated,
		Data:   responseEntity,
		Paging: nil,
	}
	return ctx.Status(fiber.StatusCreated).JSON(&apiResponse, "application/json")
}

func (t taskController) Update(ctx *fiber.Ctx) error {
	user := middleware.GetUser(ctx)

	if contentType := ctx.Get("Content-Type"); contentType != "application/json" {
		t.Log.Error("Content-Type not supported")
		return fiber.NewError(fiber.StatusBadRequest, "Content-Type must be application/json")
	}

	// Jika query param `mark=done` ada, hanya update status ke 'done'
	requestParam := dto.UpdateStatusTaskRequest{
		IdUser: user.Id,
		Id:     ctx.Params("id", ""),
		Mark:   ctx.Query("mark", ""),
	}
	if requestParam.Mark != "" {
		responseEntity, err := t.UseCase.UpdateStatus(ctx.UserContext(), &requestParam)
		if err != nil {
			return err
		}
		apiResponse := dto.ApiResponse{
			Status: fiber.StatusOK,
			Data:   responseEntity,
			Paging: nil,
		}
		return ctx.Status(fiber.StatusOK).JSON(&apiResponse, "application/json")
	}

	// Jika tidak ada query param `mark=done`, proses body untuk update data lain
	requestBody := new(dto.UpdateTaskRequest)
	if err := ctx.BodyParser(requestBody); err != nil {
		t.Log.WithError(err).Error("error parsing request body")
		return fiber.ErrInternalServerError
	}
	requestBody.Id = requestParam.Id
	requestBody.IdUser = user.Id
	responseEntity, err := t.UseCase.Update(ctx.UserContext(), requestBody)
	if err != nil {
		return err
	}

	apiResponse := dto.ApiResponse{
		Status: fiber.StatusOK,
		Data:   responseEntity,
		Paging: nil,
	}
	return ctx.Status(fiber.StatusOK).JSON(&apiResponse, "application/json")
}

func (t taskController) Delete(ctx *fiber.Ctx) error {
	user := middleware.GetUser(ctx)

	request := dto.DeleteTaskRequest{
		Id:     ctx.Params("id"),
		IdUser: user.Id,
	}
	if err := t.UseCase.Delete(ctx.UserContext(), &request); err != nil {
		return err
	}

	return ctx.SendStatus(fiber.StatusNoContent)
}

func (t taskController) FindAll(ctx *fiber.Ctx) error {
	user := middleware.GetUser(ctx)

	if contentType := ctx.Get("Content-Type"); contentType != "application/json" {
		t.Log.Error("Content-Type not supported")
		return fiber.NewError(fiber.StatusBadRequest, "Content-Type must be application/json")
	}

	request := dto.SearchTaskRequest{
		IdUser: user.Id,
		Status: ctx.Query("status", "in-progress"),
		Sort:   ctx.Query("sort", "created_at"),
		Order:  ctx.Query("order", "asc"),
		Page:   ctx.QueryInt("page", 1),
		Limit:  ctx.QueryInt("limit", 20),
	}
	responseEntity, total, err := t.UseCase.FindAll(ctx.UserContext(), &request)
	if err != nil {
		return err
	}

	paging := dto.PageMetadata{
		TotalData: total,
		Page:      request.Page,
		Limit:     request.Limit,
		TotalPage: int64(math.Ceil(float64(total) / float64(request.Limit))),
	}

	apiResponse := dto.ApiResponse{
		Status: fiber.StatusOK,
		Data:   responseEntity,
		Paging: &paging,
	}
	return ctx.Status(fiber.StatusOK).JSON(&apiResponse, "application/json")
}
