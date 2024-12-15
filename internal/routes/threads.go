package routes

import (
	"encoding/json"
	"net/http"
	threads "github.com/CATISNOTSODIUM/taggy-backend/internal/handlers/threads"
	"github.com/go-chi/chi/v5"
)


func ThreadRoutes(r chi.Router) {
	// list all threads
	r.Get("/threads", func(w http.ResponseWriter, req *http.Request) {
			response, _ := threads.HandleList(w, req)
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(response)
	})
	// get individual thread
	r.Post("/threads", func(w http.ResponseWriter, req *http.Request) {
		response, _ := threads.HandleRetrieve(w, req)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	})
	r.Post("/threads/create", func(w http.ResponseWriter, req *http.Request) {
		response, _ := threads.HandleCreate(w, req)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	})
	r.Post("/threads/update", func(w http.ResponseWriter, req *http.Request) {
		response, _ := threads.HandleUpdate(w, req)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	})
	r.Post("/threads/reaction", func(w http.ResponseWriter, req *http.Request) {
		response, _ := threads.HandleLikeThread(w, req)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	})
}