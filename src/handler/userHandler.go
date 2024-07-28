package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/Mamvriyskiy/shortLink/tree/develop/src/structure"
	"net/http"
	"fmt"
)

func (h *Handler) RegisterUser(c *gin.Context) {
	var user structure.User
	if err := c.BindJSON(&user); err != nil {
		//TODO: logger
		return
	}

	fmt.Println(user)
	userID, err := h.services.CreateUser(user)
	fmt.Println(userID, err)

	c.JSON(http.StatusOK, map[string]interface{}{})

	return 
}

func (h *Handler) GetUser(c *gin.Context) {
	var user structure.User
	if err := c.BindJSON(&user); err != nil {
		fmt.Println(err)
		//TODO: logger
		return
	}

	fmt.Println(user)
	userID, err := h.services.GetUser(user)
	//TODO: user no search
	if err != nil {
		return 
	}

	fmt.Println(userID)

	token, err := h.services.CreateToken(userID)
	if err != nil {
		return 
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"token": token,
	})

	return 
}
