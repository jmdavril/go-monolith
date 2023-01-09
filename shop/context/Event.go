package context

import (
	"github.com/jmdavril/pubsub"
)

const PRODUCT_CREATED_TYPE = "product_created"

type ProductCreatedEvent struct {
	Sku string
}

func (ProductCreatedEvent) Type() pubsub.EventType {
	return PRODUCT_CREATED_TYPE
}
