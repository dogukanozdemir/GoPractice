package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func index(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", nil)
}

type Movie struct {
	ID       string    `json:"id"`
	Isbn     string    `json"isbn"`
	Title    string    `json:"title"`
	Director *Director `json:"director"`
}

type Director struct {
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
}

var movies []Movie

func getMovies(ctx *gin.Context) {
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, movies)
}

func getMovie(ctx *gin.Context) {
	ctx.Header("Content-Type", "application/json")
	id := ctx.Param("id")
	for _, item := range movies {
		if item.ID == id {
			ctx.JSON(http.StatusOK, item)
			return
		}
	}
}
func deleteMovie(ctx *gin.Context) {
	ctx.Header("Content-Type", "application/json")
	id := ctx.Param("id")
	var deletedMovie Movie
	for index, item := range movies {

		if item.ID == id {
			deletedMovie = movies[index]
			movies = append(movies[:index], movies[index+1:]...)
		}
	}
	ctx.JSON(http.StatusOK, deletedMovie)
}
func createMovie(ctx *gin.Context) {
	ctx.Header("Content-Type", "application/json")
	var movie Movie
	if err := ctx.BindJSON(&movie); err == nil {
		movie.ID = strconv.Itoa(len(movies))
		movies = append(movies, movie)
		ctx.JSON(http.StatusOK, movie)
	} else {
		ctx.JSON(http.StatusInternalServerError, "error while binding JSON")
	}
}
func updateMovie(ctx *gin.Context) {
	ctx.Header("Content-Type", "application/json")
	id := ctx.Param("id")
	for index, item := range movies {
		if item.ID == id {
			movies = append(movies[:index], movies[index+1:]...)
			var movie Movie
			movie.ID = id
			if err := ctx.BindJSON(&movie) ; err == nil {
				movies = append(movies, movie)
				ctx.JSON(http.StatusOK,movie)
			} else {
				ctx.JSON(http.StatusInternalServerError, "error while binding JSON")
			}
		}
	}
}

func main() {
	movies = append(movies, Movie{ID: "0", Isbn: "4387225", Title: "Movie1", Director: &Director{FirstName: "John", LastName: "Doe"}})
	movies = append(movies, Movie{ID: "1", Isbn: "4387235", Title: "Movie2", Director: &Director{FirstName: "Steve", LastName: "Spielberg"}})
	router := gin.Default()
	router.SetTrustedProxies([]string{"localhost"})
	router.LoadHTMLGlob("assets/*.html")
	router.Static("/assets", "./assets")
	router.GET("/", index)
	router.GET("/movies/:id", getMovie)
	router.POST("/movies", createMovie)
	router.PUT("/movies/:id", updateMovie)
	router.DELETE("/movies/:id", deleteMovie)
	router.GET("/movies", getMovies)
	router.Run(":8080")
}
