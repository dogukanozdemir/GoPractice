package main

import (
	"net/http"

	controller "github.com/dogukanozdemir/go-todo-mongo/controllers"

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
	router.GET("/todos", controller.GetTodos)
	router.GET("/todo/:id", controller.GetTodo)

	router.POST("/todo", controller.AddTodo)
	router.DELETE("/todo/:id", controller.DeleteTodo)
	router.DELETE("/todos",controller.ClearAll)
	router.PUT("/todo/:id",controller.UpdateTodo)
	router.Run(":8080")

}
