package routes

import (
	"github.com/gofiber/fiber/v2"
	"personal-blog/internal/delivery/http"
)

type Route struct {
	App               *fiber.App
	ArticleController http.ArticleControllerInterface
	AuthMiddleware    fiber.Handler
}

func (self *Route) Setup() {
	self.SetupStaticRoute()
	self.SetupGuestRoute()
	self.SetupAuthRoute()
}

func (self *Route) SetupStaticRoute() {
	self.App.Static("/static/css/output.css", "./web/src/css/output.css")
	self.App.Static("/static/js/jquery-3.7.1.min.js", "./web/src/js/jquery-3.7.1.min.js")
	self.App.Static("/static/js/new-article.js", "./web/src/js/new-article.js")
	self.App.Static("/static/js/index.js", "./web/src/js/index.js")
}

func (self *Route) SetupGuestRoute() {
	self.App.Get("/", self.ArticleController.Home)
	self.App.Get("/article/:id", self.ArticleController.FindById)
	self.App.Get("/login", self.ArticleController.PageLogin)
	self.App.Post("/login", self.ArticleController.Login)
}

func (self *Route) SetupAuthRoute() {
	self.App.Use(self.AuthMiddleware)

	self.App.Get("/admin", self.ArticleController.PageAdminHome)
	self.App.Get("/new", self.ArticleController.PageNewArticle)
	self.App.Get("/edit/:id", self.ArticleController.PageEditArticle)
	self.App.Post("/new", self.ArticleController.Create)
	self.App.Post("/edit/:id", self.ArticleController.Update)
	self.App.Get("/article", self.ArticleController.Delete)
	self.App.Get("/logout", self.ArticleController.Logout)
}
