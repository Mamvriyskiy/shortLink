package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Handler struct {
	services *service.Services
}

func NewHandler(services *service.Services) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRouters() *gin.Engine {
	router := gin.New()

	router.Static("/css", "./index/css")
	router.LoadHTMLGlob("index/*.html")

	app := router.Group("/app")
	app.GET("/", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "menu.html", nil)
	})

	return h.InitRouters()
}
