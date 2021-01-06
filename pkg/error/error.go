package error

import "errors"

var (
	ErrEmptyDescription = errors.New("description should not be empty")
	ErrNotPositivePrice = errors.New("price should be positive number")
	ErrNotPositiveId    = errors.New("id should be positive number")
	ErrWrongSortField   = errors.New("wrong sort param")
)
