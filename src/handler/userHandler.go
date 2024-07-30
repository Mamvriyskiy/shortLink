package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/Mamvriyskiy/shortLink/tree/develop/src/structure"
	"github.com/Mamvriyskiy/shortLink/tree/develop/src/logger"
	"net/http"
	"fmt"
)

func (h *Handler) RegisterUser(c *gin.Context) {
	var user structure.User
	if err := c.BindJSON(&user); err != nil {
		logger.Log("Error", "c.BindJSON(&user)", "Error bind json:", err)
		return
	}
	
	_, err := h.services.CreateUser(user)
	if err != nil {
		logger.Log("Error", "h.services.CreateUser(user)", "Error create user:", err, user.Email)
		return 
	}

	logger.Log("Info", "", fmt.Sprintf("User %s was created", user.Email), nil, "")

	c.JSON(http.StatusOK, map[string]interface{}{})

	return 
}

func (h *Handler) GetUser(c *gin.Context) {
	var user structure.User
	if err := c.BindJSON(&user); err != nil {
		logger.Log("Error", "c.BindJSON(&user)", "Error bind json:", err)
		return
	}
	
	userID, err := h.services.GetUser(user)
	if err != nil {
		logger.Log("Error", "h.services.GetUser(user)", "User no search", err)
		return 
	}

	logger.Log("Info", "", fmt.Sprintf("Get user: %s", user.Email), nil, "")

	token, err := h.services.CreateToken(userID)
	if err != nil {
		logger.Log("Error", "h.services.CreateToken(userID)", fmt.Sprintf("Error created token for %s", user.Email), err)
		return 
	}

	logger.Log("Info", "", fmt.Sprintf("Token for % was created", user.Email), nil, "")

	c.JSON(http.StatusOK, map[string]interface{}{
		"token": token,
	})

	return 
}
