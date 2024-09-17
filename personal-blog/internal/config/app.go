package config

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
	"github.com/sirupsen/logrus"
	"personal-blog/internal/delivery/http"
	"personal-blog/internal/delivery/http/middleware"
	"personal-blog/internal/delivery/http/routes"
	"personal-blog/internal/repositories"
	"personal-blog/internal/usecase"
)

type BootstrapConfig struct {
	App          *fiber.App
	Log          *logrus.Logger
	Validate     *validator.Validate
	DirFile      string
	StoreSession *session.Store
}

func NewBootstrap(config *BootstrapConfig) {
	//setup repositories
	articleRepository := repositories.NewArticleRepository(config.Log)

	// setup usecase
	articleUseCase := usecase.NewArticleUseCase(config.DirFile, config.Log, articleRepository, config.Validate)

	// setup controllers
	articleController := http.NewArticleController(articleUseCase, config.Log, config.StoreSession)

	// setup middleware
	authMiddleware := middleware.NewAuthMiddleware(config.StoreSession)

	// setup routes
	route := routes.Route{
		App:               config.App,
		ArticleController: articleController,
		AuthMiddleware:    authMiddleware,
	}
	route.Setup()
}
