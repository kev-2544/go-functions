package validation

import (
	"errors"
	"strings"
)

func ValidatePhoneNumber(phone string) error {
	if phone == "" {
		return errors.New("phone number cannot be empty")
	}
	for _, ch := range phone {
		if ch < '0' || ch > '9' {
			return errors.New("phone number can only contains digits(no spaces or letters)")
		}
	}
	if !strings.HasPrefix(phone, "254") {
		return errors.New("phone no must start with 254")
	}
	if len(phone) != 12 {
		return errors.New("phone no must be 12 digits")
	}

	return nil
}
