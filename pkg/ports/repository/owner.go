package infrastructure

import (
	"context"
	"gc-hexa-go/pkg/domain"
)

type OwnerRepository interface {
	FindAllOwners(ctx context.Context) ([]*domain.Owner, error)
	FindOwnerById(id int, ctx context.Context) (*domain.Owner, error)
	CreateNewOwner(reqBody *domain.Owner, ctx context.Context) (int, error)
	UpdateOwnerDetail(id int, reqBody *domain.Owner, ctx context.Context) error
	RemoveOwner(id int, ctx context.Context) error
	DeactivateOwner(id int, ctx context.Context) error
}
