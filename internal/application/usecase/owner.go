package usecase

import (
	"context"
	"gc-hexa-go/pkg/domain"
	r "gc-hexa-go/pkg/ports/repository"
	u "gc-hexa-go/pkg/ports/usecase"
)

type OwnerUseCase struct {
	OwnerRepo r.OwnerRepository
}

func NewOwnerUseCase(repo r.OwnerRepository) u.OwnerUseCase {
	return &OwnerUseCase{
		OwnerRepo: repo,
	}
}

func (u *OwnerUseCase) FindAllOwners(ctx context.Context) ([]*domain.Owner, error) {
	users, err := u.OwnerRepo.FindAllOwners(ctx)
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (u *OwnerUseCase) FindOwnerById() (*domain.Owner, error) {
	return nil, nil
}

func (u *OwnerUseCase) CreateNewOwner() error {
	return nil
}

func (u *OwnerUseCase) UpdateOwnerDetail() error {
	return nil
}

func DeactivateOwner() error {
	return nil
}

func (u *OwnerUseCase) RemoveOwner() error {
	return nil
}

func (u *OwnerUseCase) DeactivateOwner() error {
	return nil
}
