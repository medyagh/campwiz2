package main

import (
	"log"
	"net/http"
	"strconv"
	"strings"
	"text/template"
	"time"

	"github.com/medyagh/campwiz2/pkg/book"
	"github.com/medyagh/campwiz2/pkg/ramerica"
)

// indexHandler handles the index page for both GET and POST requests
func indexHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		formHandler(w, r)
	}
	if r.Method == "POST" {
		resultsHandler(w, r)
	}
}

func formHandler(w http.ResponseWriter, r *http.Request) {
	type FormPage struct {
		Today string
	}
	tmpl, err := template.New("searchForm").Parse(searchFormTemplate)
	if err != nil {
		log.Fatalf("failed to parse template")
	}

	err = tmpl.Execute(w, &FormPage{
		Today: time.Now().Format("2006-01-02"),
	})
	if err != nil {
		panic(err)
	}
}

// handler that shows the search results in a table
func resultsHandler(w http.ResponseWriter, r *http.Request) {
	var startDate = time.Now()
	var lenOfStay, maxDistance int

	err := r.ParseForm()
	if err != nil {
		log.Printf("error parsing the form: %v", err)
	}
	if r.Form["nights"] != nil {
		lenOfStay, err = strconv.Atoi(r.Form["nights"][0])
		if err != nil {
			log.Printf("Error parsing night %q : %v", r.Form["nights"][0], err)
		}
	}

	if r.Form["dates"] != nil {
		startDate, err = time.Parse("2006-01-02", r.Form["dates"][0])
		if err != nil {
			log.Printf("Error parsing date %q : %v", r.Form["dates"][0], err)
		}

	}
	if r.Form["distance"] != nil {
		maxDistance, err = strconv.Atoi(r.Form["distance"][0])
		if err != nil {
			log.Printf("Error parsing distance %q : %v", r.Form["distance"][0], err)
		}
	}
	log.Println("todo later user MaxDistance", maxDistance)

	c := ramerica.Criteria{
		Longitude:    -122.07237049999999,
		Latitude:     37.4092297,
		ArrivalDate:  startDate.Format("2006-01-02"),
		LengthOfStay: lenOfStay,
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
