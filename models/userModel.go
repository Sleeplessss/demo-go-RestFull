package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)
// User struct NOTE user struct
type User struct {
	ID     primitive.ObjectID `bson:"_id,omitempty"`
	Name string `bson:"name,omitempty"`
	Surname string `bson:"surname,omitempty"`
}