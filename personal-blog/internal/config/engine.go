package config

import (
	"github.com/gofiber/template/html/v2"
	"net/http"
)

func NewEngine(fs http.FileSystem, ext string) *html.Engine {
	return html.NewFileSystem(fs, ext)
}
