package presenter

import "net/http"

const (
	ErrConflict   = http.StatusConflict            // action cannot be performed - example: duplicate email
	ErrInternal   = http.StatusInternalServerError // referring to the language itself or to the server where the code runs. Example: saving a file, marshaling a json
	ErrBadRequest = http.StatusBadRequest          // The server cannot or will not process the request. example: malformed request syntax, invalid request message framing, or deceptive request routing
	ErrNotFound   = http.StatusNotFound            // page or entity does not exist
)

// Web response provide response struct formatting
// Representing a success of endpoint hit
type WebResponse struct {
	Code   int         `json:"code"`
	Status string      `json:"status"`
	Data   interface{} `json:"data"`
}
