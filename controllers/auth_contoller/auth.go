package controllers

import (
	"net/http"

	"ankit/authentication/dto"
	"ankit/authentication/services"

	"github.com/gin-gonic/gin"
)

func SignUp(c *gin.Context) {
	var req dto.SignUpRequest

	// 1. Bind JSON body to DTO
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	// 2. Call service layer
	service := services.AuthService{}
	err := service.SignUp(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// 3. Respond back
	c.JSON(http.StatusCreated, gin.H{
		"message": "User registered successfully",
	})
}
