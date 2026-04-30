package main

import (
	"fmt"
	"log"
	"net/http"
)

func hellohandler(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	if name == "" {
		name = "World"
	}
	fmt.Fprintf(w, "hello %s!\n", name)
}

func main() {
	http.HandleFunc("/hello", hellohandler)
	log.Println("Server listening on :8000")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
