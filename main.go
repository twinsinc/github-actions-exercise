package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", root)
	log.Fatalf("error: %s", http.ListenAndServe(":8080", nil))
}
func root(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World.")
}
