package usecase

import (
	"context"
	"gc-hexa-go/pkg/domain"
)

type PetUseCase interface {
	GetAllPets(ctx context.Context) ([]*domain.Pet, error)
	GetPetDetail(petID int, ctx context.Context) (*domain.Pet, error)
	AddNewPet(reqBody *domain.Pet, ctx context.Context) (*domain.Pet, error)
	UpdatePet(reqBody *domain.Pet, petID int, ctx context.Context) error
	DeletePet(petID int, ctx context.Context) error
}
