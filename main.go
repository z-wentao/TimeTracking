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
	tpl, err := views.Parse("templates/home.gohtml")
	if err != nil {
		panic(err)
	}
	r.Get("/", controllers.StaticHandler(tpl))

	tpl, err = views.Parse("templates/log.gohtml")
	if err != nil {
		panic(err)
	}
	r.Get("/log", controllers.StaticHandler(tpl))

	tpl, err = views.Parse("templates/reflection.gohtml")
	if err != nil {
		panic(err)
	}
	r.Get("/reflection", controllers.StaticHandler(tpl))

	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "oh, we cannot found this page", http.StatusNotFound)
	})
	fmt.Println("starting the server on 3001...")
	http.ListenAndServe(":3001", r)
}
