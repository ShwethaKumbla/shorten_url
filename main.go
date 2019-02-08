package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/shorten_url/controllers"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/{endpoint}", controllers.Redirect).Methods("GET")
	router.HandleFunc("/list/", controllers.Retrieve).Methods("GET")
	router.HandleFunc("/create", controllers.Create).Methods("POST")
	router.HandleFunc("/delete", controllers.Delete).Methods("DELETE")
	router.HandleFunc("/update", controllers.Update).Methods("PUT")

	log.Println("listening on :8090")
	log.Fatal(http.ListenAndServe(":8090", router))
}
