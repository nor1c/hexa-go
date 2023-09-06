package customerror

import "errors"

var (
	ErrTimedOut           = errors.New("Request timed out.")
	ErrInvalidRequestBody = errors.New("Invalid request body.")
	ErrReturningZero      = errors.New("Expecting non-zero value, get 0 instead.")
)
