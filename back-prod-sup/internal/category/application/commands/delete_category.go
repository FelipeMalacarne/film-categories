package commands

import (
	"github.com/felipemalacarne/back-prod-sup/internal/category/domain"
	"github.com/google/uuid"
)

type DeleteCategoryCommand struct {
	ID uuid.UUID `json:"id"`
}

type DeleteCategoryHandler struct {
	repository domain.CategoryRepository
}

func NewDeleteCategoryHandler(repository domain.CategoryRepository) *DeleteCategoryHandler {
	return &DeleteCategoryHandler{repository}
}

func (h *DeleteCategoryHandler) Handle(command DeleteCategoryCommand) error {
	return h.repository.Delete(command.ID)
}
