package Threads

import (
	"github.com/CATISNOTSODIUM/taggy-backend/internal/models"
	
)

type ThreadCreateRequest struct {
    Title       string 		`json:"title"`
    Content     string 		`json:"content"`
	User		models.User `json:"user"`
}

type ThreadUpdateRequest struct {
	ThreadID    string 		`json:"threadID"`
    Title       string 		`json:"title"`
    Content     string 		`json:"content"`
	User		models.User `json:"user"`
}


type ThreadRetrieveRequest struct {
	ThreadID       string 	`json:"threadID"`
}

type ThreadLikeRequest struct {
    ThreadID       string 	`json:"threadID"`
	Reaction	   ReactionType `json:"reaction"`
}

type ReactionType int32
const (
	LIKE	ReactionType = 1
	UNLIKE	ReactionType = 2
	VIEW 	ReactionType = 0
)

