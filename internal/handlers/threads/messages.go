package Threads

// contains common message types

const (
	ListThreads = "Threads.HandleList"
	CountThreads = "Threads.HandleCount"
	ListThreadsTag = "Threads.HandleTag"
	RetrieveThread = "Threads.HandleRetrieve"
 	ReactThread = "Threads.HandleReactThread"
	DeleteThread = "Threads.DeleteThread"
	CreateNewThread = "Threads.HandleCreate"
	UpdateThread = "Threads.HandleUpdate"

	SuccessfulRequest = "Succesfully Requested"
	SuccessfulListThreadsMessage = "Successfully listed Threads"
	SuccessfulRetrieveThreadsMessage = "Successfully retrieve thread in %s"
	SuccessfulCreateNewThreadMessage = "Successfully create new thread in %s"
	SuccessfulUpdateThreadMessage    = "Successfully update thread in %s"
	SuccessfulDeleteThread			 = "successfully delete thread in %s"
	ErrRetrieveDatabase        = "Failed to retrieve database in %s"
	ErrRetrieveThreads           = "Failed to retrieve Threads in %s"
	ErrEncodeView              = "Failed to retrieve Threads in %s"
	ErrUpdateThreads		   = "[Failed to update Threads in %s] %s"
	ErrCreateThread	 				 = "Failed to create new thread in %s"
	ErrDeleteThread				= "Failed to delete thread in %s"
	
	
	// HTTP Request
	ErrInvalidPostRequest			 = "Method not allowed in %s"
	ErrBadRequest					 = "Bad request in %s"
	ErrParsingParams			= "Failed to parse URL params in %s"

	// Reaction
	ErrInvalidReactionType = "Invalid reaction type"
	ErrUpdateReaction = "Unable to update reaction. Possible duplicate."

)

