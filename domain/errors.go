package domain

import "errors"

var (
	ErrExists            = errors.New("qcbalu: The item exists")
	ErrBadRequest        = errors.New("qcbalu: Bad request")
	ErrNotFound          = errors.New("qcbalu: The item was not found")
	ErrVersionIsNotValid = errors.New("qcbalu: The item version is not valid")
)
