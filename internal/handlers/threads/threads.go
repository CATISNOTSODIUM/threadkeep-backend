package Threads

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/CATISNOTSODIUM/taggy-backend/internal/api"
	"github.com/CATISNOTSODIUM/taggy-backend/internal/dataaccess/query"
	"github.com/CATISNOTSODIUM/taggy-backend/internal/database"
	"github.com/pkg/errors"
)

const (
	ListThreads = "Threads.HandleList"

	SuccessfulListThreadsMessage = "Successfully listed Threads"
	ErrRetrieveDatabase        = "Failed to retrieve database in %s"
	ErrRetrieveThreads           = "Failed to retrieve Threads in %s"
	ErrEncodeView              = "Failed to retrieve Threads in %s"
)

func HandleList(w http.ResponseWriter, r *http.Request) (*api.Response, error) {
	db, err := database.Connect()
	defer db.Close()
	
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf(ErrRetrieveDatabase, ListThreads))
	}
	Threads, err := query.GetThreads(db)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf(ErrRetrieveThreads, ListThreads))
	}
	data, err := json.Marshal(Threads)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf(ErrEncodeView, ListThreads))
	}
	return &api.Response{
		Payload: api.Payload{
			Data: data,
		},
		Messages: []string{SuccessfulListThreadsMessage},
	}, nil
}
