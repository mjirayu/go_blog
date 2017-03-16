package main

import (
	"blog/controllers"
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	router.PathPrefix("/assets/").Handler(http.StripPrefix("/assets/", http.FileServer(http.Dir("assets"))))

	router.HandleFunc("/", controllers.Home).Methods("GET")
	router.HandleFunc("/posts", controllers.PostsHandler).Methods("GET")

	fmt.Println("Starting server on port :8080")
	http.ListenAndServe("localhost:8080", handlers.LoggingHandler(os.Stdout, router))
}
