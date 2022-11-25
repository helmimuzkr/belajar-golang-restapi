package exception

import (
	"encoding/json"
	"net/http"

	"github.com/helmimuzkr/golang-restapi/presenter"
)

func ErrorWebResponse(w http.ResponseWriter, code int, err error) {

	switch code {
	case http.StatusBadRequest:
		writeErrorResponse(w, code, "BAD_REQUEST", err)
	case http.StatusNotFound:
		writeErrorResponse(w, code, "NOT_FOUND", err)
	default:
		writeErrorResponse(w, code, "INTERNAL_SERVER_ERROR", err)
	}

}

func writeErrorResponse(w http.ResponseWriter, code int, status string, err error) {
	errorWebResponse := &presenter.WebResponse{
		Code:   code,
		Status: status,
		Data:   err.Error(),
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)

	encoder := json.NewEncoder(w)
	err = encoder.Encode(errorWebResponse)
	if err != nil {
		http.Error(w, "INTERNAL_SERVER_ERROR", 500)
	}
}
