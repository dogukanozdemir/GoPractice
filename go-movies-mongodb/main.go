package main

import (
	"net/http"

	"github.com/dogukanozdemir/go-movies-mongodb/controllers"
	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})

	router.SetTrustedProxies([]string{"localhost"})
	router.LoadHTMLGlob("assets/*.html")
	router.Static("/assets", "./assets")
	router.GET("/movies/:id", controllers.GetMovie)
	router.GET("/movies", controllers.GetAllMovies)
	router.PUT("/movies", controllers.UpdateMovie)
	router.DELETE("/movies/:id", controllers.DeleteMovie)
	router.POST("/movies", controllers.CreateMovie)
	router.Run(":8080")
}
