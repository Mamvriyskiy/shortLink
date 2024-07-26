package handler 

import (
	"github.com/gin-gonic/gin"
	"github.com/Mamvriyskiy/shortLink/tree/develop/src/structure"
	"github.com/Mamvriyskiy/shortLink/tree/develop/src/logger"
	"net/http"
	"fmt"
)

func (h *Handler) GetLongLink(shortLink string) (string, error) {
	return h.services.GetLongLink(shortLink)
}

func (h *Handler) CreateShortLink(c *gin.Context) {
	var link structure.Link

	if err := c.BindJSON(&link); err != nil {
		logger.Log("Error", "c.BindJSON(&link)", "Error bind json:", err, link)
		return
	}

	var result string
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

	result = "https://localhost:8000/" + result
	//TODO: clinetID
	linkID, err := h.services.AddLink(link, 1)
	if err != nil {
		logger.Log("Error", " h.services.AddLink(link, 1)", "Error addlink:", err, link.ShortLink, link.LongLink)
	}

	fmt.Println(link, linkID)

	c.JSON(http.StatusOK, map[string]interface{}{
		"shortlink": result,
	})

	return
}
