package main

import (
	"fmt"
	"log"
	"net/http"
	"encoding/json"
	"github.com/gorilla/mux"
	"strconv"
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
	router.HandleFunc("/posts/{id}", getPost).Methods("GET")
	router.HandleFunc("/post/{id}", deletePost).Methods("DELETE")

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

func getAllPosts(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(posts)
}

func getPost(w http.ResponseWriter, req *http.Request) {
	var idParam string = mux.Vars(req)["id"]
	id, err := strconv.Atoi(idParam)
	if err != nil {
		w.WriteHeader(400)
		w.Write([]byte("No post exist for the specified Id"))
	}
	if id >= len(posts) {
		w.WriteHeader(404)
		w.Write([]byte("No post found with specified ID"))
		return
	}
	post := posts[id]
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(post)
}

func deletePost(w http.ResponseWriter, req *http.Request) {
	var idParam string = mux.Vars(req)["id"]
	id, err := strconv.Atoi(idParam)
	if err != nil {
		w.WriteHeader(400)
		w.Write([]byte("Invalid Id"))
		return
	}

	if id >= len(posts) {
		w.WriteHeader(404)
		w.Write([]byte("No post found with specified ID"))
		return
	}

	// Delete the post from the slice
	// https://github.com/golang/go/wiki/SliceTricks#delete
	posts = append(posts[:id], posts[id+1:]...)
	w.WriteHeader(200)
}