package httperror

import (
	"errors"
	"net/http"
	"strings"
)

// Status implements a error-like HTTP status.
type Status int

// Error returns default text for HTTP status.
func (s Status) Error() string {
	return http.StatusText(int(s))
}

// Write tries to find Status in err and use it as return status, otherwise prints 500.
func Write(w http.ResponseWriter, err error) {
	errorMessage := err.Error()
	var statusCode int
	var e Status
	if errors.As(err, &e) {
		statusCode = int(e)
		errorMessage = strings.TrimSpace(strings.TrimPrefix(errorMessage, http.StatusText(statusCode)+":"))
	} else {
		statusCode = http.StatusInternalServerError
	}
	http.Error(w, errorMessage, statusCode)
}
