package router

import (
	"github.com/CATISNOTSODIUM/taggy-backend/internal/routes"
	"github.com/go-chi/chi/v5"
)

func Setup() chi.Router {
	r := chi.NewRouter()
	setUpRoutes(r)
	return r
}

func setUpRoutes(r chi.Router) {
	r.Group(routes.GetRoutes())
}
