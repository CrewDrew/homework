package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/items/*", items)
	log.Fatal(http.ListenAndServe("localhost:8888", nil))
}

func items(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:

	case http.MethodPost:

	case http.MethodDelete:
		// Remove the record.
	default:
		// Give an error message.
	}
}
