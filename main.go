package main

import (
	"github.com/gorilla/mux"
	"net/http"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/article", ArticleHandler).Methods("POST")
	router.HandleFunc("/article/{id:[0-9]+}", ArticleHandler).Methods("GET")

	router.HandleFunc("/", HomeHandler)

	http.Handle("/", router)
	http.ListenAndServe(":8080", nil)
}
