package routes

import (
	"github.com/CATISNOTSODIUM/threadkeep-backend/internal/middleware"
	"encoding/json"
	"net/http"
	handler "github.com/CATISNOTSODIUM/threadkeep-backend/internal/handlers"
	"github.com/go-chi/chi/v5"
)



func GetRoutes() func(r chi.Router) {
	return func(r chi.Router) {
		r.Get("/", func(w http.ResponseWriter, req *http.Request) {
			response, httpCode := handler.Index(w, req)
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(httpCode)
			json.NewEncoder(w).Encode(response)
		})
		UserRoutes(r)
	}
}

func GetProtectedRoutes() func(r chi.Router) {
	return func(r chi.Router) {
		r.Use(middleware.JWT_Middleware()) // user authentication
		ThreadRoutes(r)
		CommentRoutes(r)
	}
}