package command

import "github.com/felipemalacarne/back-prod-sup/internal/supplier/domain"

type CreateSupplierCommand struct {
	name  string
	email string
	phone string
}

type CreateSupplierHandler struct {
	repository domain.SupplierRepository
}

func NewCreateSupplierHandler(repository domain.SupplierRepository) *CreateSupplierHandler {
    return &CreateSupplierHandler{repository}
}

func (h CreateSupplierHandler) Handle(command CreateSupplierCommand) (domain.Supplier, error) {
    supplier, err := domain.NewSupplier(command.name, command.email, command.phone)
    if err != nil {
        return domain.Supplier{}, err
    }
    return h.repository.Create(supplier)
}
