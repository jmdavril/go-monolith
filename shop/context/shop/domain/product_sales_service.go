package domain

import (
	"github.com/jmdavril/template/shop/context/shop/data"
)

type ProductSalesService struct {
	productSalesRepo *data.ProductSalesRepo
}

func NewProudctSalesService(r *data.ProductSalesRepo) *ProductSalesService {
	return &ProductSalesService{
		productSalesRepo: r,
	}
}

func (s *ProductSalesService) InitProductSales(sku string) error {
	return s.productSalesRepo.InsertProductSales(sku)
}
