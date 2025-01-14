package users

import (
	"encoding/json"
	"net/http"
	"github.com/CATISNOTSODIUM/threadkeep-backend/internal/api"
	"github.com/CATISNOTSODIUM/threadkeep-backend/internal/dataaccess/query"
	"github.com/CATISNOTSODIUM/threadkeep-backend/internal/database"
	"github.com/CATISNOTSODIUM/threadkeep-backend/internal/utils"
	"github.com/pkg/errors"
)



func HandleVerify(w http.ResponseWriter, r *http.Request) (*api.Response, int) {
	if r.Method != http.MethodPost {
		err := errors.New(ErrInvalidPostRequest)
		return utils.WrapHTTPError(err, http.StatusBadRequest)
	}

	userRequest := &UserVerifyRequest{}
	err := json.NewDecoder(r.Body).Decode(userRequest)

	if err != nil {
		return utils.WrapHTTPError(err, http.StatusBadRequest)
	}


	db, err := database.Connect()
	if err != nil {
		return utils.WrapHTTPError(err, http.StatusInternalServerError)
	}

	defer db.Close()

	user, isVerified := query.VerifyUser(db, userRequest.Name, userRequest.Password)
	
	data, err := json.Marshal(isVerified)
	
	if isVerified {
		data, err = json.Marshal(user)
	}
	
	if err != nil {
		return utils.WrapHTTPError(err, http.StatusBadRequest)
	}

	return utils.WrapHTTPPayload(data, SuccessfulVerifyUser)
	
}