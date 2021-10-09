package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

//Create Struct

type User struct {
	ID       primitive.ObjectID `json:"_id" bson:"_id,omitempty"`
	Name     string             `json:"name" bson:"name"`
	Email    string             `json:"email" bson:"email"`
	Password string             `json:"password" bson:"password"`
	Posts    *Post              `json:"posts,omitempty" bson:"posts,omitempty"`
}

type Post struct {
	EID       primitive.ObjectID `json:"pid" bson:"pid,omitempty"`
	Caption   string             `json:"caption" bson:"caption"`
	ImageURL  string             `json:"imageURL" bson:"imageURL"`
	Timestamp time.Time          `json:"timestamp" bson:"timestamp"`
}
