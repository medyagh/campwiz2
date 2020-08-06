package main

import (
	"log"
	"net/http"
	"text/template"

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
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
