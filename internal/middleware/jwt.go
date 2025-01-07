package middleware

import (
	"encoding/json"
	"errors"
	"net/http"
	"github.com/CATISNOTSODIUM/threadkeep-backend/internal/utils"
)

func JWT_Middleware() func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			tokenString := r.Header.Get("Authorization")
			if tokenString == "" {
				err_json, _ := utils.WrapHTTPError(errors.New("missing authorization header"), http.StatusUnauthorized)
				json.NewEncoder(w).Encode(err_json)
				return
			}
			tokenString = tokenString[len("Bearer "):]
			
			err := utils.VerifyToken(tokenString)
			if err != nil {
				err_json, _ := utils.WrapHTTPError(errors.New("invalid jwt token"), http.StatusUnauthorized)
				json.NewEncoder(w).Encode(err_json)
				return
			}
			
			next.ServeHTTP(w, r)
		})
	}
}
