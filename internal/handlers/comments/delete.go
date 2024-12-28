package Comments

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/CATISNOTSODIUM/taggy-backend/internal/api"
	"github.com/CATISNOTSODIUM/taggy-backend/internal/dataaccess/mutation"
	"github.com/CATISNOTSODIUM/taggy-backend/internal/database"
	"github.com/pkg/errors"
)


func HandleDelete(w http.ResponseWriter, r *http.Request) (*api.Response, error) {
	if r.Method != http.MethodPost {
		errorMessage := fmt.Sprintf(ErrInvalidPostRequest, DeleteComment)
		http.Error(w, errorMessage, http.StatusMethodNotAllowed)
		return nil, errors.New(errorMessage)
	}

	comment := &CommentDeleteRequest{}
	err := json.NewDecoder(r.Body).Decode(comment)

	if err != nil {
		errorMessage := fmt.Sprintf(ErrBadRequest, DeleteComment)
		http.Error(w, errorMessage, http.StatusBadRequest)
		return nil, errors.Wrap(err, errorMessage)
	}


	db, err := database.Connect()
	if err != nil {
		errorMessage := fmt.Sprintf(ErrRetrieveDatabase, DeleteComment)
		http.Error(w, errorMessage, http.StatusBadRequest)
		return nil, errors.Wrap(err, errorMessage)
	}

	defer db.Close()

	threadObject, err := mutation.DeleteComment(db, comment.CommentID)
	
	if err != nil {
		errorMessage := fmt.Sprintf(ErrDeleteComment, DeleteComment)
		http.Error(w, errorMessage, http.StatusBadRequest)
		return nil, errors.Wrap(err, errorMessage)
	}
	

	data, err := json.Marshal(threadObject)
	if err != nil {
		errorMessage := fmt.Sprintf(ErrEncodeView, DeleteComment)
		return nil, errors.Wrap(err, errorMessage)
	}

	return &api.Response{
		Payload: api.Payload{
			Data: data,
		},
		Messages: []string{fmt.Sprintf(SuccessfulDeleteComment, DeleteComment)},
	}, nil
	
}