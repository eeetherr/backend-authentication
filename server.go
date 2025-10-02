package main

import (
	"ankit/authentication/constants"
	controllers "ankit/authentication/controllers/auth_contoller"

	"github.com/gin-gonic/gin"
)

func Routes(router *gin.Engine) {
	auth := router.Group(constants.Auth)
	// comms := router.Group(constants.Comms)
	// users := router.Group(constants.Users)
	AuthPath(auth)
	// CommsPath(comms)
	// UsersPath(users)
}

func AuthPath(authRoute *gin.RouterGroup) {
	r1 := authRoute.Group(constants.V1)
	r1.POST(constants.SignUp, controllers.SignUp)

}

// func CommsPath(commsRoute *gin.RouterGroup){
// 	r1 := commsRoute.Group(constants.V1)
// }

// func UsersPath(usersRoute *gin.RouterGroup){
// 	r1 := usersRoute.Group(constants.V1)
// }
