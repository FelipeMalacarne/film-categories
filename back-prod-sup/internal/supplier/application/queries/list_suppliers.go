package queries

import "github.com/felipemalacarne/back-prod-sup/internal/supplier/domain"

type ListSuppliersQuery struct {}

type ListSuppliersHandler struct {
    repository domain.SupplierRepository
}

func NewListSuppliersHandler(repository domain.SupplierRepository) *ListSuppliersHandler {
    return &ListSuppliersHandler{repository}
}

func (h ListSuppliersHandler) Handle(query ListSuppliersQuery) ([]domain.Supplier, error) {
    return h.repository.FindAll()
}
