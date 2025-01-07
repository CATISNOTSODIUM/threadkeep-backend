package routes

import (
	"encoding/json"
	"net/http"
	threads "github.com/CATISNOTSODIUM/threadkeep-backend/internal/handlers/threads"
	"github.com/go-chi/chi/v5"
)


func ThreadRoutes(r chi.Router) {
	// list all threads
	r.Get("/threads", func(w http.ResponseWriter, req *http.Request) {
			response, _ := threads.HandleList(w, req)
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(response)
	})
	r.Get("/threads/count", func(w http.ResponseWriter, req *http.Request) {
			response, _ := threads.HandleCount(w, req)
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
	r.Post("/threads/delete", func(w http.ResponseWriter, req *http.Request) {
		response, _ := threads.HandleDelete(w, req)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	})
	r.Post("/threads/tags", func(w http.ResponseWriter, req *http.Request) {
		response, _ := threads.HandleTag(w, req)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	})
	r.Get("/threads/tags", func(w http.ResponseWriter, req *http.Request) {
		response, _ := threads.HandleTagList(w, req)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	})
	r.Post("/threads/reaction", func(w http.ResponseWriter, req *http.Request) {
		response, _ := threads.HandleReactThread(w, req)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	})
	r.Post("/threads/reaction/isLike", func(w http.ResponseWriter, req *http.Request) {
		response, _ := threads.HandleIsLikeThread(w, req)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	})
}