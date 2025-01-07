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


func HandleUpdate(w http.ResponseWriter, r *http.Request) (*api.Response, error) {
	if r.Method != http.MethodPost {
		err := errors.New(ErrInvalidPostRequest)
		return utils.WrapHTTPError(err, http.StatusBadRequest)
	}

	thread := &ThreadUpdateRequest{}
	err := json.NewDecoder(r.Body).Decode(thread)

	if err != nil {
		return utils.WrapHTTPError(err, http.StatusBadRequest)
	}

	db, err := database.Connect()
	if err != nil {
		return utils.WrapHTTPError(err, http.StatusInternalServerError)
	}

	defer db.Close()

	threadObject, err := mutation.UpdateThread(db, thread.ThreadID, &thread.User, thread.Title, thread.Content)
	
	if err != nil {
		return utils.WrapHTTPError(err, http.StatusInternalServerError)
	}
	

	data, err := json.Marshal(threadObject)
	if err != nil {
		return utils.WrapHTTPError(err, http.StatusInternalServerError)
	}

	return utils.WrapHTTPPayload(data, SuccessfulUpdateThreadMessage)
	
}