package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/medyagh/campwiz2/pkg/ramerica"
)

func handler(w http.ResponseWriter, r *http.Request) {
	res, err := ramerica.Search()
	if err != nil {
		log.Printf("failed to search reserve america. Resp: %v  error: %v", res, err)
	}

	// fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
	fmt.Fprintf(w, "\n%s\n", res.Body)

}

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
