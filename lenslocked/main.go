package main

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/jemajet/lenslocked/controllers"
	"github.com/jemajet/lenslocked/templates"
	"github.com/jemajet/lenslocked/views"
	"net/http"
)

func main() {
	// Create Router
	r := chi.NewRouter()

	// Setup Routing, parsing templates first
	tpl := views.Must(views.ParseFS(templates.FS, "home.gohtml", "layout-parts.gohtml"))
	r.Get("/", controllers.StaticHandler(tpl))

	tpl = views.Must(views.ParseFS(templates.FS, "contact.gohtml"))
	r.Get(
		"/contact",
		controllers.StaticHandler(
			views.Must(
				views.ParseFS(
					templates.FS,
					"contact.gohtml",
				),
			),
		),
	)

	tpl = views.Must(views.ParseFS(templates.FS, "faq.gohtml"))
	r.Get("/faq", controllers.FAQ(tpl))

	r.NotFound(
		func(w http.ResponseWriter, r *http.Request) {
			http.Error(w, "Page not found", http.StatusNotFound)
		},
	)

	// Start serving with our router
	fmt.Println("Starting the server on :3000...")
	http.ListenAndServe(":3000", r)
}
