package routes

import (
	"encoding/json"
	"net/http"
	"github.com/go-chi/chi/v5"
	comments "github.com/CATISNOTSODIUM/threadkeep-backend/internal/handlers/comments"
)


func CommentRoutes(r chi.Router) {
	r.Post("/comments/create", func(w http.ResponseWriter, req *http.Request) {
			response, httpCode := comments.HandleCreate(w, req)
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(httpCode)
			json.NewEncoder(w).Encode(response)
	})
	r.Post("/comments/update", func(w http.ResponseWriter, req *http.Request) {
			response, httpCode := comments.HandleUpdate(w, req)
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(httpCode)
			json.NewEncoder(w).Encode(response)
	})
	
	r.Post("/comments/delete", func(w http.ResponseWriter, req *http.Request) {
			response, httpCode := comments.HandleDelete(w, req)
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(httpCode)
			json.NewEncoder(w).Encode(response)
	})

	r.Post("/comments", func(w http.ResponseWriter, req *http.Request) {
			response, httpCode := comments.HandleRetrieve(w, req)
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(httpCode)
			json.NewEncoder(w).Encode(response)
	})
	
}