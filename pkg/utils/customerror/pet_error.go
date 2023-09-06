package customerror

import "errors"

var (
	ErrPetsEmpty       = errors.New("Pets empty.")
	ErrPetNotFound     = errors.New("Pet not found.")
	ErrPetUpdateFailed = errors.New("Failed updating pet data.")
	ErrPetDeleteFailed = errors.New("Failed to delete selected pet.")
)
