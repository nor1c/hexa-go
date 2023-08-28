package repository

import (
	"context"
	"gc-hexa-go/pkg/domain"
	"gc-hexa-go/pkg/utils/customerror"

	"github.com/jmoiron/sqlx"
)

type ownerRepository struct {
	db *sqlx.DB
}

func NewOwnerRepository(db *sqlx.DB) *ownerRepository {
	return &ownerRepository{
		db: db,
	}
}

func (r *ownerRepository) FindAllOwners(ctx context.Context) ([]*domain.Owner, error) {
	// prepare query
	stmt, err := r.db.PrepareContext(ctx, "SELECT * FROM owners")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	// run query
	rows, err := stmt.QueryContext(ctx)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// scan to struct
	var owners []*domain.Owner
	err = sqlx.StructScan(rows, &owners)
	if err != nil {
		return nil, err
	}

	// check if owners is empty
	if len(owners) == 0 {
		return nil, customerror.ErrOwnersEmpty
	}

	return owners, nil
}

func (u *ownerRepository) FindOwnerById() (*domain.Owner, error) {
	return nil, nil
}

func (u *ownerRepository) CreateNewOwner() error {
	return nil
}

func (u *ownerRepository) UpdateOwnerDetail() error {
	return nil
}

func DeactivateOwner() error {
	return nil
}

func (u *ownerRepository) RemoveOwner() error {
	return nil
}

func (u *ownerRepository) DeactivateOwner() error {
	return nil
}
