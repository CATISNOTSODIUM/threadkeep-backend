package models

import (
	"time"
)

type Thread struct {
    ID          string `json:"id"`
    Title       string `json:"title"`
    Content     string `json:"content"`
    Tags        []Tag   `json:"tags"`
	
    Likes       int    `json:"likes"`
    Views       int    `json:"views"`

	User		User    `json:"user"`

	CreatedAt time.Time `json:"createdAt"`
    UpdatedAt time.Time `json:"updatedAt"`
    
    IsSaved     bool    `json:"isSaved"`
}
