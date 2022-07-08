package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	Id     primitive.ObjectID  `bson:"_id"`
	Name   string        	   `json:"name"   bson:"name"`
	Gender string        	   `json:"gender" bson:"gender"`
	Age    int           	   `json:"age"    bson:"age"`
	User_id       string       `json:"user_id"`
}
