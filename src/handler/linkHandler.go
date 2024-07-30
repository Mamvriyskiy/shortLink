package handler 

import (
	"github.com/gin-gonic/gin"
	"github.com/Mamvriyskiy/shortLink/tree/develop/src/structure"
	"github.com/Mamvriyskiy/shortLink/tree/develop/src/logger"
	"github.com/spf13/viper"
	"net/http"
	"fmt"
)

func (h *Handler) GetLongLink(shortLink string) (string, error) {
	return h.services.GetLongLink(shortLink)
}

func (h *Handler) CreateShortLink(c *gin.Context) {
	clientID, ok := c.Get("userId")
	if !ok {
		logger.Log("Warning", "Get", "Error get userID from context", nil, "userId")
		return
	}

	var link structure.Link

	if err := c.BindJSON(&link); err != nil {
		logger.Log("Error", "c.BindJSON(&link)", "Error bind json:", err, link)
		return
	}

	for true {
		result, err := h.services.CreateShortLink(link)
		if err != nil {
			logger.Log("Error", "h.services.CreateShortLink(link)", "Error create shortlink:", err)
		}

		flag, err := h.services.CheckDuplicateShortLink(result)
		if err != nil {
			logger.Log("Error", "h.services.CheckDuplicateShortLink(result)", "Error CheckDuplicateShortLink:", err)
			return
		}

		if !flag {
			link.ShortLink = result
			break
		}
	}

	logger.Log("Info", "", fmt.Sprintf("Create shortLink %s", link.ShortLink), nil, "")

	var intVal float64
	if val, ok := clientID.(float64); ok {
		intVal = val
	} else {
		logger.Log("Error", "userID.(float64)", "Error:", nil, "")
	}

	_, err := h.services.AddLink(link, int(intVal))
	if err != nil {
		logger.Log("Error", " h.services.AddLink(link, 1)", "Error addlink:", err, link.ShortLink, link.LongLink)
		return
	}

	resultLink := fmt.Sprintf("https://localhost:%s/", viper.GetString("server.port")) + link.ShortLink
	viper.GetString("server.port")
	c.JSON(http.StatusOK, map[string]interface{}{
		"shortlink": resultLink,
	})

	return
}
