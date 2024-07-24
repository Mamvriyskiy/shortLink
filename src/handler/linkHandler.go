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

    _, err := h.services.LinkServices.CreateShortLink(link)
	if err != nil {
		fmt.Println("error create short link")
	}

	return
}
