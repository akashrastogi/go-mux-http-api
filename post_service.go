package main

import (
	"fmt"
    "log"
	"net/http"
	"github.com/gorilla/mux"
)

func main() {
	fmt.Printf("Starting server at port 8080\n")
	router := mux.NewRouter()

    if err := http.ListenAndServe(":8080", router); err != nil {
        log.Fatal(err)
    }
}