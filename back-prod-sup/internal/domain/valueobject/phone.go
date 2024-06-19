package valueobject

import "regexp"

type Phone struct {
	number string
}

func NewPhone(number string) (Phone, error) {
	if !isValidPhone(number) {
		return Phone{}, ErrInvalidPhone
	}
	return Phone{number: number}, nil
}

func (p Phone) String() string {
	return p.number
}

func isValidPhone(phone string) bool {
	// Simple regex for phone validation
	const phoneRegex = `^\+?[1-9]\d{1,14}$`
	re := regexp.MustCompile(phoneRegex)
	return re.MatchString(phone)
}
