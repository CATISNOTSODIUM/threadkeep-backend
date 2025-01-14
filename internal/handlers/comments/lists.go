package Comments

import (
	"encoding/json"
	"net/http"

	"github.com/CATISNOTSODIUM/threadkeep-backend/internal/api"
	"github.com/CATISNOTSODIUM/threadkeep-backend/internal/dataaccess/query"
	"github.com/CATISNOTSODIUM/threadkeep-backend/internal/database"
	"github.com/CATISNOTSODIUM/threadkeep-backend/internal/utils"
	"github.com/pkg/errors"
)


func HandleRetrieve(w http.ResponseWriter, r *http.Request) (*api.Response, int) {
	if r.Method != http.MethodPost {
		err := errors.New(ErrInvalidPostRequest)
		return utils.WrapHTTPError(err, http.StatusBadRequest)
	}

	commentRequest := &CommentRetrieveRequest{}
	err := json.NewDecoder(r.Body).Decode(commentRequest)
	if err != nil {
		return utils.WrapHTTPError(err, http.StatusBadRequest)
	}

	db, err := database.Connect()
	if err != nil {
		return utils.WrapHTTPError(err, http.StatusInternalServerError)
	}
	defer db.Close()

	commentObject, err := query.GetCommentsByThreadID(db, commentRequest.ThreadID)

	if err != nil {
		return utils.WrapHTTPError(err, http.StatusBadRequest)
	}
	
	data, err := json.Marshal(commentObject)
	if err != nil {
		return utils.WrapHTTPError(err, http.StatusInternalServerError)
	}

	return utils.WrapHTTPPayload(data, SuccessfulRetrieveComments)
}