package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1> This website helps me tracking my time and self-reflection daily.")
}

func LogHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>April 30 Time Log</h1>")
	fmt.Fprintf(w, "8.00AM Morning Routine<br>")
	fmt.Fprintf(w, "9.00AM Interleaving\n")
}

func ReflectionHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Self-Reflection based on daily Log</h1>")
	w.Header().Set("Content-Type", "text/hmtl; charset=utf-8")
	fmt.Fprintf(w, "It works well")
}

type Router struct{}

func main() {
	r := chi.NewRouter()
	r.Get("/", HomeHandler)
	r.Get("/log", LogHandler)
	r.Get("/reflection", ReflectionHandler)
	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "oh, we cannot found this page", http.StatusNotFound)
	})
	fmt.Println("starting the server on 3001...")
	http.ListenAndServe(":3001", r)
}
