package data

import "github.com/google/uuid"

type LineItemEntity struct {
	OrderId   uuid.UUID
	LineIndex int64
	Sku       string
	Quantity  int64
	UnitPrice float64
}
