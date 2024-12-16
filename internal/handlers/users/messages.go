package users

const (
	CreateUser = "users.HandleCreate"
	ListUsers  = "users.HandleList"
	VerifyUser = "users.HandleVerify"
	SuccessfulListUsersMessage = "Successfully listed users"
	SuccessfulCreateUser       = "Successfully create new user in %s"
	SuccessfulUpdateUser       = "Successfully update user in %s"
	SuccessfulVerifyUser	   = "Successfully verify user"
	ErrRetrieveUsers      = "Failed to retrieve users in %s"
	ErrRetrieveDatabase   = "Failed to retrieve database in %s"
	ErrRetrieveUser       = "Failed to retrieve users in %s"
	ErrEncodeView         = "Failed to retrieve users in %s"
	ErrUpdateUsers        = "Failed to update users in %s"
	ErrCreateUser         = "Failed to create new user in %s"
	ErrInvalidPostRequest = "Method not allowed in %s"
	ErrBadRequest         = "Bad request in %s"
	ErrVerifyUser 		  = "Failed to verify user in %s"
)
