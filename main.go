package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/z-wentao/TimeTracking/controllers"
	"github.com/z-wentao/TimeTracking/templates"
	"github.com/z-wentao/TimeTracking/views"
)

type Router struct{}

func main() {
	r := chi.NewRouter()
	r.Get("/", controllers.StaticHandler(
		views.Must(views.ParseFS(templates.FS, "home.gohtml"))))

	r.Get("/log", controllers.StaticHandler(
		views.Must(views.ParseFS(templates.FS, "log.gohtml"))))

	r.Get("/reflection", controllers.StaticHandler(
		views.Must(views.ParseFS(templates.FS, "reflection.gohtml"))))

	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "oh, we cannot found this page", http.StatusNotFound)
	})
	fmt.Println("starting the server on 3001...")
	http.ListenAndServe(":3001", r)
}
