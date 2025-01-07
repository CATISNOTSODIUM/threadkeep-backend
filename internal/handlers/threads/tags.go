package Threads

import (
	"encoding/json"
	"net/http"

	"github.com/CATISNOTSODIUM/threadkeep-backend/internal/api"
	"github.com/CATISNOTSODIUM/threadkeep-backend/internal/dataaccess/query"
	"github.com/CATISNOTSODIUM/threadkeep-backend/internal/database"
	"github.com/CATISNOTSODIUM/threadkeep-backend/internal/utils"
	"github.com/pkg/errors"
)



func HandleTag(w http.ResponseWriter, r *http.Request) (*api.Response, error) {
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

	threadObject, err := query.GetThreadTagsByID(db, threadRequest.ThreadID)

	if err != nil {
		return utils.WrapHTTPError(err, http.StatusInternalServerError)
	}
	

	data, err := json.Marshal(threadObject)
	if err != nil {
		return utils.WrapHTTPError(err, http.StatusInternalServerError)
	}

	return utils.WrapHTTPPayload(data, SuccessfulListThreadsMessage)
}


func HandleTagList(w http.ResponseWriter, r *http.Request) (*api.Response, error) {
	db, err := database.Connect()
	
	if err != nil {
		return utils.WrapHTTPError(errors.Wrap(err, ErrRetrieveDatabase), http.StatusInternalServerError)
	}

	defer db.Close()

	threadsObject, err := query.GetTags(db)
	if err != nil {
		return utils.WrapHTTPError(errors.Wrap(err, ErrRetrieveDatabase), http.StatusInternalServerError)
	}
	data, err := json.Marshal(threadsObject)
	if err != nil {
		return utils.WrapHTTPError(err, http.StatusInternalServerError)
	}
	
	return utils.WrapHTTPPayload(data, SuccessfulListThreadsMessage)
}