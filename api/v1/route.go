package v1

import (
	"github.com/go-chi/chi/v5"
)

func Route() *chi.Mux {
	r := chi.NewRouter()

	return r
}
