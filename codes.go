package httperror

import "net/http"

// TODO: generate these.

const (
	BadRequest          Status = http.StatusBadRequest
	Forbidden           Status = http.StatusForbidden
	InternalServerError Status = http.StatusInternalServerError
	NotFound            Status = http.StatusNotFound
	Unauthorized        Status = http.StatusUnauthorized
)
