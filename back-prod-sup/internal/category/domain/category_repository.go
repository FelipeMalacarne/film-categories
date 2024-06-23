package domain

import "github.com/google/uuid"

type CategoryRepository interface {
	Create(category *Category) (Category, error)
	FindAll() ([]Category, error)
	FindByID(id uuid.UUID) (Category, error)
    Update(category *Category) (Category, error)
    Delete(id uuid.UUID) error
}
