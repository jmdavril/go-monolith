package data

import "github.com/google/uuid"

type CustomerEntity struct {
	ID    uuid.UUID
	Email string
}
