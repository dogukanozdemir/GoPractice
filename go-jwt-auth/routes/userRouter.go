package routes

import (
	controller "github.com/dogukanozdemir/GoPractice/go-jwt-auth/controllers"
	"github.com/dogukanozdemir/GoPractice/go-jwt-auth/middleware"

	"github.com/gin-gonic/gin"
)

func userRoutes(IncomingRoutes *gin.Engine) {
	IncomingRoutes.Use(middleware.Authenticate())
	IncomingRoutes.GET("/users", controller.GetUsers())
	IncomingRoutes.GET("/users/:user_id", controller.GetUser())
}
