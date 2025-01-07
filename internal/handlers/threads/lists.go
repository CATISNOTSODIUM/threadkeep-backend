package Threads

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"github.com/CATISNOTSODIUM/threadkeep-backend/internal/api"
	"github.com/CATISNOTSODIUM/threadkeep-backend/internal/dataaccess/query"
	"github.com/CATISNOTSODIUM/threadkeep-backend/internal/database"
	"github.com/CATISNOTSODIUM/threadkeep-backend/internal/models"
	"github.com/CATISNOTSODIUM/threadkeep-backend/internal/utils"
	"github.com/pkg/errors"
)


func HandleList(w http.ResponseWriter, r *http.Request) (*api.Response, error) {
	_skip := r.URL.Query().Get("skip")
	_max_per_page := r.URL.Query().Get("max_per_page")
	name := r.URL.Query().Get("name")
	_tags := r.URL.Query().Get("tags")
	skip, err := strconv.Atoi(_skip)
	if err != nil {
		skip = 0 // default
	}
	userID := r.URL.Query().Get("userID") // fetch saved thread
	max_per_page, err := strconv.Atoi(_max_per_page)
	if err != nil {
		max_per_page = 10 // default
	}
	tags := strings.Split(_tags, ",")
	
	if (_tags == ""){
		tags = []string{}
	}

	db, err := database.Connect()
	
	if err != nil {
		return utils.WrapHTTPError(err, http.StatusInternalServerError)
	}

	defer db.Close()


	threadsObject := []*models.Thread{}
	if (len(userID) > 0) { 
		threadsObject, err = query.GetSavedThreads(db, userID)
	} else {
		threadsObject, err = query.GetThreads(db, skip, max_per_page, name, tags)
	}

	if err != nil {
		return utils.WrapHTTPError(err, http.StatusInternalServerError)
	}

	data, err := json.Marshal(threadsObject)

	if err != nil {
		return utils.WrapHTTPError(err, http.StatusInternalServerError)
	}
	
	return utils.WrapHTTPPayload(data, SuccessfulListThreadsMessage)
}

func HandleRetrieve(w http.ResponseWriter, r *http.Request) (*api.Response, error) {
	if r.Method != http.MethodPost {
		err := errors.New(ErrInvalidPostRequest)
		return utils.WrapHTTPError(err, http.StatusBadRequest)
	}

	threadRequest := &ThreadRetrieveRequest{}
	err := json.NewDecoder(r.Body).Decode(threadRequest)

	if err != nil {
		return utils.WrapHTTPError(err, http.StatusBadRequest)
	}


	db, err := database.Connect()
	if err != nil {
		return utils.WrapHTTPError(err, http.StatusInternalServerError)
	}

	defer db.Close()

	threadObject, err := query.GetThreadByID(db, threadRequest.ThreadID)

	if err != nil {
		return utils.WrapHTTPError(err, http.StatusInternalServerError)
	}
	

	data, err := json.Marshal(threadObject)
	if err != nil {
		return utils.WrapHTTPError(err, http.StatusInternalServerError)
	}

	return utils.WrapHTTPPayload(data, SuccessfulRetrieveThreadsMessage)
}

func HandleCount(w http.ResponseWriter, r *http.Request) (*api.Response, error) {
	db, err := database.Connect()
	
	if err != nil {
		return utils.WrapHTTPError(err, http.StatusInternalServerError)
	}

	defer db.Close()

	threadsObject, err := query.CountThreads(db)
	if err != nil {
		return utils.WrapHTTPError(err, http.StatusInternalServerError)
	}
	data, err := json.Marshal(threadsObject)
	if err != nil {
		return utils.WrapHTTPError(err, http.StatusInternalServerError)
	}
	
	return utils.WrapHTTPPayload(data, SuccessfulListThreadsMessage)
}