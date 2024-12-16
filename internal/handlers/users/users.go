package users

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/CATISNOTSODIUM/taggy-backend/internal/api"
	"github.com/CATISNOTSODIUM/taggy-backend/internal/dataaccess/query"
	"github.com/CATISNOTSODIUM/taggy-backend/internal/database"
	"github.com/pkg/errors"
)


func HandleList(w http.ResponseWriter, r *http.Request) (*api.Response, error) {
	db, err := database.Connect()
	
	
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf(ErrRetrieveDatabase, ListUsers))
	}

	defer db.Close()

	users, err := query.GetUsers(db)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf(ErrRetrieveUsers, ListUsers))
	}
	data, err := json.Marshal(users)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf(ErrEncodeView, ListUsers))
	}
	return &api.Response{
		Payload: api.Payload{
			Data: data,
		},
		Messages: []string{SuccessfulListUsersMessage},
	}, nil
}
