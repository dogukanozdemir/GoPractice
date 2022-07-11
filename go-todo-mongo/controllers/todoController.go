package controller

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/dogukanozdemir/go-todo-mongo/database"
	"github.com/dogukanozdemir/go-todo-mongo/models"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var todoCollection *mongo.Collection = database.OpenCollection(database.Client, "todos")

func GetTodo(c *gin.Context) {
	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)

	id := c.Param("id")
	objId, _ := primitive.ObjectIDFromHex(id)

	var todo models.Todo
	err := todoCollection.FindOne(ctx, bson.M{"_id": objId}).Decode(&todo)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error ": err.Error()})
	}

	defer cancel()
	c.JSON(http.StatusOK, todo)
}

func GetTodos(c *gin.Context) {
	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)

	findResult, err := todoCollection.Find(ctx, bson.D{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"FindError": err.Error()})
		return
	}

	var todos []models.Todo
	for findResult.Next(ctx) {
		var todo models.Todo
		err := findResult.Decode(&todo)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"Decode Error": err.Error()})
			return
		}
		todos = append(todos, todo)
	}
	defer cancel()

	c.JSON(http.StatusOK, todos)
}

func DeleteTodo(c *gin.Context) {
	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)

	id := c.Param("id")
	objId, _ := primitive.ObjectIDFromHex(id)
	deleteResult, err := todoCollection.DeleteOne(ctx, bson.M{"_id": objId})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if deleteResult.DeletedCount == 0 {
		msg := fmt.Sprintf("No todo with id : %v was found, no deletion occurred.", id)
		c.JSON(http.StatusBadRequest, gin.H{"error": msg})
		return
	}
	defer cancel()

	msg := fmt.Sprintf("todo with id : %v is was deleted successfully.", id)
	c.JSON(http.StatusOK, gin.H{"success": msg})

}

func AddTodo(c *gin.Context) {
	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)

	var todo models.Todo
	if err := c.BindJSON(&todo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	todo.ID = primitive.NewObjectID()

	_, err := todoCollection.InsertOne(ctx, todo)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer cancel()

	msg := fmt.Sprintf("A new todo '%v' created!", todo.Name)
	c.JSON(http.StatusOK, gin.H{"success": msg})
}
