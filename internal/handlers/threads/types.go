package Threads

import (
	"github.com/CATISNOTSODIUM/threadkeep-backend/internal/models"
)

type ThreadCreateRequest struct {
    Title       string 		`json:"title"`
    Content     string 		`json:"content"`
	User		models.User `json:"user"`
	Tags		[]models.Tag `json:"tags"`
}

type ThreadUpdateRequest struct {
	ThreadID    string 		`json:"threadID"`
    Title       string 		`json:"title"`
    Content     string 		`json:"content"`
	User		models.User `json:"user"`
}

type ThreadDeleteRequest struct {
	ThreadID    string 		`json:"threadID"`
}

type ThreadRetrieveRequest struct {
	ThreadID       string 	`json:"threadID"`
}

type ThreadLikeRequest struct {
	UserID		   string 	`json:"userID"`
    ThreadID       string 	`json:"threadID"`
	Reaction	   ReactionType `json:"reaction"`
}

type ThreadIsLikeRequest struct {
	UserID		   string 	`json:"userID"`
    ThreadID       string 	`json:"threadID"`
}


type ReactionType int32
const (
	VIEW 	ReactionType = 0
	LIKE	ReactionType = 1
	UNLIKE	ReactionType = 2
	// save thread
	SAVED	ReactionType = 3 
	UNSAVE	ReactionType = 4
)

