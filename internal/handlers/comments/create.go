package Comments

import (
	"encoding/json"
	"net/http"
	"github.com/CATISNOTSODIUM/threadkeep-backend/internal/api"
	"github.com/CATISNOTSODIUM/threadkeep-backend/internal/dataaccess/mutation"
	"github.com/CATISNOTSODIUM/threadkeep-backend/internal/database"
	"github.com/CATISNOTSODIUM/threadkeep-backend/internal/utils"
	"github.com/pkg/errors"
)



func HandleCreate(w http.ResponseWriter, r *http.Request) (*api.Response, error) {
	if r.Method != http.MethodPost {
		err := errors.New(ErrInvalidPostRequest)
		return utils.WrapHTTPError(err, http.StatusBadRequest)
	}

	comment := &CommentCreateRequest{}
	err := json.NewDecoder(r.Body).Decode(comment)

	if err != nil {
		return utils.WrapHTTPError(err, http.StatusBadRequest)
	}

	db, err := database.Connect()
	if err != nil {
		return utils.WrapHTTPError(err, http.StatusInternalServerError)
	}

	defer db.Close()

	commentObject, err := mutation.CreateComment(db, comment.User.ID, comment.ThreadID, comment.Content)
	
	if err != nil {
		return utils.WrapHTTPError(err, http.StatusInternalServerError)
	}
	
	data, err := json.Marshal(commentObject)
	if err != nil {
		return utils.WrapHTTPError(err, http.StatusInternalServerError)
	}

	return utils.WrapHTTPPayload(data, SuccessfulCreateNewComments)
}

