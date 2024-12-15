package Comments

import (
	"github.com/CATISNOTSODIUM/taggy-backend/internal/models"
)

type CommentCreateRequest struct {
    Content     string 		`json:"content"`
	User		models.User `json:"user"`
	ThreadID	string
}

type CommentRetrieveRequest struct {
	ThreadID	string
}