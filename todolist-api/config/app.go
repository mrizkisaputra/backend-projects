package config

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"todolist-api/internal/delivery/http"
	"todolist-api/internal/delivery/http/middleware"
	"todolist-api/internal/delivery/http/routes"
	"todolist-api/internal/repositories"
	"todolist-api/internal/usecase"
)

type BootstrapConfig struct {
	App      *fiber.App
	DB       *gorm.DB
	Log      *logrus.Logger
	Validate *validator.Validate
}

// tempat menyiapkan semua kompoonen-komponen aplikasi
func NewBootstrap(config *BootstrapConfig) {
	// setup repository
	userRepository := repositories.NewUserRepository()
	taskRepository := repositories.NewTaskRepository()

	// setup usecase
	userUseCase := usecase.NewUserUseCase(config.Log, config.DB, *userRepository, config.Validate)
	taskUseCase := usecase.NewTaskUseCase(config.Log, config.DB, *taskRepository, config.Validate)

	// setup controller
	userController := http.NewUserController(userUseCase, config.Log)
	taskController := http.NewTaskController(taskUseCase, config.Log)

	// setup middleware
	authMiddleware := middleware.NewAuthMiddleware(userUseCase, config.Log)

	// setup route
	route := routes.NewRoute(&routes.RouteConfig{
		App:            config.App,
		UserController: userController,
		TaskController: taskController,
		AuthMiddleware: authMiddleware,
	})
	route.Setup()
}
