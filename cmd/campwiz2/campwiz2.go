package main

import (
	"log"
	"net/http"
	"strings"
	"text/template"

	"github.com/medyagh/campwiz2/pkg/book"
	"github.com/medyagh/campwiz2/pkg/ramerica"
)

func handler(w http.ResponseWriter, r *http.Request) {
	results, err := ramerica.Search()
	if err != nil {
		log.Printf("failed to search reserve america. resp: %v  error: %v", results, err)
	}

	applyBookRating(results)

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

func applyBookRating(rs []*ramerica.Record) {
	for _, r := range rs {
		for k, v := range book.LoadedEntries {
			if strings.Contains(k, r.Name) {
				log.Printf("\n ------ found book record %q for %q: %+v ------------------ \n ", r.Name, k, r)
				r.BookRecord = v
			}
		}
	}
}
