package domain

import (
	"github.com/google/uuid"
	"github.com/jmdavril/template/shop/context/shop/data"
)

type Customer struct {
	ID    uuid.UUID
	Email string
}

func (c Customer) CustomerEntity() data.CustomerEntity {
	return data.CustomerEntity{
		ID:    c.ID,
		Email: c.Email,
	}
}

func NewCustomer(dto data.CustomerEntity) Customer {
	return Customer{
		ID:    dto.ID,
		Email: dto.Email,
	}
}
