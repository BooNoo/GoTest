package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	log.Println("Starting server on :8080")
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s from %s", r.Method, r.URL.Path, r.RemoteAddr)
		fmt.Fprintln(w, "Hello world")
	})

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
