package valueobject

import "testing"

func TestNewPhone(t *testing.T) {
	tests := []struct {
		phone      string
		expectErr  bool
	}{
		{"+1234567890", false},
		{"1234567890", false},
		{"+1-234-567-890", true},
		{"invalid-phone", true},
		{"", true},
	}

	for _, tt := range tests {
		t.Run(tt.phone, func(t *testing.T) {
			_, err := NewPhone(tt.phone)
			if (err != nil) != tt.expectErr {
				t.Errorf("NewPhone() error = %v, expectErr %v", err, tt.expectErr)
			}
		})
	}
}

func TestPhone_String(t *testing.T) {
	phoneStr := "+1234567890"
	phone, err := NewPhone(phoneStr)
	if err != nil {
		t.Fatalf("NewPhone() error = %v", err)
	}

	if got := phone.String(); got != phoneStr {
		t.Errorf("Phone.String() = %v, want %v", got, phoneStr)
	}
}
