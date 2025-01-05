package Threads

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/CATISNOTSODIUM/threadkeep-backend/internal/api"
	"github.com/CATISNOTSODIUM/threadkeep-backend/internal/dataaccess/query"
	"github.com/CATISNOTSODIUM/threadkeep-backend/internal/database"
	"github.com/pkg/errors"
)



func HandleTag(w http.ResponseWriter, r *http.Request) (*api.Response, error) {
	if r.Method != http.MethodPost {
		errorMessage := fmt.Sprintf(ErrInvalidPostRequest, ListThreadsTag)
		http.Error(w, errorMessage, http.StatusMethodNotAllowed)
		return nil, errors.New(errorMessage)
	}

	threadRequest := &ThreadRetrieveRequest{}
	err := json.NewDecoder(r.Body).Decode(threadRequest)

	if err != nil {
		errorMessage := fmt.Sprintf(ErrBadRequest, ListThreadsTag)
		http.Error(w, errorMessage, http.StatusBadRequest)
		return nil, errors.Wrap(err, errorMessage)
	}


	db, err := database.Connect()
	if err != nil {
		errorMessage := fmt.Sprintf(ErrRetrieveDatabase, ListThreadsTag)
		http.Error(w, errorMessage, http.StatusBadRequest)
		return nil, errors.Wrap(err, errorMessage)
	}

	defer db.Close()

	threadObject, err := query.GetThreadTagsByID(db, threadRequest.ThreadID)

	if err != nil {
		errorMessage := fmt.Sprintf(ErrCreateThread, ListThreadsTag)
		http.Error(w, errorMessage, http.StatusBadRequest)
		return nil, errors.Wrap(err, errorMessage)
	}
	

	data, err := json.Marshal(threadObject)
	if err != nil {
		errorMessage := fmt.Sprintf(ErrEncodeView, ListThreadsTag)
		return nil, errors.Wrap(err, errorMessage)
	}

	return &api.Response{
		Payload: api.Payload{
			Data: data,
		},
		Messages: []string{SuccessfulListThreadsMessage},
	}, nil
}


func HandleTagList(w http.ResponseWriter, r *http.Request) (*api.Response, error) {
	db, err := database.Connect()
	
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf(ErrRetrieveDatabase, ListThreads))
	}

	defer db.Close()

	threadsObject, err := query.GetTags(db)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf(ErrRetrieveThreads, ListThreads))
	}
	data, err := json.Marshal(threadsObject)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf(ErrEncodeView, ListThreads))
	}
	
	return &api.Response{
		Payload: api.Payload{
			Data: data,
		},
		Messages: []string{SuccessfulListThreadsMessage},
	}, nil

}