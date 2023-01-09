package api

import (
	"github.com/jmdavril/template/shop/context/shop/domain"
	"github.com/samber/lo"
)

type ProductSalesListDto struct {
	Sales []ProductSalesDto `json:"sales"`
}

type ProductSalesDto struct {
	SKU        string  `json:"sku"`
	Quantity   int64   `json:"quantity"`
	TotalSales float64 `json:"totalSales"`
}

func NewProductSalesDto(ps domain.ProductSales) ProductSalesDto {
	return ProductSalesDto{
		SKU:        ps.SKU,
		Quantity:   ps.Quantity,
		TotalSales: ps.TotalSales,
	}
}

func NewProductSalesListDto(s []domain.ProductSales) ProductSalesListDto {
	return ProductSalesListDto{
		Sales: lo.Map(s, func(e domain.ProductSales, _ int) ProductSalesDto { return NewProductSalesDto(e) }),
	}
}
