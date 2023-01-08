package domain

import (
	"github.com/jmdavril/template/shop/app/event"
)

// product_created

var ProductCreatedEventType event.EventType = event.EventType{
	BC:   "catalog",
	Name: "product_created",
}

type ProductCreatedEvent struct {
	Sku string
}

func (ProductCreatedEvent) Type() event.EventType {
	return ProductCreatedEventType
}

func NewProductCreatedEvent(sku string) ProductCreatedEvent {
	return ProductCreatedEvent{
		Sku: sku,
	}
}
