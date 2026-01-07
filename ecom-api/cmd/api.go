package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

type application struct {
	config config
}

// run
// mount
func (app *application) mount() http.Handler {
	r := chi.NewRouter()

	//Middlewares
	r.Use(middleware.Logger)
	r.Use(middleware.Recover)

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("root."))
	})

	// http.ListenAndServe(":3333", r)

	return r

}

type config struct {
	addr string
	db   dbConfig
}

type dbConfig struct {
	dsn string
}
