package users

import (
	"net/http"

	"github.com/CATISNOTSODIUM/threadkeep-backend/internal/api"
)

func Index(w http.ResponseWriter, r *http.Request) (*api.Response, error) {
	return &api.Response{
		Messages: []string{"Welcome to our threadkeep api server v2."},
	}, nil
}
