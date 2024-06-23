package commands

import "github.com/felipemalacarne/back-prod-sup/internal/category/domain"

type CreateCategoryCommand struct {
	Name string `json:"name"`
}

type CreateCategoryHandler struct {
	repository domain.CategoryRepository
}

func (h *CreateCategoryHandler) Handle(command CreateCategoryCommand) (domain.Category, error) {
	category := domain.NewCategory(command.Name)
	err := category.Validate()
	if err != nil {
		return domain.Category{}, err
	}
	return h.repository.Create(category)
}
