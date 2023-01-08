package domain

import (
	"github.com/google/uuid"
	"github.com/jmdavril/template/shop/context/catalog/data"
)

type Product struct {
	ID    uuid.UUID
	Sku   string
	Name  string
	Price float64
}

func (p Product) ProductEntiy() data.ProductEntity {
	return data.ProductEntity{
		ID:    p.ID,
		Sku:   p.Sku,
		Name:  p.Name,
		Price: p.Price,
	}
}

func NewProduct(dto data.ProductEntity) Product {
	return Product{
		ID:    dto.ID,
		Sku:   dto.Sku,
		Name:  dto.Name,
		Price: dto.Price,
	}
}
