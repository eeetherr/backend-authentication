package healthcontroller

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

func HealthCheck(ctx *gin.Context){
	ctx.JSON(http.StatusOK, gin.H{
		"status": "OK",
		"message": "Service is healthy",
	})
}