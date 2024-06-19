package supplier

import "testing"

type testData struct {
	name      string
	email     string
	phone     string
	expectErr bool
}

func TestNewSupplier(t *testing.T) {
	data := testData{"Test Supplier", "test@email.com", "+5554997158012", false}

	supplier, err := New(data.name, data.email, data.phone)
	if (err != nil) != data.expectErr {
		t.Errorf("NewSupplier() error = %v, expectErr %v", err, data.expectErr)
	}

	if supplier.Name != data.name {
		t.Errorf("Supplier.Name = %v, want %v", supplier.Name, data.name)
	}

	if supplier.Email.String() != data.email {
		t.Errorf("Supplier.Email = %v, want %v", supplier.Email.String(), data.email)
	}
}

func TestNewSupplierInvalidName(t *testing.T) {
	data := testData{"te", "test@email.com", "+5554997158012", false}

	_, err := New(data.name, data.email, data.phone)
	if data.expectErr && err == nil {
		t.Errorf("NewSupplier() error = %v, expectErr %v", err, data.expectErr)
	}
}

func TestNewSupplierInvalidEmail(t *testing.T) {
	data := testData{"Test Supplier", "invalid-email", "+5554997158012", true}

	_, err := New(data.name, data.email, data.phone)
	if data.expectErr && err == nil {
		t.Errorf("NewSupplier() error = %v, expectErr %v", err, data.expectErr)
	}
}

func TestNewSupplierInvalidPhone(t *testing.T) {
	data := testData{"Test Supplier", "test@email.com", "invalid-phone", true}

	_, err := New(data.name, data.email, data.phone)
	if data.expectErr && err == nil {
		t.Errorf("NewSupplier() error = %v, expectErr %v", err, data.expectErr)
	}
}
