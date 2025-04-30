package main

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	tpl, err := template.ParseFiles("templates/home.gohtml")
	if err != nil {
		panic(err)
	}
	err = tpl.Execute(w, nil)
	if err != nil {
		panic(err)
	}
}

func LogHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>April 30 Time Log</h1>")
	fmt.Fprintf(w, "8.00AM Morning Routine<br>")
	fmt.Fprintf(w, "9.00AM Interleaving\n")

	tpl, err := template.ParseFiles("templates/log.gohtml")
	if err != nil {
		panic(err)
	}
	err = tpl.Execute(w, nil)
	if err != nil {
		panic(err)
	}
}

func ReflectionHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Self-Reflection based on daily Log</h1>")
	w.Header().Set("Content-Type", "text/hmtl; charset=utf-8")
	fmt.Fprintf(w, "It works well")

	tpl, err := template.ParseFiles("templates/reflection.gohtml")
	if err != nil {
		panic(err)
	}
	err = tpl.Execute(w, nil)
	if err != nil {
		panic(err)
	}
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
