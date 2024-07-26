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
	password := "abcd"
	email := "asdf@mail.ru"
	login := "Mamre32"
	
	user := structure.User{
		Login: login,
		Password: password,
		Email: email,
	}

	fmt.Println(user)
	userID, err := h.services.CreateUser(user)
	_ = userID
	_ = err

	return 
}

func (h *Handler) GetUser(*gin.Context) {
	email := "asdf@mail.ru"
	password := "abcd"

	user := structure.User{
		Password: password,
		Email: email,
	}

	userID, err := h.services.GetUser(user)
	_ = userID
	_ = err
	
	return 
}
