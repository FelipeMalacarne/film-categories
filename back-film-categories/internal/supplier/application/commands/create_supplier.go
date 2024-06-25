package commands

import "github.com/felipemalacarne/back-film-categories/internal/supplier/domain"

type CreateSupplierCommand struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Phone string `json:"phone"`
}

type CreateSupplierHandler struct {
	repository domain.SupplierRepository
}

func NewCreateSupplierHandler(repository domain.SupplierRepository) *CreateSupplierHandler {
	return &CreateSupplierHandler{repository}
}

func (h CreateSupplierHandler) Handle(command CreateSupplierCommand) (domain.Supplier, error) {
	supplier, err := domain.NewSupplier(command.Name, command.Email, command.Phone)
	if err != nil {
		return domain.Supplier{}, err
	}
	return h.repository.Create(supplier)
}

