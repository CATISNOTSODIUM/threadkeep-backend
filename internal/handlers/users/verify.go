package users

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/CATISNOTSODIUM/threadkeep-backend/internal/api"
	"github.com/CATISNOTSODIUM/threadkeep-backend/internal/dataaccess/query"
	"github.com/CATISNOTSODIUM/threadkeep-backend/internal/database"
	"github.com/pkg/errors"
)



func HandleVerify(w http.ResponseWriter, r *http.Request) (*api.Response, error) {
	if r.Method != http.MethodPost {
		errorMessage := fmt.Sprintf(ErrInvalidPostRequest, VerifyUser)
		http.Error(w, errorMessage, http.StatusMethodNotAllowed)
		return nil, errors.New(errorMessage)
	}

	userRequest := &UserVerifyRequest{}
	err := json.NewDecoder(r.Body).Decode(userRequest)

	if err != nil {
		errorMessage := fmt.Sprintf(ErrBadRequest, VerifyUser)
		http.Error(w, errorMessage, http.StatusBadRequest)
		return nil, errors.Wrap(err, errorMessage)
	}


	db, err := database.Connect()
	if err != nil {
		errorMessage := fmt.Sprintf(ErrRetrieveDatabase, VerifyUser)
		http.Error(w, errorMessage, http.StatusBadRequest)
		return nil, errors.Wrap(err, errorMessage)
	}

	defer db.Close()

	user, isVerified := query.VerifyUser(db, userRequest.Name, userRequest.Password)
	
	data, err := json.Marshal(isVerified)
	
	if isVerified {
		data, err = json.Marshal(user)
	}
	
	if err != nil {
		errorMessage := fmt.Sprintf(ErrEncodeView, VerifyUser)
		return nil, errors.Wrap(err, errorMessage)
	}

	return &api.Response{
		Payload: api.Payload{
			Data: data,
		},
		Messages: []string{SuccessfulVerifyUser},
	}, nil
	
}