package domain

import "github.com/jmdavril/template/shop/context/shop/data"

type ProductSales struct {
	SKU        string
	Quantity   int64
	TotalSales float64
}

func NewProductSales(entity data.ProductSalesEntity) ProductSales {
	return ProductSales{
		SKU:        entity.Sku,
		Quantity:   entity.Quantity,
		TotalSales: entity.TotalSales,
	}
}
