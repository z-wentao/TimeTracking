package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/z-wentao/TimeTracking/controllers"
	"github.com/z-wentao/TimeTracking/views"
)

type Router struct{}

func main() {
	r := chi.NewRouter()
	tpl := views.Must(views.Parse("templates/home.gohtml"))
	r.Get("/", controllers.StaticHandler(tpl))

	tpl = views.Must(views.Parse("templates/log.gohtml"))
	r.Get("/log", controllers.StaticHandler(tpl))

	tpl = views.Must(views.Parse("templates/reflection.gohtml"))
	r.Get("/reflection", controllers.StaticHandler(tpl))

	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "oh, we cannot found this page", http.StatusNotFound)
	})
	fmt.Println("starting the server on 3001...")
	http.ListenAndServe(":3001", r)
}
