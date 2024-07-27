package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"github.com/Mamvriyskiy/shortLink/tree/develop/src/service"
	"fmt"
)

type Handler struct {
	services *service.Services
}

func NewHandler(services *service.Services) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRouters() *gin.Engine {
	router := gin.New()

	router.Static("/css", "./template/css")
	router.LoadHTMLGlob("template/*.html")

	router.GET("/:path", func(c *gin.Context) {
		currentURL := c.Param("path")
		// fmt.Println(currentURL)
		longLink, err := h.GetLongLink(currentURL)
		if err != nil {
			//TODO: error status
			return
		}
		fmt.Println(longLink, currentURL, err)

		c.Redirect(http.StatusFound, longLink)
		return
	})

	app := router.Group("/app")
	app.GET("/create", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "index.html", nil)
	})

	app.GET("/sign-up", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "signup.html", nil)
	})

	app.GET("/sign-in", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "signin.html", nil)
	})
	


	api := router.Group("/api")
	api.POST("/create", h.CreateShortLink)
	api.POST("/sing-up", h.RegisterUser)
	api.POST("/sing-in", h.GetUser)

	

	return router
}
