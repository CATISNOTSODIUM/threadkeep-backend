package Threads

import (
	"encoding/json"
	"net/http"

	"github.com/CATISNOTSODIUM/threadkeep-backend/internal/api"
	"github.com/CATISNOTSODIUM/threadkeep-backend/internal/dataaccess/mutation"
	"github.com/CATISNOTSODIUM/threadkeep-backend/internal/dataaccess/query"
	"github.com/CATISNOTSODIUM/threadkeep-backend/internal/database"
	"github.com/CATISNOTSODIUM/threadkeep-backend/internal/utils"
	"github.com/pkg/errors"
)


func HandleIsLikeThread(w http.ResponseWriter, r *http.Request) (*api.Response, error) {
	if r.Method != http.MethodPost {
		err := errors.New(ErrInvalidPostRequest)
		return utils.WrapHTTPError(err, http.StatusBadRequest)
	}

	threadRequest := &ThreadIsLikeRequest{}
	if err := json.NewDecoder(r.Body).Decode(threadRequest); err != nil {
		return utils.WrapHTTPError(err, http.StatusBadRequest)
	}

	db, err := database.Connect()
	if err != nil {
		return utils.WrapHTTPError(err, http.StatusBadRequest)
	}

	defer db.Close()

	isLike, err := query.IsLikeThread(db, threadRequest.UserID, threadRequest.ThreadID)
	if err != nil {
		return utils.WrapHTTPError(err, http.StatusBadRequest)
	}

	data, err := json.Marshal(isLike)
	
	if err != nil {
		return utils.WrapHTTPError(err, http.StatusBadRequest)
	}

	return utils.WrapHTTPPayload(data, ReactThread)
}


func HandleReactThread(w http.ResponseWriter, r *http.Request) (*api.Response, error) {
	if r.Method != http.MethodPost {
		err := errors.New(ErrInvalidPostRequest)
		return utils.WrapHTTPError(err, http.StatusBadRequest)
	}

	threadRequest := &ThreadLikeRequest{}
	err := json.NewDecoder(r.Body).Decode(threadRequest)

	if err != nil {
		return utils.WrapHTTPError(err, http.StatusBadRequest)
	}


	db, err := database.Connect()
	if err != nil {
		return utils.WrapHTTPError(err, http.StatusBadRequest)
	}

	defer db.Close()

	count := 0

	if threadRequest.Reaction == VIEW {
		count, err = mutation.ViewThread(db, threadRequest.ThreadID)
	} else if threadRequest.Reaction == LIKE {
		count, err = mutation.LikeThread(db, threadRequest.UserID, threadRequest.ThreadID)
	} else if threadRequest.Reaction == UNLIKE {
		count, err = mutation.UnlikeThread(db, threadRequest.UserID, threadRequest.ThreadID)
	} else if threadRequest.Reaction == SAVED {
		count, err = mutation.SaveThread(db, threadRequest.UserID, threadRequest.ThreadID)
	} else if threadRequest.Reaction == UNSAVE {
		count, err = mutation.UnsaveThread(db, threadRequest.UserID, threadRequest.ThreadID)
	} else {
		err = errors.New(ErrInvalidReactionType)
		return utils.WrapHTTPError(err, http.StatusBadRequest)
	}
	
	if err != nil {
		return utils.WrapHTTPError(err, http.StatusBadRequest)
	}

	if (count == 1) {
		return utils.WrapHTTPSuccess(SuccessfulRequest)
	} else {
		return utils.WrapHTTPError(errors.New(ErrUpdateReaction), http.StatusInternalServerError) 
	}
}