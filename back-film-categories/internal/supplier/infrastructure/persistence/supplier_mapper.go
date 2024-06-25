package persistence

import (
	"time"

	"github.com/felipemalacarne/back-film-categories/internal/supplier/domain"
	"github.com/felipemalacarne/back-film-categories/utils"
	"github.com/google/uuid"
)

type dynamoSupplier struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	Phone     string `json:"phone"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

func toDynamoSupplier(supplier *domain.Supplier) *dynamoSupplier {
	return &dynamoSupplier{
		ID:        supplier.ID.String(),
		Name:      supplier.Name,
		Email:     supplier.Email,
		Phone:     supplier.Phone,
		CreatedAt: supplier.CreatedAt.Format(time.RFC3339),
		UpdatedAt: supplier.UpdatedAt.Format(time.RFC3339),
	}
}

func toSupplier(ds *dynamoSupplier) *domain.Supplier {
	return &domain.Supplier{
		ID:        uuid.MustParse(ds.ID),
		Name:      ds.Name,
		Email:     ds.Email,
		Phone:     ds.Phone,
		CreatedAt: utils.ParseTime(ds.CreatedAt),
		UpdatedAt: utils.ParseTime(ds.UpdatedAt),
	}
}
