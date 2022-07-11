package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func index(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", nil)
}


func main() {
	router := gin.Default()
	router.LoadHTMLGlob("assets/*.html")
	router.Static("/assets", "./assets")
	router.GET("/", index)
	router.Run(":8080")
}