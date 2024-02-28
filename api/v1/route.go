package v1

import (
	"github.com/alaa-aqeel/school-ms-api-golang/api/v1/handler"
	"github.com/go-chi/chi/v5"
)

func Route() *chi.Mux {
	r := chi.NewRouter()
	r.Get("/", handler.Index)
	return r
}
