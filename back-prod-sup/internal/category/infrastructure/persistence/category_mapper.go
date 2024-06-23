package persistence

import (
	"time"

	"github.com/felipemalacarne/back-prod-sup/internal/category/domain"
	"github.com/felipemalacarne/back-prod-sup/utils"
	"github.com/google/uuid"
)

type dynamoCategory struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

func toDynamoCategory(category *domain.Category) *dynamoCategory {
	return &dynamoCategory{
		ID:        category.ID.String(),
		Name:      category.Name,
		CreatedAt: category.CreatedAt.Format(time.RFC3339),
		UpdatedAt: category.UpdatedAt.Format(time.RFC3339),
	}
}

func toCategory(dc *dynamoCategory) *domain.Category {
	return &domain.Category{
		ID:        uuid.MustParse(dc.ID),
		Name:      dc.Name,
		CreatedAt: utils.ParseTime(dc.CreatedAt),
		UpdatedAt: utils.ParseTime(dc.UpdatedAt),
	}
}
