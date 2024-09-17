package http

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
	"github.com/sirupsen/logrus"
	"personal-blog/internal/models/dto"
	. "personal-blog/internal/usecase"
	"strings"
)

type articleController struct {
	Log          *logrus.Logger
	UseCase      ArticleUseCaseInterface
	StoreSession *session.Store
}

func NewArticleController(
	articleUseCase ArticleUseCaseInterface,
	log *logrus.Logger,
	StoreSession *session.Store,
) ArticleControllerInterface {
	return &articleController{
		Log:          log,
		UseCase:      articleUseCase,
		StoreSession: StoreSession,
	}
}

func (self *articleController) Home(ctx *fiber.Ctx) error {
	articles, err := self.UseCase.FindAll()
	if err != nil {
		return err
	}

	return ctx.Render("home", map[string]any{
		"Articles": articles,
	})
}

func (self *articleController) PageAdminHome(ctx *fiber.Ctx) error {
	articles, err := self.UseCase.FindAll()
	if err != nil {
		return err
	}
	return ctx.Render("admin_home", map[string]any{
		"Articles": articles,
	})
}

func (self *articleController) PageNewArticle(ctx *fiber.Ctx) error {
	return ctx.Render("new_article", nil)
}

func (self *articleController) PageLogin(ctx *fiber.Ctx) error {
	return ctx.Render("login", nil)
}

func (self *articleController) Login(ctx *fiber.Ctx) error {
	username := ctx.FormValue("username")
	password := ctx.FormValue("password")

	if username == "admin" && password == "admin123" {
		sess, _ := self.StoreSession.Get(ctx)
		sess.Set("authenticated", true)
		if err := sess.Save(); err != nil {

		}
		return ctx.Redirect("/admin")
	} else {
		return ctx.Status(fiber.StatusBadRequest).Render("login", fiber.Map{
			"Error": "Invalid username or password",
		})
	}
}

func (self *articleController) Logout(ctx *fiber.Ctx) error {
	sess, _ := self.StoreSession.Get(ctx)
	sess.Destroy()
	return ctx.Redirect("/")
}

func (self *articleController) PageEditArticle(ctx *fiber.Ctx) error {
	articleId, _ := ctx.ParamsInt("id", 0)
	articlePayloadParamIdRequest := dto.ArticlePayloadParamIdRequest{
		Id: int64(articleId),
	}
	article, err := self.UseCase.FindById(&articlePayloadParamIdRequest)
	if err != nil {
		return err
	}
	return ctx.Render("edit_article", map[string]any{
		"Article": article,
		"Tags":    strings.Join(article.Tags, ","),
	})
}

func (self *articleController) Create(ctx *fiber.Ctx) error {
	var articlePayloadBodyRequest = new(dto.ArticlePayloadBodyRequest)
	if err := ctx.BodyParser(articlePayloadBodyRequest); err != nil {
		self.Log.WithError(err).Error("error parsing request body")
		return fiber.ErrBadRequest
	}

	if err := self.UseCase.Create(articlePayloadBodyRequest); err != nil {
		self.Log.WithError(err).Error("error insert article")
		return err
	}

	return ctx.SendStatus(fiber.StatusOK)
}

func (self *articleController) FindById(ctx *fiber.Ctx) error {
	articleId, _ := ctx.ParamsInt("id", 0)
	requestId := dto.ArticlePayloadParamIdRequest{
		Id: int64(articleId),
	}
	article, err := self.UseCase.FindById(&requestId)
	if err != nil {
		return err
	}
	return ctx.Render("detail-article", map[string]any{
		"Article":     article,
		"Tags":        strings.Join(article.Tags, ","),
		"PublishDate": article.PublishDate.Format("January 02 2006"),
	})
}

func (self *articleController) Delete(ctx *fiber.Ctx) error {
	articleId := ctx.QueryInt("id", 0)
	articlePayloadParamIdRequest := dto.ArticlePayloadParamIdRequest{
		Id: int64(articleId),
	}

	if err := self.UseCase.Delete(&articlePayloadParamIdRequest); err != nil {
		return err
	}
	return ctx.Redirect("/admin")
}

func (self *articleController) Update(ctx *fiber.Ctx) error {
	articleId, _ := ctx.ParamsInt("id", 0)
	requestBody := new(dto.ArticlePayloadBodyRequest)
	if err := ctx.BodyParser(requestBody); err != nil {
		return err
	}
	requestId := dto.ArticlePayloadParamIdRequest{
		Id: int64(articleId),
	}

	if err := self.UseCase.Update(&requestId, requestBody); err != nil {
		return err
	}
	return ctx.Redirect("/admin")
}
