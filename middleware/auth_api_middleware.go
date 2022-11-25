package middleware

import (
	"encoding/json"
	"net/http"

	"github.com/helmimuzkr/golang-restapi/presenter"
	"github.com/julienschmidt/httprouter"
)

type authApiMiddleware struct {
	Next http.Handler

	ApiKey string
}

func NewAuthApiMiddleware(router *httprouter.Router, apiKey string) *authApiMiddleware {
	return &authApiMiddleware{
		Next:   router,
		ApiKey: apiKey,
	}
}

func (middleware authApiMiddleware) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	key := r.Header.Get("X-Api-Key")
	if key == middleware.ApiKey {
		middleware.Next.ServeHTTP(w, r)
	} else {
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnauthorized)

		errorResponse := &presenter.WebResponse{
			Code:   http.StatusUnauthorized,
			Status: "UNAUTHORIZED",
			Data:   nil,
		}

		encoder := json.NewEncoder(w)
		encoder.Encode(errorResponse)
	}
}
