package valueobject

import "testing"

func TestNewEmail(t *testing.T) {
	tests := []struct {
		email     string
		expectErr bool
	}{
		{"test@example.com", false},
		{"user@domain.co", false},
		{"invalid-email", true},
		{"@example.com", true},
		{"user@.com", true},
	}

	for _, tt := range tests {
		t.Run(tt.email, func(t *testing.T) {
			_, err := NewEmail(tt.email)
			if (err != nil) != tt.expectErr {
				t.Errorf("NewEmail() error = %v, expectErr %v", err, tt.expectErr)
			}
		})
	}
}

func TestEmail_String(t *testing.T) {
	emailStr := "test@example.com"
	email, err := NewEmail(emailStr)
	if err != nil {
		t.Fatalf("NewEmail() error = %v", err)
	}

	if got := email.String(); got != emailStr {
		t.Errorf("Email.String() = %v, want %v", got, emailStr)
	}
}
