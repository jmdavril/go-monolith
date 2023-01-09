package domain

import (
	"github.com/google/uuid"
	"github.com/jmdavril/pubsub"
	"github.com/jmdavril/template/shop/context"
	"github.com/jmdavril/template/shop/context/catalog/data"
)

type ProductService struct {
	productRepo *data.ProductRepo
	pub         pubsub.Pub
}

func NewProductService(pr *data.ProductRepo, p *pubsub.PubSub) *ProductService {
	return &ProductService{
		productRepo: pr,
		pub:         p,
	}
}

func (s *ProductService) CreateNewProduct(p Product) (uuid.UUID, error) {
	id, err := s.productRepo.InsertProduct(p.ProductEntiy())

	s.pub.Publish(context.ProductCreatedEvent{Sku: p.Sku})

	return id, err
}

func (s *ProductService) UpdateProductInfo(p Product) error {
	return s.productRepo.UpdateProduct(p.ProductEntiy())
}

func (s *ProductService) ReadProduct(productId uuid.UUID) (Product, error) {
	dto, err := s.productRepo.SelectProduct(productId)
	return NewProduct(dto), err
}
