package domain

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

type Film struct {
	ReleaseDate time.Time `json:"release_date"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Duration    uint16    `json:"duration"`
	ID          uuid.UUID `json:"id"`
}

func NewFilm(name string, description string, duration uint16, releaseDate time.Time) *Film {
	now := time.Now()
	return &Film{
		ID:          uuid.New(),
		Name:        name,
		Description: description,
		Duration:    duration,
		ReleaseDate: releaseDate,
		CreatedAt:   now,
		UpdatedAt:   now,
	}
}

func (f *Film) Validate() error {
	err := f.validateName()
	if err != nil {
		return err
	}
	err = f.validateDescription()
	if err != nil {
		return err
	}
	err = f.validateDuration()
	if err != nil {
		return err
	}
	err = f.validateReleaseDate()
	if err != nil {
		return err
	}
	return nil
}

func (f *Film) validateName() error {
	if len(f.Name) > 100 {
		return errors.New("name must have at most 100 characters")
	}
	if len(f.Name) == 0 {
		return errors.New("name cannot be empty")
	}

	return nil
}

func (f *Film) validateDescription() error {
	if len(f.Description) > 255 {
		return errors.New("description must have at most 255 characters")
	}
	if len(f.Description) == 0 {
		return errors.New("description cannot be empty")
	}

	return nil
}

func (f *Film) validateDuration() error {
    if f.Duration <= 0 {
        return errors.New("duration must be greater than 0")
    }
	if f.Duration > 60000 {
		return errors.New("duration must be less than 60000")
	}

	return nil
}

func (f *Film) validateReleaseDate() error {
	if f.ReleaseDate.After(time.Now()) {
		return errors.New("release date must be in the past")
	}
	return nil
}
