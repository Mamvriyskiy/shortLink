package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"github.com/Mamvriyskiy/shortLink/tree/develop/src/service"
)

type Handler struct {
	services *service.Services
}

func NewHandler(services *service.Services) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRouters() *gin.Engine {
	router := gin.New()

	// router.Static("/css", "./template/css")
	// router.LoadHTMLGlob("template/*.html")

	app := router.Group("/app")
	app.GET("/", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "menu.html", nil)
	})

	api := router.Group("/api")
	api.GET("/create", h.CreateShortLink)

	return router
}
