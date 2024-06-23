package queries

import "github.com/felipemalacarne/back-prod-sup/internal/category/domain"

type ListCategoriesQuery struct{}

type ListCategoriesHandler struct {
	repository domain.CategoryRepository
}

func NewListCategoriesHandler(repository domain.CategoryRepository) *ListCategoriesHandler {
	return &ListCategoriesHandler{repository}
}

func (h *ListCategoriesHandler) Handle(query ListCategoriesQuery) ([]domain.Category, error) {
	return h.repository.FindAll()
}
