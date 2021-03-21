package main

import (
	"fmt"
	"log"
	"net/http"
	"encoding/json"
	"github.com/gorilla/mux"
)

// Model to define Author of the Post
type User struct {
	Name string `json:"name"`
	UserName string `json:"userName"`
	Email string `json:"email"`
}

// Post model to store data
type Post struct {
	Title  string `json:"title"`
	Body   string `json:"body"`
	Author User   `json:"author"`
}

var posts []Post = []Post {}

func main() {
	fmt.Printf("Starting server at port 8080\n")
	router := mux.NewRouter()

	router.HandleFunc("/post", addPost).Methods("POST")
	router.HandleFunc("/posts", getAllPosts).Methods("GET")

    if err := http.ListenAndServe(":8080", router); err != nil {
        log.Fatal(err)
    }
}

func addPost(w http.ResponseWriter, req *http.Request) {
	var newPost Post
	json.NewDecoder(req.Body).Decode(&newPost)
	posts = append(posts, newPost)
	w.Header().Set("Content-Type", "applicaton/json")
	json.NewEncoder(w).Encode(posts)
}

func getAllPosts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(posts)
}