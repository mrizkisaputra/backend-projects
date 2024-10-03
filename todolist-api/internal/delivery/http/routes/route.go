package routes

import (
	"github.com/gofiber/fiber/v2"
	"todolist-api/internal/delivery/http"
)

type RouteConfig Route
type Route struct {
	App            *fiber.App
	UserController http.UserControllerInterface
	TaskController http.TaskControllerInterface
	AuthMiddleware fiber.Handler
}

func NewRoute(config *RouteConfig) *Route {
	return &Route{
		App:            config.App,
		UserController: config.UserController,
		TaskController: config.TaskController,
		AuthMiddleware: config.AuthMiddleware,
	}
}

func (r *Route) Setup() {
	r.guestRoutes()
	r.authRoutes()
}

func (r *Route) guestRoutes() {
	r.App.Post("/api/register", r.UserController.Register)
	r.App.Post("/api/login", r.UserController.Login)
}

func (r *Route) authRoutes() {
	r.App.Use(r.AuthMiddleware)
	r.App.Post("/api/todos", r.TaskController.Create)
	r.App.Post("/api/todos/:id", r.TaskController.Update)
	r.App.Delete("/api/todos/:id", r.TaskController.Delete)
	r.App.Get("/api/todos", r.TaskController.FindAll)
}
