package domain

import (
	"github.com/jmdavril/template/shop/context/shop/data"
	"github.com/samber/lo"
)

type ProductSalesService struct {
	productSalesRepo *data.ProductSalesRepo
}

func NewProductSalesService(r *data.ProductSalesRepo) *ProductSalesService {
	return &ProductSalesService{
		productSalesRepo: r,
	}
}

func (s *ProductSalesService) InitProductSales(sku string) error {
	return s.productSalesRepo.InsertProductSales(sku)
}

func (s *ProductSalesService) ReadProductSales() ([]ProductSales, error) {
	productSales, err := s.productSalesRepo.ProductSales()

	return lo.Map(productSales, func(e data.ProductSalesEntity, _ int) ProductSales { return NewProductSales(e) }), err
}
