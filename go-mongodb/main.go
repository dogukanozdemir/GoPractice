package main

import (
	"github.com/dogukanozdemir/GoPractice/go-mongodb/controllers"
	"github.com/gin-gonic/gin"
)



func main() {

	r := gin.New()
	r.Use(gin.Logger())

	r.GET("/check", controllers.CheckDB)
 	r.GET("/user/:id", controllers.GetUser)
	r.POST("/user", controllers.CreateUser)
	r.DELETE("/user/:id", controllers.DeleteUser)
	r.GET("/users", controllers.GetAllUsers)
	r.PUT("/user", controllers.UpdateUser)

	r.Run(":9000")
}
