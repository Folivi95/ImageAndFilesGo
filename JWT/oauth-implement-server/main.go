package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	handleRequests()
}
func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "dummy secret information")
}

func handleRequests() {
	http.HandleFunc("/", homePage)
	log.Fatal(http.ListenAndServe(":9000", nil))
}
