package domain

import "github.com/google/uuid"

type FilmRepository interface {
	Create(film *Film) (Film, error)
	FindAll() ([]Film, error)
	FindByID(id uuid.UUID) (Film, error)
	Update(film *Film) (Film, error)
	Delete(id uuid.UUID) error
}
