package commands

import (
	"github.com/felipemalacarne/back-prod-sup/internal/category/domain"
	"github.com/google/uuid"
)

type UpdateCategoryCommand struct {
	Name string    `json:"name"`
	ID   uuid.UUID `json:"id"`
}

type UpdateCategoryHandler struct {
	repository domain.CategoryRepository
}

func NewUpdateCategoryHandler(repository domain.CategoryRepository) *UpdateCategoryHandler {
    return &UpdateCategoryHandler{repository}
}

func (h *UpdateCategoryHandler) Handle(command UpdateCategoryCommand) (domain.Category, error) {
	category, err := h.repository.FindByID(command.ID)
	if err != nil {
		return domain.Category{}, err
	}
	category.Update(&command.Name)
	err = category.Validate()
	if err != nil {
		return domain.Category{}, err
	}
	return h.repository.Update(&category)
}
