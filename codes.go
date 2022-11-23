package httperror

import "net/http"

// TODO: generate these.

const (
	BadRequest          Status = http.StatusBadRequest
	InternalServerError Status = http.StatusInternalServerError
)
