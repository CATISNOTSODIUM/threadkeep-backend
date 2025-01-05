package Comments

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/CATISNOTSODIUM/threadkeep-backend/internal/api"
	"github.com/CATISNOTSODIUM/threadkeep-backend/internal/dataaccess/mutation"
	"github.com/CATISNOTSODIUM/threadkeep-backend/internal/database"
	"github.com/pkg/errors"
)



func HandleCreate(w http.ResponseWriter, r *http.Request) (*api.Response, error) {
	if r.Method != http.MethodPost {
		errorMessage := fmt.Sprintf(ErrInvalidPostRequest, CreateNewComments)
		http.Error(w, errorMessage, http.StatusMethodNotAllowed)
		return nil, errors.New(errorMessage)
	}

	comment := &CommentCreateRequest{}
	err := json.NewDecoder(r.Body).Decode(comment)

	if err != nil {
		errorMessage := fmt.Sprintf(ErrBadRequest, CreateNewComments)
		http.Error(w, errorMessage, http.StatusBadRequest)
		return nil, errors.Wrap(err, errorMessage)
	}

	db, err := database.Connect()
	if err != nil {
		errorMessage := fmt.Sprintf(ErrRetrieveDatabase, CreateNewComments)
		http.Error(w, errorMessage, http.StatusBadRequest)
		return nil, errors.Wrap(err, errorMessage)
	}

	defer db.Close()

	commentObject, err := mutation.CreateComment(db, comment.User.ID, comment.ThreadID, comment.Content)
	
	if err != nil {
		errorMessage := fmt.Sprintf(ErrCreateComment, CreateNewComments)
		http.Error(w, errorMessage, http.StatusBadRequest)
		return nil, errors.Wrap(err, errorMessage)
	}
	
	data, err := json.Marshal(commentObject)
	if err != nil {
		errorMessage := fmt.Sprintf(ErrEncodeView, CreateNewComments)
		return nil, errors.Wrap(err, errorMessage)
	}

	return &api.Response{
		Payload: api.Payload{
			Data: data,
		},
		Messages: []string{SuccessfulCreateNewComments},
	}, nil
}

