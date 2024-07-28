package handler 

import (
	"github.com/gin-gonic/gin"
	"github.com/Mamvriyskiy/shortLink/tree/develop/src/structure"
	"github.com/Mamvriyskiy/shortLink/tree/develop/src/logger"
	"net/http"
	//"fmt"
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

	var intVal float64
	if val, ok := clientID.(float64); ok {
		intVal = val
	} else {
		logger.Log("Error", "userID.(float64)", "Error:", nil, "")
	}

	_, err := h.services.AddLink(link, int(intVal))
	if err != nil {
		logger.Log("Error", " h.services.AddLink(link, 1)", "Error addlink:", err, link.ShortLink, link.LongLink)
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"shortlink": "https://localhost:8000/" + link.ShortLink,
	})

	return
}
