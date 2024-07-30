package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"github.com/Mamvriyskiy/shortLink/tree/develop/src/service"
	"fmt"
	jwt "github.com/dgrijalva/jwt-go"
	"strings"
	"github.com/Mamvriyskiy/shortLink/tree/develop/src/logger"
)

const (
	signingKey = "jaskljfkdfndnznmckmdkaf3124kfdlsf"
)

type Handler struct {
	services *service.Services
}

func NewHandler(services *service.Services) *Handler {
	return &Handler{services: services}
}

type markerClaims struct {
	jwt.StandardClaims
}

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Проверить URL запроса
		if !strings.HasPrefix(c.Request.URL.Path, "/api"){
			// Если URL начинается с /auth, пропустить проверку JWT
			c.Next()
			return
		}

		// Получить токен из заголовка запроса или из куки
		tokenString := c.GetHeader("Authorization")
		var err error
		fmt.Println("Token:", tokenString)
		if tokenString == "" {
			// Если токен не найден в заголовке, попробуйте из куки
			tokenString, err = c.Cookie("jwt")
			if err != nil {
				logger.Log("Error", "c.Cookie(jwt)", "Error", err, "jwt")
			}
		}

		// Проверить, что токен не пустой
		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Empty token"})
			c.Abort()
			return
		}

		// Парсинг токена
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			// Здесь нужно вернуть ключ для проверки подписи токена.
			// В реальном приложении, возможно, это будет случайный секретный ключ.
			return []byte(signingKey), nil
		})
		// Проверить наличие ошибок при парсинге токена
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized", "detail:": err.Error()})
			c.Abort()
			return
		}

		// Добавить данные из токена в контекст запроса
		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			c.Set("userId", claims["userId"])
			c.Next()
		} else {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}
	}
}

func (h *Handler) InitRouters() *gin.Engine {
	router := gin.New()

	router.Use(AuthMiddleware())

	//router.Static("/css", "./template/css")
	router.StaticFS("/css", http.Dir("./template/css"))

	router.Static("/script", "./template/script")
	
	router.LoadHTMLGlob("template/*.html")

	router.GET("/:path", func(c *gin.Context) {
		currentURL := c.Param("path")
		longLink, err := h.GetLongLink(currentURL)
		if err != nil {
			logger.Log("Error", "h.GetLongLink(currentURL)", "Error get long link for% ", nil, currentURL)
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

	auth := router.Group("/auth")
	auth.GET("/sign-up", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "signup.html", nil)
	})

	auth.GET("/sign-in", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "signin.html", nil)
	})

	auth.POST("/sing-up", h.RegisterUser)
	auth.POST("/sing-in", h.GetUser)

	api := router.Group("/api")
	api.POST("/create", h.CreateShortLink)

	return router
}
