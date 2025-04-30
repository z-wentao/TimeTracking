package main

import (
	"fmt"
	"net/http"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1> This website helps me tracking my time and self-reflection daily.")
}

func LogHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>April 30 Time Log</h1>")
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprintf(w, "8.00AM Morning Routine")
}

func ReflectionHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Self-Reflection based on daily Log</h1>")
	w.Header().Set("Content-Type", "text/hmtl; charset=utf-8")
	fmt.Fprintf(w, "It works well")
}

type Router struct{}

func (router Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/":
		HomeHandler(w, r)
	case "/log":
		LogHandler(w, r)
	case "/reflection":
		ReflectionHandler(w, r)
	default:
		http.Error(w, "oh, we cannot found this page", http.StatusNotFound)
	}
}

func main() {
	var router Router
	fmt.Println("starting the server on 3001...")
	http.ListenAndServe(":3001", router)
}
