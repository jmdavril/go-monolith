package data

import "github.com/google/uuid"

type OrderEntity struct {
	ID         uuid.UUID
	TotalSpent float64
	CustomerId uuid.UUID
	Items      []LineItemEntity
}
