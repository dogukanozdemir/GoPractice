package controllers

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/dogukanozdemir/GoPractice/go-mongodb/database"
	"github.com/dogukanozdemir/GoPractice/go-mongodb/models"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var userCollection *mongo.Collection = database.OpenCollection(database.Client, "user")

func GetUser(c *gin.Context) {
	id := c.Param("id")
	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)

	var user models.User
	objId, _ := primitive.ObjectIDFromHex(id)
	err := userCollection.FindOne(ctx, bson.M{"_id": objId}).Decode(&user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer cancel()

	c.JSON(http.StatusOK, user)
}

func CreateUser(c *gin.Context) {

	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	var user models.User

	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	user.Id = primitive.NewObjectID()

	_, insertErr := userCollection.InsertOne(ctx, user)
	if insertErr != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": insertErr.Error()})
		return
	}
	defer cancel()

	c.JSON(http.StatusOK, user)
}

func GetAllUsers(c *gin.Context) {
	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)

	cursor, err := userCollection.Find(ctx, bson.D{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var users []models.User
	for cursor.Next(ctx) {
		var user models.User
		err := cursor.Decode(&user)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		users = append(users, user)
	}
	defer cancel()

	c.JSON(http.StatusOK, users)
}

func DeleteUser(c *gin.Context) {
	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)

	id := c.Param("id")
	objId, _ := primitive.ObjectIDFromHex(id)
	deleteResult, err := userCollection.DeleteOne(ctx, bson.M{"_id": objId})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if deleteResult.DeletedCount == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}
	defer cancel()

	msg := fmt.Sprintf("User with id %v deleted", id)
	c.JSON(http.StatusOK, gin.H{"message": msg})
}

func UpdateUser(c *gin.Context) {
	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	var user models.User

	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	_, err := userCollection.UpdateOne(ctx, bson.M{"_id": user.Id}, bson.M{"$set": user})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	defer cancel()
	c.JSON(http.StatusOK, user)
}
