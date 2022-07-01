package routes

import (
	controller "github.com/dogukanozdemir/GoPractice/go-jwt-auth/controllers"

	"github.com/gin-gonic/gin"
)

func AuthRoutes(IncomingRoutes *gin.Engine) {
	IncomingRoutes.POST("users/signup", controller.Signup())
	IncomingRoutes.POST("users/login", controller.Login())
}
