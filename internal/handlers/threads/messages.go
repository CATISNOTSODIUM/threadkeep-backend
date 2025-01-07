package Threads

// contains common message types

const (
	// Successful message
	SuccessfulRequest = "Succesfully Requested"
	SuccessfulListThreadsMessage = "Successfully listed Threads"
	SuccessfulRetrieveThreadsMessage = "Successfully retrieve thread in %s"
	SuccessfulCreateNewThreadMessage = "Successfully create new thread in %s"
	SuccessfulUpdateThreadMessage    = "Successfully update thread in %s"
	SuccessfulDeleteThread			 = "successfully delete thread in %s"
	ErrRetrieveDatabase        = "Failed to retrieve database in %s"
	
	// HTTP Request
	ErrInvalidPostRequest			 = "Method not allowed in %s"
	ErrBadRequest					 = "Bad request in %s"

	// Reaction
	ErrInvalidReactionType = "Invalid reaction type"
	ErrUpdateReaction = "Unable to update reaction. Possible duplicate."

)

