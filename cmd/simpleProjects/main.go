package main

import (
	// "io"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/ritikasarkar18/Creating-APIs/cmd/simpleProjects/helper"
	"github.com/ritikasarkar18/Creating-APIs/cmd/simpleProjects/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//Connection mongoDB with helper class
var collection = helper.ConnectDB()

//dummy database
type allUsers []models.User

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

//dummy database
type allPosts []models.Post

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

// create user
func createUser(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		w.Header().Set("Content-Type", "application/json")

		var newUser models.User

		//var newUser User //another user of type user
		reqBody, err := ioutil.ReadAll(r.Body)
		if err != nil {
			fmt.Fprintf(w, "Kindly enter data properly")
		}

		json.Unmarshal(reqBody, &newUser)
		result, err := collection.InsertOne(context.TODO(), newUser)

		if err != nil {
			helper.GetError(err, w)
			return
		}

		json.NewEncoder(w).Encode(result)
		// users = append(users, newUser)

		// w.WriteHeader(http.StatusCreated)

		// json.NewEncoder(w).Encode(newUser)

	} else {
		http.Redirect(w, r, "/", http.StatusFound)
	}
}

//get an user
func getUser(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		w.Header().Set("Content-Type", "application/json")
		var user models.User
		userID := strings.Split(r.URL.String()[len(`/users/`):], `/`)
		//userID := mux.Vars(r)["id"]
		uid, _ := primitive.ObjectIDFromHex(userID[0])

		filter := bson.M{"_id": uid}
		err := collection.FindOne(context.TODO(), filter).Decode(&user)

		if err != nil {
			helper.GetError(err, w)
			return
		}

		json.NewEncoder(w).Encode(user)

		///* the following commented code is for standalone API (without mongo)*/
		// for _, singleUser := range users {
		// 	if singleUser.ID == uid {
		// 		json.NewEncoder(w).Encode(singleUser)
		// 	}
		// }

	} else {
		http.Redirect(w, r, "/", http.StatusFound)
	}

}

//create a post
// When creating a post, we get data from the user’s end. The user enters data which is in the form of http Request data.
// The request data is not is a human-readable format hence we use the package ioutil to convert it into a slice.
func createPost(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		w.Header().Set("Content-Type", "application/json")
		//var newPost Post //another post of type post
		var newPost models.Post

		reqBody, err := ioutil.ReadAll(r.Body)
		if err != nil {
			fmt.Fprintf(w, "Kindly enter data properly")
		}

		json.Unmarshal(reqBody, &newPost)

		result, err := collection.InsertOne(context.TODO(), newPost)
		if err != nil {
			helper.GetError(err, w)
			return
		}

		json.NewEncoder(w).Encode(result)

		// posts = append(posts, newPost)
		// w.WriteHeader(http.StatusCreated)

		// json.NewEncoder(w).Encode(newPost)

	} else {
		http.Redirect(w, r, "/", http.StatusFound)
	}
}

//get a post
func getOnePost(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		w.Header().Set("Content-Type", "application/json")
		var post models.Post
		postID := strings.Split(r.URL.String()[len(`/posts/`):], `/`)
		//r.URL.Path[len("/posts/{id}"):] --- wrong
		//postID := mux.Vars(r)["id"]
		pid, _ := primitive.ObjectIDFromHex(postID[0])

		// for _, singlePost := range posts {
		// 	if singlePost.EID == pid {
		// 		json.NewEncoder(w).Encode(singlePost)
		// 	}
		// }

		filter := bson.M{"pid": pid}
		err := collection.FindOne(context.TODO(), filter).Decode(&post)

		if err != nil {
			helper.GetError(err, w)
			return
		}

		json.NewEncoder(w).Encode(post)

	} else {
		http.Redirect(w, r, "/", http.StatusFound)
	}

}

//get all posts
func getAllPosts(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		w.Header().Set("Content-Type", "application/json")

		var posts models.Post

		userID := strings.Split(r.URL.String()[len(`/posts/users/`):], `/`)
		//r.URL.Path[len("/posts/{id}"):] --- wrong
		//postID := mux.Vars(r)["id"]
		upid, _ := primitive.ObjectIDFromHex(userID[0])

		// for _, singleUser := range users {
		// 	if singleUser.ID == upid {
		// 		json.NewEncoder(w).Encode(posts)
		// 	}
		// }

		filter := bson.M{"_id": upid}
		cursor, err := collection.Find(context.TODO(), filter) // #### check

		if err != nil {
			helper.GetError(err, w)
			return
		}
		fmt.Println(cursor)
		json.NewEncoder(w).Encode(posts)

	} else {
		http.Redirect(w, r, "/", http.StatusFound)
	}

}

func homeLink(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome")
}

func main() {
	// router:=mux.NewRouter().StrictSlash(true)
	// helloHandler := func(hw http.ResponseWriter, req *http.Request) {
	// 	io.WriteString(hw, "Hello Ritika\n")
	// }
	http.HandleFunc("/", homeLink)
	http.HandleFunc("/users", createUser)
	http.HandleFunc("/users/", getUser)
	http.HandleFunc("/posts", createPost)
	http.HandleFunc("/posts/", getOnePost)
	http.HandleFunc("/posts/users/", getAllPosts) // ### not working

	log.Println("Listing for requests at http://localhost:8000/")
	log.Fatal(http.ListenAndServe(":8000", nil))
}
