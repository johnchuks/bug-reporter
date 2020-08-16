package responses

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// HTTPError a defined struct for a jsonified error
type HTTPError struct {
	Error string `json:"error"`
}

// JSON returns a well formated response with a status code
func JSON(w http.ResponseWriter, statusCode int, data interface{}) {
	w.WriteHeader(statusCode)
	err := json.NewEncoder(w).Encode(data)

	if err != nil {
		fmt.Fprintf(w, "%s", err.Error())
	}
}

// ERROR returns a jsonified error response along with a status code
func ERROR(w http.ResponseWriter, statusCode int, err error) {

	if err != nil {
		newError := HTTPError{
			Error: err.Error(),
		}
		JSON(w, statusCode, newError)
		return
	}

	JSON(w, http.StatusBadRequest, nil)
}
