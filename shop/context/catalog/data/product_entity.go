package data

import "github.com/google/uuid"

type ProductEntity struct {
	ID    uuid.UUID
	Sku   string
	Name  string
	Price float64
}
