package domain

import (
	"github.com/jmdavril/template/shop/app/event"
	"github.com/jmdavril/template/shop/context/catalog/domain"
)

type EventHandler struct {
	productSalesService *ProductSalesService
}

func NewEventHandler(pss *ProductSalesService) *EventHandler {
	return &EventHandler{
		productSalesService: pss,
	}
}

func (h *EventHandler) StartSubscriptions(p *event.Publisher) {
	p.Subscribe(domain.ProductCreatedEventType, h.handleProductCreatedEvent)
}

func (h *EventHandler) handleProductCreatedEvent(e event.Event) error {
	pce, _ := e.(domain.ProductCreatedEvent)
	err := h.productSalesService.InitProductSales(pce.Sku)
	return err
}
