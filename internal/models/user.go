package models


type User struct {
	ID   string    `json:"id"`
	Name string `json:"name"`
	JWTToken string `json:"jwt_token"`
}


