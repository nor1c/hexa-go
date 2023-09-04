package usecase

import (
	"context"
	"gc-hexa-go/pkg/domain"
	repo "gc-hexa-go/pkg/ports/repository"
	uc "gc-hexa-go/pkg/ports/usecase"
	"gc-hexa-go/pkg/utils/customerror"
)

type OwnerUseCase struct {
	OwnerRepo repo.OwnerRepository
}

func NewOwnerUseCase(repo repo.OwnerRepository) uc.OwnerUseCase {
	return &OwnerUseCase{
		OwnerRepo: repo,
	}
}

func (uc *OwnerUseCase) FindAllOwners(ctx context.Context) ([]*domain.Owner, error) {
	users, err := uc.OwnerRepo.FindAllOwners(ctx)
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (uc *OwnerUseCase) FindOwnerById(id int, ctx context.Context) (*domain.Owner, error) {
	owner, err := uc.OwnerRepo.FindOwnerById(id, ctx)
	if err != nil {
		return nil, err
	}

	return owner, nil
}

func (uc *OwnerUseCase) CreateNewOwner(reqBody *domain.Owner, ctx context.Context) (*domain.Owner, error) {
	// save user input to database
	createdOwnerId, err := uc.OwnerRepo.CreateNewOwner(reqBody, ctx)
	if err != nil {
		return nil, err
	}

	// if returned ID is 0
	if createdOwnerId == 0 {
		return nil, customerror.ErrReturningZero
	}

	// get owner detail by ID
	owner, err := uc.OwnerRepo.FindOwnerById(createdOwnerId, ctx)
	if err != nil {
		return nil, err
	}

	return owner, nil
}

func (uc *OwnerUseCase) UpdateOwnerDetail(id int, reqBody *domain.Owner, ctx context.Context) (*domain.Owner, error) {
	err := uc.OwnerRepo.UpdateOwnerDetail(id, reqBody, ctx)
	if err != nil {
		return nil, err
	}

	// get owner detail by ID
	owner, err := uc.FindOwnerById(id, ctx)
	if err != nil {
		return nil, err
	}

	return owner, nil
}

func (uc *OwnerUseCase) RemoveOwner(id int, ctx context.Context) error {
	err := uc.OwnerRepo.RemoveOwner(id, ctx)
	if err != nil {
		return err
	}

	return nil
}

func (uc *OwnerUseCase) DeactivateOwner(id int, ctx context.Context) error {
	err := uc.OwnerRepo.DeactivateOwner(id, ctx)
	if err != nil {
		return err
	}

	return nil
}
