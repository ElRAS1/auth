package utils

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

// EncryptedPassw: hashing password: use bcrypt
func EncryptedPassw(passw string) (string, error) {
	const nm string = "[EncryptedPassw]"

	passwB, err := bcrypt.GenerateFromPassword([]byte(passw), bcrypt.MinCost)
	if err != nil {
		return "", fmt.Errorf("%s %w", nm, err)
	}

	if err = bcrypt.CompareHashAndPassword(passwB, []byte(passw)); err != nil {
		return "", fmt.Errorf("%s %w", nm, err)
	}

	return string(passwB), nil
}
