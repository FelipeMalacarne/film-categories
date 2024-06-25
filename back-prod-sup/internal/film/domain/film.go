package domain

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

type Film struct {
	ReleaseDate time.Time  `json:"release_date"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
	CategoryID  *uuid.UUID `json:"category_id,omitempty"`
	Name        string     `json:"name"`
	Description string     `json:"description"`
	Author      string     `json:"author"`
	Duration    uint16     `json:"duration"`
	ID          uuid.UUID  `json:"id"`
}

func NewFilm(name string, description string, duration uint16, releaseDate time.Time, author string) *Film {
	now := time.Now()
	return &Film{
		ID:          uuid.New(),
		Name:        name,
		Author:      author,
		Description: description,
		Duration:    duration,
		ReleaseDate: releaseDate,
		CreatedAt:   now,
		UpdatedAt:   now,
	}
}

func (f *Film) Update(name *string, description *string, duration *uint16, releaseDate *time.Time) {
	if name != nil {
		f.Name = *name
	}
	if description != nil {
		f.Description = *description
	}
	if duration != nil {
		f.Duration = *duration
	}
	if releaseDate != nil {
		f.ReleaseDate = *releaseDate
	}
	f.UpdatedAt = time.Now()
}

func (f *Film) SetCategoryID(categoryID uuid.UUID) {
    f.CategoryID = &categoryID
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

func (f *Film) validateAuthor() error {
	if len(f.Author) > 100 {
		return errors.New("author must have at most 100 characters")
	}
	if len(f.Author) == 0 {
		return errors.New("author cannot be empty")
	}
	return nil
}
