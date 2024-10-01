package validations

import (
	"fmt"
	"unicode"
)

const (
	errNoCapitalLetter = "password must contain at least one capital letter"
	errNoSmallLetter   = "password must contain at least one small letter"
	errNoDigit         = "password must contain at least one digit"
	errNoSpecialChar   = "password must contain at least one of these characters '!@#$%^&*()_+-={}:<>?,./;'"

	errPasswordTooShort = "password length is less than 8"
	errPasswordTooLong  = "password length should not exceed 50 characters"
)

const (
	maxPasswordLength = 50
	minPasswordLength = 8
)

func CheckPassword(password string) error {
	if len(password) < minPasswordLength {
		return fmt.Errorf("password error: %s", errPasswordTooShort)
	}

	if len(password) > maxPasswordLength {
		return fmt.Errorf("password error: %s", errPasswordTooLong)
	}

	var (
		hasUpper   = false
		hasLower   = false
		hasNumber  = false
		hasSpecial = false
	)

	for _, sym := range password {

		if hasUpper && hasLower && hasNumber && hasSpecial {
			return nil
		}

		switch {
		case unicode.IsUpper(sym):
			hasUpper = true
		case unicode.IsLower(sym):
			hasLower = true
		case unicode.IsDigit(sym):
			hasNumber = true
		case unicode.IsSymbol(sym) || unicode.IsPunct(sym):
			hasSpecial = true

		}
	}

	if !hasUpper {
		return fmt.Errorf("password error: %s", errNoCapitalLetter)
	}
	if !hasLower {
		return fmt.Errorf("password error: %s", errNoSmallLetter)
	}
	if !hasNumber {
		return fmt.Errorf("password error: %s", errNoDigit)
	}
	if !hasSpecial {
		return fmt.Errorf("password error: %s", errNoSpecialChar)
	}

	return nil
}
