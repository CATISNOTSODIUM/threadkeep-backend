package routes

import (
	"encoding/json"
	"net/http"
	handler "github.com/CATISNOTSODIUM/taggy-backend/internal/handlers"
	users "github.com/CATISNOTSODIUM/taggy-backend/internal/handlers/users"
	"github.com/go-chi/chi/v5"
)

// dummy data structure

func GetRoutes() func(r chi.Router) {
	return func(r chi.Router) {
		r.Get("/", func(w http.ResponseWriter, req *http.Request) {
			response, _ := handler.Index(w, req)
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(response)
		})
		r.Get("/users", func(w http.ResponseWriter, req *http.Request) {
			response, _ := users.HandleList(w, req)
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(response)
		})
		ThreadRoutes(r)
	}
}
