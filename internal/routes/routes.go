package routes

import (
	"encoding/json"
	"net/http"
	handler "github.com/CATISNOTSODIUM/threadkeep-backend/internal/handlers"
	"github.com/go-chi/chi/v5"
)

func GetRoutes() func(r chi.Router) {
	return func(r chi.Router) {
		r.Get("/", func(w http.ResponseWriter, req *http.Request) {
			response, _ := handler.Index(w, req)
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(response)
		})
		UserRoutes(r)
		ThreadRoutes(r)
		CommentRoutes(r)
	}
}
