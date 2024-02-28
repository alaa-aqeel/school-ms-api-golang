package main

import (
	"net/http"

	v1 "github.com/alaa-aqeel/school-ms-api-golang/api/v1"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Mount("/api/v1", v1.Route())
	http.ListenAndServe(":8080", r)
}
