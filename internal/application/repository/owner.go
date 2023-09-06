package repository

import (
	"context"
	"database/sql"
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
	stmt, err := r.db.PreparexContext(ctx, "SELECT * FROM owners")
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

func (r *ownerRepository) FindOwnerById(id int, ctx context.Context) (*domain.Owner, error) {
	row := r.db.QueryRowxContext(ctx, "SELECT * FROM owners WHERE id=?", id)

	var owner domain.Owner
	err := row.StructScan(&owner)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, customerror.ErrOwnerNotFound
		} else {
			return nil, err
		}
	}

	return &owner, nil
}

func (u *ownerRepository) CreateNewOwner(reqBody *domain.Owner, ctx context.Context) (int, error) {
	// begin transaction
	tx, err := u.db.Begin()
	if err != nil {
		return 0, err
	}
	defer tx.Rollback()

	// prepare statement
	txStmt, err := tx.PrepareContext(ctx, "INSERT INTO owners (name, age, address) VALUES (?,?,?)")
	if err != nil {
		return 0, err
	}
	defer txStmt.Close()

	// run insert query
	res, err := txStmt.ExecContext(ctx, &reqBody.Name, &reqBody.Age, &reqBody.Address)
	if err != nil {
		return 0, err
	}

	// commit
	err = tx.Commit()
	if err != nil {
		return 0, err
	}

	// get inserted owner data
	lastInsertedID, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(lastInsertedID), nil
}

func (u *ownerRepository) UpdateOwnerDetail(id int, reqBody *domain.Owner, ctx context.Context) error {
	// prepare statement/query
	stmt, err := u.db.PreparexContext(ctx, "UPDATE owners SET name=?, age=?, address=? WHERE id=?")
	if err != nil {
		return err
	}
	defer stmt.Close()

	// execute query
	_, err = stmt.ExecContext(ctx, &reqBody.Name, &reqBody.Age, &reqBody.Address, id)
	if err != nil {
		return err
	}

	return nil
}

func (u *ownerRepository) RemoveOwner(id int, ctx context.Context) error {
	stmt, err := u.db.PreparexContext(ctx, "DELETE FROM owners WHERE id=?")
	if err != nil {
		return err
	}

	res, err := stmt.ExecContext(ctx, id)
	if err != nil {
		return err
	}

	if rowsAffected, _ := res.RowsAffected(); rowsAffected == 0 {
		return customerror.ErrOwnerRemoveFailed
	}

	return nil
}

func (u *ownerRepository) DeactivateOwner(id int, ctx context.Context) error {
	_, err := u.db.ExecContext(ctx, "UPDATE owners SET is_active=0")
	if err != nil {
		return err
	}

	return nil
}
