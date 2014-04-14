package main

import (
	"encoding/json"
	"net/http"
	"time"
)

type Response struct {
	Status string
	Body   string
	Time   int64
}

func makeResponseString(status string, body string) string {
	response := Response{status, body, time.Now().Unix()}

	jsonResponse, err := json.Marshal(response)
	if err != nil {
		panic(err)
	}

	return string(jsonResponse)
}

func getRequestPayload(r *http.Request) string {
	payload := r.PostFormValue("payload")
	if payload == "" {
		panic("No Input :\\")
	}

	return payload
}
