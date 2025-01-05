package Threads

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
		errorMessage := fmt.Sprintf(ErrInvalidPostRequest, CreateNewThread)
		http.Error(w, errorMessage, http.StatusMethodNotAllowed)
		return nil, errors.New(errorMessage)
	}

	thread := &ThreadCreateRequest{}
	err := json.NewDecoder(r.Body).Decode(thread)
	if err != nil {
		errorMessage := fmt.Sprintf(ErrBadRequest, CreateNewThread)
		http.Error(w, errorMessage, http.StatusBadRequest)
		return nil, errors.Wrap(err, errorMessage)
	}


	db, err := database.Connect()
	if err != nil {
		errorMessage := fmt.Sprintf(ErrRetrieveDatabase, CreateNewThread)
		http.Error(w, errorMessage, http.StatusBadRequest)
		return nil, errors.Wrap(err, errorMessage)
	}

	defer db.Close()

	threadObject, err := mutation.CreateThread(db, &thread.User, thread.Title, thread.Content, thread.Tags)
	
	if err != nil {
		print(err.Error())
		errorMessage := fmt.Sprintf(ErrCreateThread, CreateNewThread)
		http.Error(w, errorMessage, http.StatusBadRequest)
		return nil, errors.Wrap(err, errorMessage)
	}
	

	data, err := json.Marshal(threadObject)
	if err != nil {
		errorMessage := fmt.Sprintf(ErrEncodeView, CreateNewThread)
		return nil, errors.Wrap(err, errorMessage)
	}

	return &api.Response{
		Payload: api.Payload{
			Data: data,
		},
		Messages: []string{SuccessfulCreateNewThreadMessage},
	}, nil
	
}