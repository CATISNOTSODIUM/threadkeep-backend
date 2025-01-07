package utils

import (
	"fmt"
	"net/http"
	"runtime"

	"github.com/pkg/errors"

	"github.com/CATISNOTSODIUM/threadkeep-backend/internal/api"
)
func CallerName(skip int) string {
        pc, _, _, ok := runtime.Caller(skip + 1)
        if !ok {
                return ""
        }
        f := runtime.FuncForPC(pc)
        if f == nil {
                return ""
        }
        return f.Name()
}

// simple error handler system
// report error based on function name

func WrapHTTPError(err error, errorCode int) (*api.Response, error) {
        errorMessage := fmt.Sprintln(CallerName(1), err.Error())
        return &api.Response{ 
                Payload: api.Payload{},
                Messages: []string{errorMessage},
                ErrorCode: errorCode,
	}, errors.Wrap(err, errorMessage)
}

func WrapHTTPSuccess(message string) (*api.Response, error) {
        successMessage := fmt.Sprintln(CallerName(1), message)
        return &api.Response{ 
                Payload: api.Payload{},
                Messages: []string{successMessage},
                ErrorCode: http.StatusOK,
	}, nil
}

func WrapHTTPPayload(data []byte, message string) (*api.Response, error) {
        return &api.Response{
                Payload: api.Payload{
                        Data: data,
                },
                Messages: []string{message},
        }, nil
}