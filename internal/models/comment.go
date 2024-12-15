package models

import ("time")

type Comment struct {
	ID   		string  `json:"id"`
    Content     string 	`json:"content"`
    Tags        []Tag   `json:"tags"`
    Likes       int     `json:"likes"`
    Views       int     `json:"views"`
	User		User    `json:"user"`
	Parent		Thread	`json:"parent"`
    CreatedAt time.Time `json:"createdAt"`
    UpdatedAt time.Time `json:"updatedAt"`
}


