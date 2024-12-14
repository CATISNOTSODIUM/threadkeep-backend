package users

import (
	"net/http"

	"github.com/CATISNOTSODIUM/taggy-backend/internal/api"
)

func Index(w http.ResponseWriter, r *http.Request) (*api.Response, error) {
	return &api.Response{
		Messages: []string{"Welcome to our taggy api server!"},
	}, nil
}
