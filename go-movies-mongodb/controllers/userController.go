package controllers

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/dogukanozdemir/go-movies-mongodb/database"

	"github.com/dogukanozdemir/go-movies-mongodb/models"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var movieCollection *mongo.Collection = database.OpenCollection(database.Client, "movies")

func GetMovie(c *gin.Context) {

	id := c.Param("id")
	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)

	var movieModel models.Movie
	docId, _ := primitive.ObjectIDFromHex(id)
	err := movieCollection.FindOne(ctx, bson.M{"_id": docId}).Decode(&movieModel)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer cancel()

	c.JSON(http.StatusOK, movieModel)
}

func CreateMovie(c *gin.Context) {
	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	var movieModel models.Movie

	movieModel.Id = primitive.NewObjectID()
	if err := c.BindJSON(&movieModel); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	_, insertErr := movieCollection.InsertOne(ctx, movieModel)
	if insertErr != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": insertErr.Error()})
		return
	}
	defer cancel()

	c.JSON(http.StatusOK, movieModel)

}

func UpdateMovie(c *gin.Context) {
	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	var movie models.Movie

	if err := c.BindJSON(&movie); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	updateResult, err := movieCollection.UpdateOne(ctx, bson.M{"_id": movie.Id}, bson.M{"$set": movie})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Database Error": err.Error()})
		return
	}

	if updateResult.MatchedCount == 0 {
		msg := fmt.Sprintf("No movie with ID %v was found", movie.Id.Hex())
		c.JSON(http.StatusBadRequest, gin.H{"Match Error": msg})
		return
	}
	defer cancel()

	c.JSON(http.StatusOK, movie)
}

func GetAllMovies(c *gin.Context) {
	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)

	documents, err := movieCollection.Find(ctx, bson.D{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var movies []models.Movie
	for documents.Next(ctx) {
		var movie models.Movie
		err := documents.Decode(&movie)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		movies = append(movies, movie)
	}

	defer cancel()
	c.JSON(http.StatusOK, movies)

}

func DeleteMovie(c *gin.Context) {
	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)

	id := c.Param("id")

	docId, _ := primitive.ObjectIDFromHex(id)
	deleteResult, err := movieCollection.DeleteOne(ctx, bson.M{"_id": docId})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if deleteResult.DeletedCount == 0 {
		msg := fmt.Sprintf("Movie with id %v was not found", id)
		c.JSON(http.StatusNotFound, gin.H{"error": msg})
		return
	}
	defer cancel()

	resmsg := fmt.Sprintf("Movie with Id %v deleted", id)
	c.JSON(http.StatusOK, gin.H{"message ": resmsg})

}
