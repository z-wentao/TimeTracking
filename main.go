package main

import (
	"fmt"
	"net/http"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1> This website helps me tracking my time daily.")
}

func main() {
	http.HandleFunc("/", HomeHandler)
	fmt.Println("starting the server on 3000...")
	http.ListenAndServe(":3001", nil)
}
