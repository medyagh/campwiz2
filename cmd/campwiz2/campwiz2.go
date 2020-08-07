package main

import (
	"log"
	"net/http"
	"text/template"

	"github.com/medyagh/campwiz2/pkg/book"
	"github.com/medyagh/campwiz2/pkg/ramerica"
)

func handler(w http.ResponseWriter, r *http.Request) {
	results, err := ramerica.Search()
	if err != nil {
		log.Printf("failed to search reserve america. resp: %v  error: %v", results, err)
	}

	tmpl, err := template.New("test").Parse(ResultsTemplate)
	if err != nil {
		log.Fatalf("failed to parse template")
	}

	err = tmpl.Execute(w, results)
	if err != nil {
		panic(err)
	}

}

func main() {
	err := book.Load()
	if err != nil {
		log.Printf("Err Loading the book: %v", err)
	}

	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
