package utils

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func EncryptedPassword(password string) (string, error) {
	const nm string = "[EncryptedPassword]"

	passwB, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
	if err != nil {
		return "", fmt.Errorf("%s %w", nm, err)
	}

	return string(passwB), nil
}
