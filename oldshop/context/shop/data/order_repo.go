package data

import (
	"github.com/jmoiron/sqlx"
	"log"

	"github.com/google/uuid"
)

type OrderRepo struct {
	db *sqlx.DB
}

func NewOrderRepo(db *sqlx.DB) *OrderRepo {
	return &OrderRepo{
		db: db,
	}
}

func (r *OrderRepo) InsertOrder(o OrderEntity) (uuid.UUID, error) {
	rows, err := r.db.Query("INSERT INTO shop_order(total_spent, customer_id) VALUES($1,$2) RETURNING id", o.TotalSpent, o.CustomerId)
	if err != nil {
		log.Printf("Could not run SQL query to insert order %v: %v", o, err)
		return uuid.Nil, err
	}
	var orderId uuid.UUID
	if rows.Next() {
		err := rows.Scan(&orderId)
		if err != nil {
			return [16]byte{}, err
		}
	}

	_, err = r.db.NamedExec(`INSERT INTO line_item (order_id, line_index, sku, quantity, unit_price)
        VALUES (:order_id, :line_index, :sku, :quantity, :unit_price)`, o.Items)

	if err != nil {
		log.Printf("Could not run SQL query to insert order items %v: %v", o, err)
		return uuid.Nil, err
	}

	return orderId, nil
}
