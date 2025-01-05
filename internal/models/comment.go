package models

import (
	"time"
)

type Comment struct {
	ID   		string  `json:"id"`
    Content     string 	`json:"content"`
    Tags        []Tag   `json:"tags"`
    Likes       int     `json:"likes"`
    Views       int     `json:"views"`
	User		User  `json:"user"`
	ParentID	string	`json:"parentID"`
    CreatedAt time.Time `json:"createdAt"`
    UpdatedAt time.Time `json:"updatedAt"`
}


