package domain

import (
	"github.com/google/uuid"
	"github.com/jmdavril/template/shop/context/catalog/data"
)

type ProductService struct {
	productRepo *data.ProductRepo
}

func NewProductService(pr *data.ProductRepo) *ProductService {
	return &ProductService{
		productRepo: pr,
	}
}

func (s *ProductService) CreateNewProduct(p Product) (uuid.UUID, error) {
	return s.productRepo.InsertProduct(p.ProductEntiy())
}

func (s *ProductService) UpdateProductInfo(p Product) error {
	return s.productRepo.UpdateProduct(p.ProductEntiy())
}

func (s *ProductService) ReadProduct(productId uuid.UUID) (Product, error) {
	dto, err := s.productRepo.SelectProduct(productId)
	return NewProduct(dto), err
}
