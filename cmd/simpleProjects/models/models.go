package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

//Create Struct

type User struct {
	ID       primitive.ObjectID `json:"_id" bson:"_id"`
	Name     string             `json:"name" bson:"name"`
	Email    string             `json:"email" bson:"email"`
	Password string             `json:"password" bson:"password"`
	Posts    *Post              `json:"posts,omitempty" bson:"posts,omitempty"`
}

//dummy database
type allUsers []User

var (
	users = allUsers{
		// {
		// 	ID:       1,
		// 	Name:     "Ritika",
		// 	Email:    "ritika@gmail.com",
		// 	Password: "32fchqsH72@j",
		// },
	}
	//seq = 2
)

type Post struct {
	EID       primitive.ObjectID `json:"pid" bson:"pid"`
	Caption   string             `json:"caption" bson:"caption"`
	ImageURL  string             `json:"imageURL" bson:"imageURL"`
	Timestamp time.Time          `json:"timestamp" bson:"timestamp"`
}

//dummy database
type allPosts []Post

var (
	posts = allPosts{
		// {
		// 	EID:      1,
		// 	Caption:  "MyJourney",
		// 	ImageURL: "https://upload.wikimedia.org/wikipedia/commons/thumb/b/b6/Image_created_with_a_mobile_phone.png/1200px-Image_created_with_a_mobile_phone.png",
		// 	Timestamp:  2014-05:28:06.801064-04:00,
		// },
		// {
		// 	EID:      2,
		// 	Caption:  "Car",
		// 	ImageURL: "https://upload.wikimedia.org/wikipedia/commons/thumb/b/b6/Image_created_with_a_mobile_phone.png/1200px-Image_created_with_a_mobile_phone.png",
		// 	Timestamp:  2014-05:28:06.801064-04:00,
		// },
	}

	//pseq = 3
)
