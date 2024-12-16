package users

type UserCreateRequest struct {
	Name string 		`json:"name"`
	Password string	`json:"password"`
}

type UserVerifyRequest struct {
	Name string 		`json:"name"`
	Password string		`json:"password"`
}