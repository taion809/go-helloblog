package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	response := makeResponseString("Success", "Hello Blog, let's learn something!")
	fmt.Fprintf(w, response)
}

func ArticleHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		payload := getRequestPayload(r)

		id := writeToFile(payload)
		response := makeResponseString("Success", id)

		fmt.Fprintf(w, response)
		return
	}

	vars := mux.Vars(r)

	response := makeResponseString("Success", readFile(vars["id"]))
	fmt.Fprintf(w, response)
}
