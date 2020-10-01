package main

import (
	"log"
	"net/http"

	"github.com/medyagh/campwiz2/pkg/book"
)

func main() {
	err := book.Load()
	if err != nil {
		log.Printf("Err Loading the book: %v", err)
	}

	http.HandleFunc("/", indexHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
