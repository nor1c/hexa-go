package repository

import (
	"context"
	"database/sql"
	"errors"
	"gc-hexa-go/pkg/domain"
	"gc-hexa-go/pkg/utils/customerror"

	"github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type PetRepository struct {
	db *sqlx.DB
}

func NewPetRepository(db *sqlx.DB) *PetRepository {
	return &PetRepository{
		db: db,
	}
}

func (r *PetRepository) GetAllPets(ctx context.Context) ([]*domain.Pet, error) {
	// prepare statement
	stmt, err := r.db.PreparexContext(ctx, "SELECT * FROM pets")
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

	// map to `Pet` struct
	var pets []*domain.Pet
	err = sqlx.StructScan(rows, &pets)
	if err != nil {
		return nil, err
	}

	// if no pet found
	if len(pets) == 0 {
		return nil, customerror.ErrPetsEmpty
	}

	return pets, nil
}

func (r *PetRepository) GetPetDetail(petID int, ctx context.Context) (*domain.Pet, error) {
	row := r.db.QueryRowxContext(ctx, "SELECT * FROM pets WHERE id=?", petID)

	var pet domain.Pet
	err := row.StructScan(&pet)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, customerror.ErrPetNotFound
		} else {
			return nil, err
		}
	}

	return &pet, nil
}

func (r *PetRepository) AddNewPet(reqBody *domain.Pet, ctx context.Context) (int, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return 0, err
	}
	defer tx.Rollback()

	// prepare query statement
	stmt, err := tx.PrepareContext(ctx, "INSERT INTO pets (kind_id, owner_id, name, age, adoption_date) VALUES (?, ?, ?, ?, ?)")
	if err != nil {
		return 0, err
	}
	defer stmt.Close()

	// execute query
	res, err := stmt.ExecContext(ctx, &reqBody.Kind_ID, &reqBody.Owner_ID, &reqBody.Name, &reqBody.Age, &reqBody.Adoption_Date)
	if err != nil {
		if mysqlErr, ok := err.(*mysql.MySQLError); ok && mysqlErr.Number == 1452 {
			return 0, errors.New("Foreign key error")
		} else {
			return 0, err
		}
	}

	err = tx.Commit()
	if err != nil {
		return 0, err
	}

	petID, err := res.LastInsertId()
	if err != nil {
		return 0, nil
	}

	return int(petID), nil
}

func (r *PetRepository) UpdatePet(reqBody *domain.Pet, petID int, ctx context.Context) error {
	tx, err := r.db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	stmt, err := tx.PrepareContext(ctx, "UPDATE pets SET kind_id=?, owner_id=?, name=?, age=?, adoption_date=? WHERE id=?")
	if err != nil {
		return err
	}
	defer stmt.Close()

	res, err := stmt.ExecContext(ctx, &reqBody.Kind_ID, &reqBody.Owner_ID, &reqBody.Name, &reqBody.Age, &reqBody.Adoption_Date, petID)
	if err != nil {
		return err
	}

	// check of affected rows
	affectedRow, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if affectedRow == 0 {
		return customerror.ErrPetUpdateFailed
	}

	// commit changes
	if err = tx.Commit(); err != nil {
		return err
	}

	return nil
}

func (r *PetRepository) DeletePet(petID int, ctx context.Context) error {
	stmt, err := r.db.PreparexContext(ctx, "DELETE FROM pets WHERE id=?")
	if err != nil {
		return err
	}
	defer stmt.Close()

	res, err := stmt.ExecContext(ctx, petID)
	if err != nil {
		return err
	}

	affectedRows, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if affectedRows == 0 {
		return customerror.ErrPetDeleteFailed
	}

	return nil
}
