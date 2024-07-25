package handler 

import (
	"github.com/gin-gonic/gin"
	"github.com/Mamvriyskiy/shortLink/tree/develop/src/structure"
	"github.com/Mamvriyskiy/shortLink/tree/develop/src/logger"
	"fmt"
)

func (h *Handler) CreateShortLink(c *gin.Context) {
	var link structure.Link

	if err := c.BindJSON(&link); err != nil {
		logger.Log("Error", "c.BindJSON(&link)", "Error bind json:", err, link)
		return
	}

	fmt.Println("+")
	for true {
		result, err := h.services.LinkServices.CreateShortLink(link)
		if err != nil {
			logger.Log("Error", "h.services.LinkServices.CreateShortLink(link)", "Error create shortlink:", err)
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

	//TODO: clinetID
	linkID, err := h.services.AddLink(link, 1)
	if err != nil {
		logger.Log("Error", " h.services.AddLink(link, 1)", "Error addlink:", err, link.ShortLink, link.LongLink)
	}

	fmt.Println(link, linkID)

	return
}
