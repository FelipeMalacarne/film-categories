package persistence

import "github.com/felipemalacarne/back-prod-sup/internal/supplier/domain"

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
		Email:     supplier.Email.String(),
		Phone:     supplier.Phone.String(),
		CreatedAt: supplier.CreatedAt.Format("2006-01-02T15:04:05"),
		UpdatedAt: supplier.UpdatedAt.Format("2006-01-02T15:04:05"),
	}
}

func toSupplier(dynamoSupplier *dynamoSupplier) *domain.Supplier {
	return &domain.Supplier{}
}
