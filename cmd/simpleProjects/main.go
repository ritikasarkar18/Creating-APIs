package main

import (
	// "io"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//Connection mongoDB with helper class
collection := helper.ConnectDB()


// create user
func createUser(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		w.Header().Set("Content-Type", "application/json")
		var newUser User //another user of type user
		reqBody, err := ioutil.ReadAll(r.Body)
		if err != nil {
			fmt.Fprintf(w, "Kindly enter data properly")
		}

		json.Unmarshal(reqBody, &newUser)
		users = append(users, newUser)

		w.WriteHeader(http.StatusCreated)

		json.NewEncoder(w).Encode(newUser)
		

	} else {
		http.Redirect(w, r, "/", http.StatusFound)
	}
}

//get an user, #####check mux
func getUser(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		w.Header().Set("Content-Type", "application/json")
		userID := strings.Split(r.URL.String()[len(`/users/`):], `/`)
		//userID := mux.Vars(r)["id"]
		uid, _ := strconv.Atoi(userID[0])
		for _, singleUser := range users {
			if singleUser.ID == uid {
				json.NewEncoder(w).Encode(singleUser)
			}
		}
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
		var newPost Post //another post of type post
		reqBody, err := ioutil.ReadAll(r.Body)
		if err != nil {
			fmt.Fprintf(w, "Kindly enter data properly")
		}

		json.Unmarshal(reqBody, &newPost)
		posts = append(posts, newPost)
		w.WriteHeader(http.StatusCreated)

		json.NewEncoder(w).Encode(newPost)

	} else {
		http.Redirect(w, r, "/", http.StatusFound)
	}
}

//get a post, #####check mux
func getOnePost(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		w.Header().Set("Content-Type", "application/json")
		postID := strings.Split(r.URL.String()[len(`/posts/`):], `/`)
		//r.URL.Path[len("/posts/{id}"):] --- wrong
		//postID := mux.Vars(r)["id"]
		pid, _ := strconv.Atoi(postID[0])

		for _, singlePost := range posts {
			if singlePost.EID == pid {
				json.NewEncoder(w).Encode(singlePost)
			}
		}
	} else {
		http.Redirect(w, r, "/", http.StatusFound)
	}

}

//get all posts
func getAllPosts(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		w.Header().Set("Content-Type", "application/json")
		userID := strings.Split(r.URL.String()[len(`/posts/users/`):], `/`)
		//r.URL.Path[len("/posts/{id}"):] --- wrong
		//postID := mux.Vars(r)["id"]
		upid, _ := strconv.Atoi(userID[0])

		for _, singleUser := range users {
			if singleUser.ID == upid {
				json.NewEncoder(w).Encode(posts)
			}
		}

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
