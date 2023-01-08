package domain

import (
	"github.com/google/uuid"
	"github.com/jmdavril/template/shop/context/shop/data"
)

type CustomerService struct {
	customerRepo     *data.CustomerRepo
	orderRepo        *data.OrderRepo
	productSalesRepo *data.ProductSalesRepo
}

func NewCustomerService(cr *data.CustomerRepo, or *data.OrderRepo, psr *data.ProductSalesRepo) *CustomerService {
	return &CustomerService{
		customerRepo:     cr,
		orderRepo:        or,
		productSalesRepo: psr,
	}
}

func (s *CustomerService) CreateNewCustomer(c Customer) (uuid.UUID, error) {
	return s.customerRepo.InsertCustomer(c.CustomerEntity())
}

func (s *CustomerService) ReadCustomer(customerId uuid.UUID) (Customer, error) {
	dto, err := s.customerRepo.SelectCustomer(customerId)
	return NewCustomer(dto), err
}

func (s *CustomerService) CreateNewOrder(o Order) (uuid.UUID, error) {
	entity := o.OrderEntity()

	orderId, err := s.orderRepo.InsertOrder(entity)
	if err != nil {
		return uuid.Nil, err
	}

	err = s.productSalesRepo.UpdateAllProductSales(entity)

	return orderId, err
}
