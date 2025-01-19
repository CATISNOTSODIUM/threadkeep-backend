package users

import (
	"encoding/json"
	"net/http"
	"github.com/CATISNOTSODIUM/threadkeep-backend/internal/api"
	"github.com/CATISNOTSODIUM/threadkeep-backend/internal/dataaccess/query"
	"github.com/CATISNOTSODIUM/threadkeep-backend/internal/database"
	"github.com/CATISNOTSODIUM/threadkeep-backend/internal/utils"
)


func HandleList(w http.ResponseWriter, r *http.Request) (*api.Response, int) {
	db, err := database.Connect()
	if err != nil {
		return utils.WrapHTTPError(err, http.StatusInternalServerError)
	}

	defer db.Close()

	users, err := query.GetUsers(db)
	if err != nil {
		return utils.WrapHTTPError(err, http.StatusInternalServerError)
	}

	data, err := json.Marshal(users)
	if err != nil {
		return utils.WrapHTTPError(err, http.StatusInternalServerError)
	}

	return utils.WrapHTTPPayload(data, SuccessfulListUsersMessage)
}
