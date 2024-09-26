package routes

import (
	"blogging-platform-api/internal/delivery/http"
	"github.com/gofiber/fiber/v2"
)

type Route struct {
	App              *fiber.App
	PostController   http.PostControllerInterface
	LoggerMiddleware fiber.Handler
}

func (r *Route) SetupRoutes() {
	r.PostRoute()
}

func (r *Route) PostRoute() {
	r.App.Use(r.LoggerMiddleware)
	r.App.Post("/api/posts", r.PostController.Create)
	r.App.Get("/api/posts/:id", r.PostController.Get)
	r.App.Delete("/api/posts/:id", r.PostController.Delete)
	r.App.Put("/api/posts/:id", r.PostController.Update)
	r.App.Get("/api/posts", r.PostController.List)
}
