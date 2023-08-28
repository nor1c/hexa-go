package usecase

import (
	"context"
	"gc-hexa-go/pkg/domain"
)

type OwnerUseCase interface {
	FindAllOwners(ctx context.Context) ([]*domain.Owner, error)
	FindOwnerById() (*domain.Owner, error)
	CreateNewOwner() error
	UpdateOwnerDetail() error
	DeactivateOwner() error
	RemoveOwner() error
}
