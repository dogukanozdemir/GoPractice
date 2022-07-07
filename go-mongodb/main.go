package main

import (
	"log"
	"net/http"
	"os"

	"github.com/dogukanozdemir/GoPractice/go-mongodb/controllers"
	"github.com/joho/godotenv"
	"github.com/julienschmidt/httprouter"
	"gopkg.in/mgo.v2"
)

func main() {

	r := httprouter.New()
	uc := controllers.NewUserController(getSession())
	r.GET("/user/:id", uc.GetUser)
	r.POST("/user", uc.CreateUser)
	r.DELETE("/user/:id", uc.DeleteUser)
	http.ListenAndServe(":9000", r)

}

func getSession() *mgo.Session {

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	url := os.Getenv("MONGODB_URL")

	s, err := mgo.Dial(url)

	if err != nil {
		panic(err)
	}

	return s
}
