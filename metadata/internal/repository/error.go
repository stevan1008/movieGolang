package repository

import "errors"

// ErrorNotFound is returned when a request record is not found.
var ErrorNotFound = errors.New("not found")
