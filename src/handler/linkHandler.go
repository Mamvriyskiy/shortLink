package handler 

import (
	"github.com/gin-gonic/gin"
	"github.com/Mamvriyskiy/shortLink/tree/develop/src/structure"
	// "encoding/json"
	"fmt"
)

func (h *Handler) CreateShortLink(c *gin.Context) {
	var link structure.Link

	if err := c.BindJSON(&link); err != nil {
		//TODO: log error
		fmt.Println("error bindjson", err)
		return
	}

	fmt.Println("+")
	for true {
		result, err := h.services.LinkServices.CreateShortLink(link)
		if err != nil {
			fmt.Println("error create short link")
		}

		flag, err := h.services.CheckDuplicateShortLink(result)
		if err != nil {
			fmt.Println(err)
			//TODO: error
			return
		}

		if !flag {
			link.ShortLink = result
			break
		}
	}

	linkID, err := h.services.AddLink(link, 1)
	if err != nil {
		fmt.Println(err)
		//TODO: error
	}

	fmt.Println(link, linkID)

	return
}
