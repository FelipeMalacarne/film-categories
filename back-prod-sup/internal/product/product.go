package domain

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

var (
	ErrNameIsRequired             = errors.New("name is required")
	ErrPriceMustBeGreaterThanZero = errors.New("price must be greater than zero")
)

type Product struct {
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Price       int       `json:"price"`
	ID          uuid.UUID `json:"id"`
}

func NewProduct(name, description string, price int) *Product {
	return &Product{
        ID:          uuid.New(),
        CreatedAt:   time.Now(),
        UpdatedAt:   time.Now(),
        Name:        name,
        Description: description,
        Price:       price,
	}
}

func (p Product) Validate() error {
	if p.Name == "" {
		return ErrNameIsRequired
	}
	if p.Price <= 0 {
		return ErrPriceMustBeGreaterThanZero
	}
	return nil
}
