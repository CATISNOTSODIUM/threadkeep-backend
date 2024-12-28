package Comments

import (
	"github.com/CATISNOTSODIUM/taggy-backend/internal/models"
)

type CommentCreateRequest struct {
    Content     string 		`json:"content"`
	User		models.User `json:"user"`
	ThreadID	string
}

type CommentUpdateRequest struct {
	CommentID	string		`json:"commentID"`
    Content     string 		`json:"content"`
	User		models.User `json:"user"`
}

type CommentRetrieveRequest struct {
	ThreadID	string `json:"threadID"`
}

type CommentDeleteRequest struct {
	CommentID string `json:"commentID"`
}