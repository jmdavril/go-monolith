package domain

import (
	"github.com/google/uuid"
	"github.com/jmdavril/template/shop/app/utils"
	"github.com/jmdavril/template/shop/context/shop/data"
)

var logger = utils.LoggerWith(utils.Shop, utils.Domain)

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

	orderID, err := s.orderRepo.InsertOrder(entity)
	if err != nil {
		return uuid.Nil, err
	}

	err = s.productSalesRepo.UpdateAllProductSales(entity)

	logger.Info().
		Str("Op", "orderCreate").
		Str("OrderID", orderID.String()).
		Msg("Updated product sales entities for order")

	return orderID, err
}
