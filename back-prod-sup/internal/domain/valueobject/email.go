package valueobject

import (
	"errors"
	"regexp"
)

var (
	ErrInvalidEmail = errors.New("invalid email address")
	ErrInvalidPhone = errors.New("invalid phone number")
)

type Email struct {
	address string
}

func NewEmail(address string) (Email, error) {
	if !isValidEmail(address) {
		return Email{}, ErrInvalidEmail
	}
	return Email{address: address}, nil
}

func (e Email) String() string {
	return e.address
}

func isValidEmail(email string) bool {
	const emailRegex = `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	re := regexp.MustCompile(emailRegex)
	return re.MatchString(email)
}

