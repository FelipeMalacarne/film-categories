package supplier

import (
	"errors"
	"time"

	"github.com/felipemalacarne/back-prod-sup/internal/domain/valueobject"
	"github.com/google/uuid"
)

type Supplier struct {
	CreatedAt time.Time         `json:"created_at"`
	UpdatedAt time.Time         `json:"updated_at"`
	Name      string            `json:"name"`
	Email     valueobject.Email `json:"email"`
	Phone     valueobject.Phone `json:"phone"`
	ID        uuid.UUID         `json:"id"`
}

func New(name string, email string, phone string) (*Supplier, error) {
	validatedEmail, err := valueobject.NewEmail(email)
	if err != nil {
		return nil, err
	}

	validatedPhone, err := valueobject.NewPhone(phone)
	if err != nil {
		return nil, err
	}

	err = validateName(name)
	if err != nil {
		return nil, err
	}

	return &Supplier{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      name,
		Email:     validatedEmail,
		Phone:     validatedPhone,
	}, nil
}

func validateName(name string) error {
	if len(name) > 100 {
		return errors.New("name must have at most 100 characters")
	}

	if len(name) < 3 {
		return errors.New("name must have at least 3 characters")
	}

	return nil
}
