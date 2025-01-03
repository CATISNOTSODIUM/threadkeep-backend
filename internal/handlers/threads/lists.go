package Threads

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/CATISNOTSODIUM/taggy-backend/internal/api"
	"github.com/CATISNOTSODIUM/taggy-backend/internal/dataaccess/query"
	"github.com/CATISNOTSODIUM/taggy-backend/internal/database"
	"github.com/pkg/errors"
)


func HandleList(w http.ResponseWriter, r *http.Request) (*api.Response, error) {
	_skip := r.URL.Query().Get("skip")
	_max_per_page := r.URL.Query().Get("max_per_page")
	name := r.URL.Query().Get("name")
	_tags := r.URL.Query().Get("tags")
	
	skip, err := strconv.Atoi(_skip)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf(ErrParsingParams, ListThreads))
	}
	max_per_page, err := strconv.Atoi(_max_per_page)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf(ErrParsingParams, ListThreads))
	}
	tags := strings.Split(_tags, ",")
	
	if (_tags == ""){
		tags = []string{}
	}

	db, err := database.Connect()
	
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf(ErrRetrieveDatabase, ListThreads))
	}

	defer db.Close()


	threadsObject, err := query.GetThreads(db, skip, max_per_page, name, tags)
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

func HandleRetrieve(w http.ResponseWriter, r *http.Request) (*api.Response, error) {
	if r.Method != http.MethodPost {
		errorMessage := fmt.Sprintf(ErrInvalidPostRequest, RetrieveThread)
		http.Error(w, errorMessage, http.StatusMethodNotAllowed)
		return nil, errors.New(errorMessage)
	}

	threadRequest := &ThreadRetrieveRequest{}
	err := json.NewDecoder(r.Body).Decode(threadRequest)

	if err != nil {
		errorMessage := fmt.Sprintf(ErrBadRequest, RetrieveThread)
		http.Error(w, errorMessage, http.StatusBadRequest)
		return nil, errors.Wrap(err, errorMessage)
	}


	db, err := database.Connect()
	if err != nil {
		errorMessage := fmt.Sprintf(ErrRetrieveDatabase, RetrieveThread)
		http.Error(w, errorMessage, http.StatusBadRequest)
		return nil, errors.Wrap(err, errorMessage)
	}

	defer db.Close()

	threadObject, err := query.GetThreadByID(db, threadRequest.ThreadID)

	if err != nil {
		errorMessage := fmt.Sprintf(ErrCreateThread, RetrieveThread)
		http.Error(w, errorMessage, http.StatusBadRequest)
		return nil, errors.Wrap(err, errorMessage)
	}
	

	data, err := json.Marshal(threadObject)
	if err != nil {
		errorMessage := fmt.Sprintf(ErrEncodeView, RetrieveThread)
		return nil, errors.Wrap(err, errorMessage)
	}


	return &api.Response{
		Payload: api.Payload{
			Data: data,
		},
		Messages: []string{fmt.Sprintf(SuccessfulRetrieveThreadsMessage, RetrieveThread)},
	}, nil
}

func HandleCount(w http.ResponseWriter, r *http.Request) (*api.Response, error) {
	db, err := database.Connect()
	
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf(ErrRetrieveDatabase, CountThreads))
	}

	defer db.Close()

	threadsObject, err := query.CountThreads(db)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf(ErrRetrieveThreads, CountThreads))
	}
	data, err := json.Marshal(threadsObject)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf(ErrEncodeView, CountThreads))
	}
	
	return &api.Response{
		Payload: api.Payload{
			Data: data,
		},
		Messages: []string{SuccessfulListThreadsMessage},
	}, nil
}