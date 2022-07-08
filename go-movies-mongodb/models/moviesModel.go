package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Movie struct {
	Id       primitive.ObjectID `bson:"_id"`
	Isbn     string             `json"isbn"      bson:"isbn"`
	Title    string             `json:"name"     bson:"name"`
	Director *Director          `json:"director" bson:"director"`
}

type Director struct {
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
}
