package Threads

import (
	"encoding/json"
	"net/http"

	"github.com/CATISNOTSODIUM/threadkeep-backend/internal/api"
	"github.com/CATISNOTSODIUM/threadkeep-backend/internal/dataaccess/mutation"
	"github.com/CATISNOTSODIUM/threadkeep-backend/internal/database"
	"github.com/CATISNOTSODIUM/threadkeep-backend/internal/utils"
	"github.com/pkg/errors"
)


func HandleCreate(w http.ResponseWriter, r *http.Request) (*api.Response, int) {
	if r.Method != http.MethodPost {
		err := errors.New(ErrInvalidPostRequest)
		return utils.WrapHTTPError(err, http.StatusBadRequest)
	}

	thread := &ThreadCreateRequest{}
	err := json.NewDecoder(r.Body).Decode(thread)
	if err != nil {
		return utils.WrapHTTPError(err, http.StatusBadRequest)
	}


	db, err := database.Connect()
	if err != nil {
		return utils.WrapHTTPError(err, http.StatusInternalServerError)
	}

	defer db.Close()

	threadObject, err := mutation.CreateThread(db, &thread.User, thread.Title, thread.Content, thread.Tags)
	
	if err != nil {
		return utils.WrapHTTPError(err, http.StatusBadRequest)
	}
	

	data, err := json.Marshal(threadObject)
	if err != nil {
		return utils.WrapHTTPError(err, http.StatusInternalServerError)
	}

	return utils.WrapHTTPPayload(data, SuccessfulCreateNewThreadMessage)
	
}