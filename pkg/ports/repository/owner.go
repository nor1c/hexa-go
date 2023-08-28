package infrastructure

import (
	"context"
	"gc-hexa-go/pkg/domain"
)

type OwnerRepository interface {
	FindAllOwners(ctx context.Context) ([]*domain.Owner, error)
	FindOwnerById() (*domain.Owner, error)
	CreateNewOwner() error
	UpdateOwnerDetail() error
	DeactivateOwner() error
	RemoveOwner() error
}
