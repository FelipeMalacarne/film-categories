package domain

import "testing"

type testData struct {
	name      string
	email     string
	phone     string
	expectErr bool
}

func TestNewSupplier(t *testing.T) {
	data := testData{"Test Supplier", "test@email.com", "+5554997158012", false}

	supplier, err := NewSupplier(data.name, data.email, data.phone)
	if (err != nil) != data.expectErr {
		t.Errorf("NewSupplier() error = %v, expectErr %v", err, data.expectErr)
	}

	if supplier.Name != data.name {
		t.Errorf("Supplier.Name = %v, want %v", supplier.Name, data.name)
	}

	if supplier.Email != data.email {
		t.Errorf("Supplier.Email = %v, want %v", supplier.Email, data.email)
	}
}

func TestNewSupplierInvalidName(t *testing.T) {
	data := testData{"te", "test@email.com", "+5554997158012", false}

	_, err := NewSupplier(data.name, data.email, data.phone)
	if data.expectErr && err == nil {
		t.Errorf("NewSupplier() error = %v, expectErr %v", err, data.expectErr)
	}
}

func TestNewSupplierInvalidEmail(t *testing.T) {
	data := testData{"Test Supplier", "invalid-email", "+5554997158012", true}

	_, err := NewSupplier(data.name, data.email, data.phone)
	if data.expectErr && err == nil {
		t.Errorf("NewSupplier() error = %v, expectErr %v", err, data.expectErr)
	}
}

func TestNewSupplierInvalidPhone(t *testing.T) {
	data := testData{"Test Supplier", "test@email.com", "invalid-phone", true}

	_, err := NewSupplier(data.name, data.email, data.phone)
	if data.expectErr && err == nil {
		t.Errorf("NewSupplier() error = %v, expectErr %v", err, data.expectErr)
	}
}
