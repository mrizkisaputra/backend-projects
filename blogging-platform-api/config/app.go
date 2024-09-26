package config

import (
	"blogging-platform-api/internal/delivery/http"
	"blogging-platform-api/internal/delivery/http/middleware"
	"blogging-platform-api/internal/delivery/http/routes"
	"blogging-platform-api/internal/repositories"
	"blogging-platform-api/internal/usecase"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type Bootstrap Config
type Config struct {
	App      *fiber.App
	Log      *logrus.Logger
	DB       *gorm.DB
	Validate *validator.Validate
}

// menyiapkan semua kompoonen-komponen aplikasi
func NewBootstrap(config Bootstrap) {
	// setup repositories
	postRepository := repositories.NewPosRepository()

	// setup usecase
	postUsecase := usecase.NewPostUseCase(config.Log, config.Validate, postRepository, config.DB)

	// setup controller
	postController := http.NewPostController(postUsecase, config.Log)

	//setup middleware
	loggerMiddleware := middleware.NewLoggerMiddleware(config.Log)

	// setup routes
	route := routes.Route{
		App:              config.App,
		PostController:   postController,
		LoggerMiddleware: loggerMiddleware,
	}
	route.SetupRoutes()
}
