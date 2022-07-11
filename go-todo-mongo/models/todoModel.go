package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Todo struct {
	ID     primitive.ObjectID `bson:"_id"`
	Name   string             `json:"name"		bson:"name"`
	Status string             `json:"status"	bson:"status"`
}
