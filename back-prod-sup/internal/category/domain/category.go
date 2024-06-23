package domain

import (
	"time"

	"github.com/google/uuid"
)

type Category struct {
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Name      string    `json:"name"`
	ID        uuid.UUID `json:"id"`
}

func NewCategory(name string) *Category {
    now := time.Now()
    return &Category{
        ID:        uuid.New(),
        Name:      name,
        CreatedAt: now,
        UpdatedAt: now,
    }
}

func (c *Category) Update(name *string) {
    if name != nil {
        c.Name = *name
    }
    c.UpdatedAt = time.Now()
}
