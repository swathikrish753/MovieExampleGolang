package repository

import "errors"

// ErrNotFound is returned when a metadata entry is not found in the repository.
var ErrNotFound = errors.New("metadata not found")
