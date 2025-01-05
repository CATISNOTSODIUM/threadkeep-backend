package Comments

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/CATISNOTSODIUM/threadkeep-backend/internal/api"
	"github.com/CATISNOTSODIUM/threadkeep-backend/internal/dataaccess/query"
	"github.com/CATISNOTSODIUM/threadkeep-backend/internal/database"
	"github.com/pkg/errors"
)


func HandleRetrieve(w http.ResponseWriter, r *http.Request) (*api.Response, error) {
	if r.Method != http.MethodPost {
		errorMessage := fmt.Sprintf(ErrInvalidPostRequest, RetrieveComments)
		http.Error(w, errorMessage, http.StatusMethodNotAllowed)
		return nil, errors.New(errorMessage)
	}

	commentRequest := &CommentRetrieveRequest{}
	err := json.NewDecoder(r.Body).Decode(commentRequest)

	if err != nil {
		errorMessage := fmt.Sprintf(ErrBadRequest, RetrieveComments)
		http.Error(w, errorMessage, http.StatusBadRequest)
		return nil, errors.Wrap(err, errorMessage)
	}


	db, err := database.Connect()
	if err != nil {
		errorMessage := fmt.Sprintf(ErrRetrieveDatabase, RetrieveComments)
		http.Error(w, errorMessage, http.StatusBadRequest)
		return nil, errors.Wrap(err, errorMessage)
	}

	defer db.Close()

	commentObject, err := query.GetCommentsByThreadID(db, commentRequest.ThreadID)

	if err != nil {
		errorMessage := fmt.Sprintf(ErrRetrieveDatabase, RetrieveComments)
		http.Error(w, errorMessage, http.StatusBadRequest)
		return nil, errors.Wrap(err, errorMessage)
	}
	
	data, err := json.Marshal(commentObject)
	if err != nil {
		errorMessage := fmt.Sprintf(ErrEncodeView, RetrieveComments)
		return nil, errors.Wrap(err, errorMessage)
	}

	return &api.Response{
		Payload: api.Payload{
			Data: data,
		},
		Messages: []string{fmt.Sprintf(SuccessfulRetrieveComments, RetrieveComments)},
	}, nil
}