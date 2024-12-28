package Comments

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/pkg/errors"

	"github.com/CATISNOTSODIUM/taggy-backend/internal/api"
	"github.com/CATISNOTSODIUM/taggy-backend/internal/dataaccess/mutation"
	"github.com/CATISNOTSODIUM/taggy-backend/internal/database"
)


func HandleUpdate(w http.ResponseWriter, r *http.Request) (*api.Response, error) {
	if r.Method != http.MethodPost {
		errorMessage := fmt.Sprintf(ErrInvalidPostRequest, UpdateComment)
		http.Error(w, errorMessage, http.StatusMethodNotAllowed)
		return nil, errors.New(errorMessage)
	}

	comment := &CommentUpdateRequest{}
	err := json.NewDecoder(r.Body).Decode(comment)

	if err != nil {
		errorMessage := fmt.Sprintf(ErrBadRequest, UpdateComment)
		http.Error(w, errorMessage, http.StatusBadRequest)
		return nil, errors.Wrap(err, errorMessage)
	}


	db, err := database.Connect()
	if err != nil {
		errorMessage := fmt.Sprintf(ErrRetrieveDatabase, UpdateComment)
		http.Error(w, errorMessage, http.StatusBadRequest)
		return nil, errors.Wrap(err, errorMessage)
	}

	defer db.Close()

	commentObject, err := mutation.UpdateComment(db, comment.CommentID, comment.Content)
	
	if err != nil {
		errorMessage := fmt.Sprintf(ErrUpdateComment, UpdateComment)
		http.Error(w, errorMessage, http.StatusBadRequest)
		return nil, errors.Wrap(err, errorMessage)
	}
	

	data, err := json.Marshal(commentObject)
	if err != nil {
		errorMessage := fmt.Sprintf(ErrEncodeView, UpdateComment)
		return nil, errors.Wrap(err, errorMessage)
	}

	return &api.Response{
		Payload: api.Payload{
			Data: data,
		},
		Messages: []string{fmt.Sprintf(SuccessfulUpdateCommentMessage, UpdateComment)},
	}, nil
	
}