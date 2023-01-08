package domain

import (
	"github.com/google/uuid"
	"github.com/jmdavril/template/shop/context/shop/data"
	"github.com/samber/lo"
)

type Order struct {
	ID         uuid.UUID
	CustomerId uuid.UUID
	Items      []LineItem
}

func (o Order) TotalSpent() float64 {
	return lo.Reduce(o.Items, func(agg float64, i LineItem, _ int) float64 {
		return agg + (i.TotalSpent())
	}, 0)
}

func (o Order) OrderDto() data.OrderEntity {
	return data.OrderEntity{
		ID:         o.ID,
		TotalSpent: o.TotalSpent(),
		Items: lo.Map(o.Items, func(i LineItem, idx int) data.LineItemEntity {
			return i.LineItemDto(o.ID, idx)
		}),
	}
}

func (o Order) CopyWithId(id uuid.UUID) Order {
	return Order{
		ID:         id,
		CustomerId: o.CustomerId,
		Items:      o.Items,
	}
}

type LineItem struct {
	OrderId   uuid.UUID
	Sku       string
	Quantity  int64
	UnitPrice float64
}

func (i LineItem) TotalSpent() float64 {
	return (i.UnitPrice * float64(i.Quantity))
}

func (i LineItem) LineItemDto(orderId uuid.UUID, lineIdx int) data.LineItemEntity {
	return data.LineItemEntity{
		OrderId:   orderId,
		LineIndex: int64(lineIdx + 1),
		Sku:       i.Sku,
		Quantity:  i.Quantity,
		UnitPrice: i.UnitPrice,
	}
}
