package domain

import (
	"errors"
	"regexp"
	"time"

	"github.com/google/uuid"
)

type Supplier struct {
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Phone     string    `json:"phone"`
	ID        uuid.UUID `json:"id"`
}

func NewSupplier(name string, email string, phone string) (*Supplier, error) {
    s := &Supplier{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      name,
        Email:     email,
        Phone:     phone,
	}

    err := s.Validate()
    if err != nil {
        return nil, err
    }

    return s, nil
}

func (s *Supplier) Validate() error {
	err := s.validateName()
	if err != nil {
		return err
	}

	err = s.validateEmail()
	if err != nil {
		return err
	}

	err = s.validatePhone()
	if err != nil {
		return err
	}

	return nil
}

func (s *Supplier) validateName() error {
	if len(s.Name) > 100 {
		return errors.New("name must have at most 100 characters")
	}

	if len(s.Name) < 3 {
		return errors.New("name must have at least 3 characters")
	}

	return nil
}

func (s *Supplier) validateEmail() error {
	if s.Email == "" {
		return errors.New("email is required")
	}
	const emailRegex = `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`

	re := regexp.MustCompile(emailRegex)
	if !re.MatchString(s.Email) {
		return errors.New("invalid email address")
	}
	return nil
}

func (s *Supplier) validatePhone() error {
	if s.Phone == "" {
		return errors.New("phone is required")
	}

	const phoneRegex = `^\+[1-9]\d{1,14}$`

	re := regexp.MustCompile(phoneRegex)
	if !re.MatchString(s.Phone) {
		return errors.New("invalid phone number")
	}

	return nil
}
