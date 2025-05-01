package main

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func ParseAndExecute(w http.ResponseWriter, filepath string) {
	tpl, err := template.ParseFiles(filepath)
	if err != nil {
		panic(err)
	}
	err = tpl.Execute(w, nil)
	if err != nil {
		panic(err)
	}
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	ParseAndExecute(w, "templates/home.gohtml")
}

func LogHandler(w http.ResponseWriter, r *http.Request) {
	ParseAndExecute(w, "templates/log.gohtml")
}

func ReflectionHandler(w http.ResponseWriter, r *http.Request) {
	ParseAndExecute(w, "templates/reflection.gohtml")
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
