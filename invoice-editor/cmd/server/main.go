package main

import (
	"invoice-editor/internal/api"
	"log"
	"net/http"
)

func main() {
	router := api.NewRouter()

	log.Println("Starting server on :8080")
	if err := http.ListenAndServe(":8080", router); err != nil {
		log.Fatal(err)
	}
}
