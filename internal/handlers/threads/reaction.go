package Threads

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/CATISNOTSODIUM/threadkeep-backend/internal/api"
	"github.com/CATISNOTSODIUM/threadkeep-backend/internal/dataaccess/mutation"
	"github.com/CATISNOTSODIUM/threadkeep-backend/internal/dataaccess/query"
	"github.com/CATISNOTSODIUM/threadkeep-backend/internal/database"
	"github.com/pkg/errors"
)


func HandleIsLikeThread(w http.ResponseWriter, r *http.Request) (*api.Response, error) {
	if r.Method != http.MethodPost {
		errorMessage := fmt.Sprintf(ErrInvalidPostRequest, ReactThread)
		http.Error(w, errorMessage, http.StatusMethodNotAllowed)
		return nil, errors.New(errorMessage)
	}

	threadRequest := &ThreadIsLikeRequest{}
	err := json.NewDecoder(r.Body).Decode(threadRequest)

	if err != nil {
		errorMessage := fmt.Sprintf(ErrBadRequest, ReactThread)
		http.Error(w, errorMessage, http.StatusBadRequest)
		return nil, errors.Wrap(err, errorMessage)
	}


	db, err := database.Connect()
	if err != nil {
		errorMessage := fmt.Sprintf(ErrRetrieveDatabase, ReactThread)
		http.Error(w, errorMessage, http.StatusBadRequest)
		return nil, errors.Wrap(err, errorMessage)
	}

	defer db.Close()

	isLike, err := query.IsLikeThread(db, threadRequest.UserID, threadRequest.ThreadID)
	if err != nil {
		errorMessage := fmt.Sprintf(ErrBadRequest, ReactThread)
		http.Error(w, errorMessage, http.StatusBadRequest)
		return nil, errors.Wrap(err, errorMessage)
	}

	data, err := json.Marshal(isLike)
	
	
	if err != nil {
		errorMessage := fmt.Sprintf(ErrBadRequest, ReactThread)
		http.Error(w, errorMessage, http.StatusBadRequest)
		return nil, errors.Wrap(err, errorMessage)
	}

	return &api.Response{
		Payload: api.Payload{
			Data: data,
		},
		Messages: []string{ReactThread},
	}, nil
}

func HandleLikeThread(w http.ResponseWriter, r *http.Request) (*api.Response, error) {
	if r.Method != http.MethodPost {
		errorMessage := fmt.Sprintf(ErrInvalidPostRequest, ReactThread)
		http.Error(w, errorMessage, http.StatusMethodNotAllowed)
		return nil, errors.New(errorMessage)
	}

	threadRequest := &ThreadLikeRequest{}
	err := json.NewDecoder(r.Body).Decode(threadRequest)

	if err != nil {
		errorMessage := fmt.Sprintf(ErrBadRequest, ReactThread)
		http.Error(w, errorMessage, http.StatusBadRequest)
		return nil, errors.Wrap(err, errorMessage)
	}


	db, err := database.Connect()
	if err != nil {
		errorMessage := fmt.Sprintf(ErrRetrieveDatabase, ReactThread)
		http.Error(w, errorMessage, http.StatusBadRequest)
		return nil, errors.Wrap(err, errorMessage)
	}

	defer db.Close()

	count := 0

	if threadRequest.Reaction == VIEW {
		count, err = mutation.ViewThread(db, threadRequest.ThreadID)
	} else if threadRequest.Reaction == LIKE {
		count, err = mutation.LikeThread(db, threadRequest.UserID, threadRequest.ThreadID)
	} else if threadRequest.Reaction == UNLIKE {
		count, err = mutation.UnlikeThread(db, threadRequest.UserID, threadRequest.ThreadID)
	} else if threadRequest.Reaction == SAVED {
		count, err = mutation.SaveThread(db, threadRequest.UserID, threadRequest.ThreadID)
	} else if threadRequest.Reaction == UNSAVE {
		count, err = mutation.UnsaveThread(db, threadRequest.UserID, threadRequest.ThreadID)
	} else {
		errorMessage := fmt.Sprintf(ErrBadRequest, ReactThread)
		http.Error(w, errorMessage, http.StatusBadRequest)
		return nil, errors.Wrap(err, errorMessage)
	}
	
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf(ErrUpdateThreads, ReactThread))
	}

	responseMessage := ""
	if (count == 1) {
		responseMessage = fmt.Sprintf(SuccessfulUpdateThreadMessage, ReactThread)
	} else {
		responseMessage = fmt.Sprintf(ErrUpdateThreads, ReactThread)
	}
	return &api.Response{
		Payload: api.Payload{},
		Messages: []string{responseMessage},
	}, nil
	
}