package router

import (
	"github.com/CATISNOTSODIUM/threadkeep-backend/internal/routes"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
)

func Setup() chi.Router {
	r := chi.NewRouter()
	r.Use(cors.Handler(cors.Options{
			// AllowedOrigins:   []string{"https://foo.com"}, // Use this to allow specific origin hosts
			AllowedOrigins:   []string{"http://localhost:3000"},
			// AllowOriginFunc:  func(r *http.Request, origin string) bool { return true },
			AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
			AllowedHeaders:   []string{"Accept", "Accept-Language", "Authorization", "Content-Type"},
			AllowCredentials: true,
			MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))
	setUpRoutes(r)
	return r
}

func setUpRoutes(r chi.Router) {
	publicRoutes := routes.GetRoutes()
	protectedRoutes := routes.GetProtectedRoutes()
	r.Group(publicRoutes)
	r.Group(protectedRoutes)
}
