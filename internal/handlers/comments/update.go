package Comments

import (
	"encoding/json"
	"net/http"

	"github.com/pkg/errors"

	"github.com/CATISNOTSODIUM/threadkeep-backend/internal/api"
	"github.com/CATISNOTSODIUM/threadkeep-backend/internal/dataaccess/mutation"
	"github.com/CATISNOTSODIUM/threadkeep-backend/internal/database"
	"github.com/CATISNOTSODIUM/threadkeep-backend/internal/utils"
)


func HandleUpdate(w http.ResponseWriter, r *http.Request) (*api.Response, int) {
	if r.Method != http.MethodPost {
		err := errors.New(ErrInvalidPostRequest)
		return utils.WrapHTTPError(err, http.StatusBadRequest)
	}

	comment := &CommentUpdateRequest{}
	err := json.NewDecoder(r.Body).Decode(comment)

	if err != nil {
		return utils.WrapHTTPError(err, http.StatusBadRequest)
	}


	db, err := database.Connect()
	if err != nil {
		return utils.WrapHTTPError(err, http.StatusInternalServerError)
	}

	defer db.Close()

	commentObject, err := mutation.UpdateComment(db, comment.CommentID, comment.Content)
	
	if err != nil {
		return utils.WrapHTTPError(err, http.StatusBadRequest)
	}
	

	data, err := json.Marshal(commentObject)
	if err != nil {
		return utils.WrapHTTPError(err, http.StatusInternalServerError)
	}

	return utils.WrapHTTPPayload(data, SuccessfulUpdateCommentMessage)
	
}