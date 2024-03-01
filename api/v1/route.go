package v1

import (
	"github.com/alaa-aqeel/school-ms-api-golang/api/middlewares"
	"github.com/alaa-aqeel/school-ms-api-golang/api/v1/handler/auth"
	"github.com/go-chi/chi/v5"
)

func Route() *chi.Mux {
	r := chi.NewRouter()
	r.Use(middlewares.AuthMiddleware)
	r.Post("/login", auth.LoginHandler)
	return r
}
