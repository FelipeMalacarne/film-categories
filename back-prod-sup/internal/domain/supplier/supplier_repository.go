package supplier

import "github.com/google/uuid"

type SupplierRepository interface {
    Create(supplier *Supplier) (Supplier, error)
    FindAll() ([]Supplier, error)
    FindByID(id uuid.UUID) (Supplier, error)
    Update(supplier *Supplier) (Supplier, error)
    Delete(id uuid.UUID) error
}
