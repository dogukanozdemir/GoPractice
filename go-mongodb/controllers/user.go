package controllers

import (
	"context"
	"fmt"
	"log"
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

func GetUser(c * gin.Context) {
	id := c.Param("id")
	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	log.Printf("%v", id)

	var user models.User
	err := userCollection.FindOne(ctx, bson.M{"user_id": id}).Decode(&user)
	defer cancel()
	if err != nil {
		c.JSON(http.StatusInternalServerError,gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, user)
}

func CheckDB(c *gin.Context) {
	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	databases, _ := database.Client.ListDatabaseNames(ctx, bson.M{})
	log.Printf("%v", databases)
	collections, _ := database.Client.Database("go-mongodb").ListCollectionNames(ctx, bson.M{})
	log.Printf("%v", collections)
	defer cancel()
	
}

func CreateUser(c *gin.Context) {
	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	var user models.User

	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	user.Id = primitive.NewObjectID()
	user.User_id = user.Id.Hex()

	_, insertErr := userCollection.InsertOne(ctx, user)
	if insertErr != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error" : insertErr.Error()})
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

func DeleteUser(c * gin.Context) {
	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)

	id := c.Param("id")
	_, err := userCollection.DeleteOne(ctx, bson.M{"user_id": id})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer cancel()

	msg := fmt.Sprintf("User with id %v deleted", id)
	c.JSON(http.StatusOK, gin.H{"message": msg})
}

func UpdateUser(c * gin.Context){
	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	var user models.User

	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	_, err := userCollection.UpdateOne(ctx, bson.M{"user_id": user.User_id}, bson.M{"$set": user})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer cancel()

	c.JSON(http.StatusOK, user)
}
