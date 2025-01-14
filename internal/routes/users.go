package routes

import (
	"encoding/json"
	"net/http"
	"github.com/go-chi/chi/v5"
	users "github.com/CATISNOTSODIUM/threadkeep-backend/internal/handlers/users"
)

func UserRoutes(r chi.Router) {
	r.Get("/users", func(w http.ResponseWriter, req *http.Request) {
		response, httpCode := users.HandleList(w, req)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(httpCode)
		json.NewEncoder(w).Encode(response)
	})
	r.Post("/users/create", func(w http.ResponseWriter, req *http.Request) {
		response, httpCode := users.HandleCreate(w, req)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(httpCode)
		json.NewEncoder(w).Encode(response)
	})
	r.Post("/users/verify", func(w http.ResponseWriter, req *http.Request) {
		response, httpCode := users.HandleVerify(w, req)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(httpCode)
		json.NewEncoder(w).Encode(response)
	})
}