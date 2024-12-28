package Comments


const (

	CreateNewComments = "Comments.HandleCreate"
	UpdateComment = "Comments.HandleUpdate"
	RetrieveComments = "Comments.HandleRetrieve"
	DeleteComment = "Comments.HandleDelete"
	SuccessfulCreateNewComments   = "Successfully create new comments in %s"
	SuccessfulRetrieveComments = "Successfully retrieve comments in %s"
	SuccessfulUpdateCommentMessage = "Successfully update comments in %s"
	SuccessfulDeleteComment = "Successfully delete comment in %s"
	ErrCreateComment		   = "Failed to create comments in %s"
	ErrRetrieveDatabase        = "Failed to retrieve database in %s"
	ErrRetrieveComments           = "Failed to retrieve Comments in %s"
	ErrDeleteComment = "Failed to delete comment in %s"
	ErrUpdateComment           = "Failed to update Comments in %s"
	ErrEncodeView              = "Failed to retrieve Comments in %s"
	ErrInvalidPostRequest			 = "Method not allowed in %s"
	ErrBadRequest					 = "Bad request in %s"
)

