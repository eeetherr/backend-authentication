package main

import (
	"ankit/authentication/constants"
	controllers "ankit/authentication/controllers/auth_contoller"
	"ankit/authentication/database"
	"ankit/authentication/utils"
	"fmt"

	"github.com/gin-gonic/gin"
)

func init() {
	fmt.Println("Loading configs...") // log is initialized after configs setup
	paths := []string{constants.ConfigPath}
	utils.SetupConfig(paths)

	fmt.Println("Connecting database........")
	database.InitDB()
}

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

	r1.POST(constants.VerifyAuth, controllers.VerifyAuth)

	r1.POST(constants.Login, controllers.Login)

}

//func CommsPath(commsRoute *gin.RouterGroup) {
//	r1 := commsRoute.Group(constants.V1)
//	r1.POST(constants.ConfirmCode, controllers.SignUp)
//}

// func UsersPath(usersRoute *gin.RouterGroup){
// 	r1 := usersRoute.Group(constants.V1)
// }
