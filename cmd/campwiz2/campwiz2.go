package main

import (
	"log"
	"net/http"
	"strings"
	"text/template"

	"github.com/medyagh/campwiz2/pkg/book"
	"github.com/medyagh/campwiz2/pkg/ramerica"
)

// func getHandler(w http.ResponseWriter, r *http.Request) {

// }

func handler(w http.ResponseWriter, r *http.Request) {
	c := ramerica.Criteria{
		Longitude:    -122.07237049999999,
		Latitude:     37.4092297,
		ArrivalDate:  "2020-08-14",
		LengthOfStay: 2,
	}
	results, err := ramerica.Search(c)
	if err != nil {
		log.Printf("failed to search reserve america. resp: %v  error: %v", results, err)
	}

	applyBookRating(results)
	results = removeUnwanted(results)
	tmpl, err := template.New("campwiz").Parse(ResultsTemplate)
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
				log.Printf("found book record %q for %q ", r.Name, k)
				r.BookRecord = v
			}
		}
	}
}

func removeUnwanted(rs []*ramerica.Record) []*ramerica.Record {
	var frs []*ramerica.Record
	for _, r := range rs {
		if !r.Details.Availability.Available && r.Details.VerifiableAvailability {
			continue
		}
		frs = append(frs, r)
	}
	return frs
}
