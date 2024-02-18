package handler

import (
	"go-wasm-srv/config"

	"github.com/gin-gonic/gin"
)

type Handler struct {
}

func NewHandler(cfg config.Config) Handler {
	return Handler{}
}

func (h *Handler) ServeFile(ctx *gin.Context) {
	path := ctx.Request.URL.Path
	switch path {
	case "/", "/index.html":
		path = "static/index.html"
	case "/favicon.ico":
		path = "static/favicon.ico"
	}
	ctx.File(path)
}
