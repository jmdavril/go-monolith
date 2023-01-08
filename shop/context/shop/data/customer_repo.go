package data

import (
	"github.com/jmoiron/sqlx"
	"log"

	"github.com/google/uuid"
)

type CustomerRepo struct {
	db *sqlx.DB
}

func NewCustomerRepo(db *sqlx.DB) *CustomerRepo {
	return &CustomerRepo{
		db: db,
	}
}

func (r *CustomerRepo) SelectCustomer(customerId uuid.UUID) (CustomerEntity, error) {
	row := r.db.QueryRowx("SELECT * FROM customer WHERE id=$1", customerId)
	var customer CustomerEntity
	err := row.StructScan(&customer)

	if err != nil {
		log.Printf("Could not run SQL query to read customer: %v", err)
		return customer, err
	}
	return customer, nil
}

func (r *CustomerRepo) InsertCustomer(c CustomerEntity) (uuid.UUID, error) {
	rows, err := r.db.Query("INSERT INTO customer(email) VALUES($1) RETURNING id", c.Email)

	if err != nil {
		log.Printf("Could not run SQL query to insert customer %v: %v", c, err)
		return uuid.Nil, err
	}

	var id uuid.UUID
	if rows.Next() {
		err := rows.Scan(&id)
		if err != nil {
			return [16]byte{}, err
		}
	}

	return id, nil
}
