package usecase

import (
	"context"
	"gc-hexa-go/pkg/domain"
	repo "gc-hexa-go/pkg/ports/repository"
	"gc-hexa-go/pkg/utils/customerror"
)

type PetUseCase struct {
	PetRepo repo.PetRepository
}

func NewPetUseCase(petRepo repo.PetRepository) *PetUseCase {
	return &PetUseCase{
		PetRepo: petRepo,
	}
}

func (uc *PetUseCase) GetAllPets(ctx context.Context) ([]*domain.Pet, error) {
	pets, err := uc.PetRepo.GetAllPets(ctx)
	if err != nil {
		return nil, err
	}

	return pets, nil
}

func (uc *PetUseCase) GetPetDetail(petID int, ctx context.Context) (*domain.Pet, error) {
	pet, err := uc.PetRepo.GetPetDetail(petID, ctx)
	if err != nil {
		return nil, err
	}

	return pet, nil
}

func (uc *PetUseCase) AddNewPet(reqBody *domain.Pet, ctx context.Context) (*domain.Pet, error) {
	petID, err := uc.PetRepo.AddNewPet(reqBody, ctx)
	if err != nil {
		return nil, err
	}

	if petID == 0 {
		return nil, customerror.ErrReturningZero
	}

	// if success, get newly created pet detail
	pet, err := uc.PetRepo.GetPetDetail(petID, ctx)
	if err != nil {
		return nil, err
	}

	return pet, nil
}

func (uc *PetUseCase) UpdatePet(reqBody *domain.Pet, petID int, ctx context.Context) error {
	err := uc.PetRepo.UpdatePet(reqBody, petID, ctx)
	if err != nil {
		return err
	}

	return nil
}

func (uc *PetUseCase) DeletePet(petID int, ctx context.Context) error {
	err := uc.PetRepo.DeletePet(petID, ctx)
	if err != nil {
		return err
	}

	return nil
}
