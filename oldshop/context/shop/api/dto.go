package api

import (
	"errors"

	"github.com/google/uuid"
	"github.com/jmdavril/template/shop/context/shop/domain"
	"github.com/samber/lo"
)

type CustomerRequest struct {
	Email string `json:"email"`
}

func (req CustomerRequest) customer() domain.Customer {
	return domain.Customer{
		Email: req.Email,
	}
}

type OrderRequest struct {
	CustomerId string        `json:"customer"`
	Items      []LineItemDto `json:"items"`
}

type LineItemDto struct {
	SKU       string  `json:"sku"`
	Quantity  int64   `json:"quantity"`
	UnitPrice float64 `json:"unitPrice"`
}

func (dto LineItemDto) lineItem() domain.LineItem {
	return domain.LineItem{
		Sku:       dto.SKU,
		Quantity:  dto.Quantity,
		UnitPrice: dto.UnitPrice,
	}
}

func (req OrderRequest) order() (domain.Order, error) {
	customerId, err := uuid.Parse(req.CustomerId)
	if err != nil {
		return domain.Order{}, errors.New("customer id must be a valid uuid")
	}

	return domain.Order{
		CustomerId: customerId,
		Items:      lo.Map(req.Items, func(x LineItemDto, _ int) domain.LineItem { return x.lineItem() }),
	}, nil
}
