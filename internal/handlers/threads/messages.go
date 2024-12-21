package Threads


const (
	ListThreads = "Threads.HandleList"
	ListThreadsTag = "Threads.HandleTag"
	RetrieveThread = "Threads.HandleRetrieve"
 	ReactThread = "Threads.HandleReactThread"
	CreateNewThread = "Threads.HandleCreate"
	UpdateThread = "Threads.HandleUpdate"
	SuccessfulListThreadsMessage = "Successfully listed Threads"
	SuccessfulRetrieveThreadsMessage = "Successfully retrieve thread in %s"
	SuccessfulCreateNewThreadMessage = "Successfully create new thread in %s"
	
	SuccessfulUpdateThreadMessage    = "Successfully update thread in %s"

	ErrRetrieveDatabase        = "Failed to retrieve database in %s"
	ErrRetrieveThreads           = "Failed to retrieve Threads in %s"
	ErrEncodeView              = "Failed to retrieve Threads in %s"
	ErrUpdateThreads		   = "Failed to update Threads in %s"
	ErrCreateThread	 				 = "Failed to create new thread in %s"
	ErrInvalidPostRequest			 = "Method not allowed in %s"
	ErrBadRequest					 = "Bad request in %s"

)

