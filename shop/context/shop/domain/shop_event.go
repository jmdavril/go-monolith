package domain

import "github.com/jmdavril/template/shop/app/event"

// order_created

var OrderCreatedEventType event.EventType = event.EventType{
	BC:   "shop",
	Name: "order_created",
}

type OrderCreatedEvent struct {
	Order Order
}

func (OrderCreatedEvent) Type() event.EventType {
	return OrderCreatedEventType
}

func NewOrderCreatedEvent(order Order) OrderCreatedEvent {
	return OrderCreatedEvent{
		Order: order,
	}
}
