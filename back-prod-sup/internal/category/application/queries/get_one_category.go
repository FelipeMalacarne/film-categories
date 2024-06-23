package queries

import (
	"github.com/felipemalacarne/back-prod-sup/internal/category/domain"
	"github.com/google/uuid"
)

type GetOneCategoryQuery struct {
	ID uuid.UUID `json:"id"`
}

type GetOneCategoryHandler struct {
	repository domain.CategoryRepository
}

func NewGetOneCategoryHandler(repository domain.CategoryRepository) *GetOneCategoryHandler {
    return &GetOneCategoryHandler{repository}
}

func (h *GetOneCategoryHandler) Handle(query GetOneCategoryQuery) (domain.Category, error) {
    return h.repository.FindByID(query.ID)
}
