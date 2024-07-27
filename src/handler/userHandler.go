package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/Mamvriyskiy/shortLink/tree/develop/src/structure"
	"fmt"
)

type User struct {
	Login string `json:"login"`
	Password string `json:"password"`
	Email string `json:"email"`
}

func (h *Handler) RegisterUser(*gin.Context) {
	var user User
	if err := c.BindJSON(&user); err != nil {
		//TODO: logger
		return
	}

	fmt.Println(user)
	userID, err := h.services.CreateUser(user)
	fmt.Println(userID, err)

	return 
}

func (h *Handler) GetUser(*gin.Context) {
	var user User
	if err := c.BindJSON(&user); err != nil {
		//TODO: logger
		return
	}

	user := structure.User{
		Password: password,
		Email: email,
	}

	userID, err := h.services.GetUser(user)
	_ = userID
	_ = err

	return 
}
