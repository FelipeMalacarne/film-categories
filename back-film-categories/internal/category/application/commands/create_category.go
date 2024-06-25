package commands

import "github.com/felipemalacarne/back-film-categories/internal/category/domain"

type CreateCategoryCommand struct {
	Name string `json:"name"`
}

type CreateCategoryHandler struct {
	repository domain.CategoryRepository
}

func NewCreateCategoryHandler(repository domain.CategoryRepository) *CreateCategoryHandler {
	return &CreateCategoryHandler{repository}
}

func (h *CreateCategoryHandler) Handle(command CreateCategoryCommand) (domain.Category, error) {
	category := domain.NewCategory(command.Name)
	err := category.Validate()
	if err != nil {
		return domain.Category{}, err
	}
	return h.repository.Create(category)
}
