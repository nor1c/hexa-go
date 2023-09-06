package customerror

import "errors"

var (
	ErrOwnersEmpty       = errors.New("Owners empty.")
	ErrOwnerNotFound     = errors.New("Owner not found.")
	ErrOwnerDuplicate    = errors.New("Owner already exists.")
	ErrOwnerUpdateFailed = errors.New("Failed updating owner data.")
	ErrOwnerRemoveFailed = errors.New("Failed to remove selected owner.")
)
