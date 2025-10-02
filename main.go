package main

import (
	"ankit/authentication/constants"
	"ankit/authentication/controllers/health_controller"
	// test "ankit/authentication/testing"

	"github.com/gin-gonic/gin"
)

func main() {

	// test.RunDevTest()

	router := gin.Default()
	Routes(router)
	router.GET(constants.Health, healthcontroller.HealthCheck)

	router.Run(":8081") // listen and serve on

}

// func Routes(router *gin.Engine) {
// 	panic("unimplemented")
// }
