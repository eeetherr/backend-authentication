package controllers

import (
	"ankit/authentication/dto/auth"
	"net/http"

	"ankit/authentication/services"

	"github.com/gin-gonic/gin"
)

func SignUp(c *gin.Context) {
	var req auth.SignUpRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}
	service := services.AuthService{}
	err := service.SignUp(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "User registered successfully",
	})
}
