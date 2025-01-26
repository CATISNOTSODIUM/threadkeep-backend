package utils

import (
	"fmt"
	"net/http"
	"runtime"
	"github.com/CATISNOTSODIUM/threadkeep-backend/internal/api"
)

// Trace the origin of the function name
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

func WrapHTTPError(err error, errorCode int) (*api.Response, int) {
        errorMessage := fmt.Sprintln(CallerName(1), err.Error())
        return &api.Response{ 
                Payload: api.Payload{},
                Messages: []string{errorMessage},
                ErrorCode: errorCode,
	}, errorCode
}

func WrapHTTPSuccess(message string) (*api.Response, int) {
        successMessage := fmt.Sprintln(CallerName(1), message)
        return &api.Response{ 
                Payload: api.Payload{},
                Messages: []string{successMessage},
                ErrorCode: http.StatusOK,
	}, http.StatusOK
}

func WrapHTTPPayload(data []byte, message string) (*api.Response, int) {
        return &api.Response{
                Payload: api.Payload{
                        Data: data,
                },
                Messages: []string{message},
        }, http.StatusOK
}