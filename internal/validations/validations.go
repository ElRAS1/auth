package validations

import (
	"fmt"
	"net/mail"
	"regexp"
	"strings"

	"github.com/ELRAS1/auth/internal/models/user/model"
)

const (
	errMaxLengthName   = "name length exceeds maximum allowed limit of 50 characters"
	errMinLengthName   = "name length is less than 1"
	errContainsDigits  = "name cannot contain any digits"
	errContainsSpecial = "name cannot contain any special characters"
	errEmptyFields     = "at least one field must be filled"
	errNotMatch        = "passwords don't match"
)

const (
	maxNameLength = 50
	minNameLength = 1
)

func CheckCreate(req *model.CreateRequest) error {
	if err := CheckName(req.Name); err != nil {
		return err
	}

	if err := CheckEmail(req.Email); err != nil {
		return err
	}

	if strings.Compare(req.Password, req.PasswordConfirm) != 0 {
		return fmt.Errorf("%s", errNotMatch)
	}

	if err := CheckPassword(req.Password); err != nil {
		return err
	}

	return nil
}

func CheckUpdate(req *model.UpdateRequest) error {
	if req.Name == "" && req.Email == "" {
		return fmt.Errorf("%s", errEmptyFields)
	}

	if req.Name != "" {
		if err := CheckName(req.Name); err != nil {
			return err
		}
	}

	if req.Email != "" {
		if err := CheckEmail(req.Email); err != nil {
			return err
		}
	}

	return nil
}

func CheckEmail(email string) error {
	_, err := mail.ParseAddress(email)
	if err != nil {
		return fmt.Errorf("email error: %w", err)
	}

	return nil
}

func CheckName(name string) error {
	if len(name) > maxNameLength {
		return fmt.Errorf("%s", errMaxLengthName)
	}

	if len(name) < minNameLength {
		return fmt.Errorf("%s", errMinLengthName)
	}

	if regexp.MustCompile(`\d`).MatchString(name) {
		return fmt.Errorf("%s", errContainsDigits)
	}

	if regexp.MustCompile(`[^a-zA-Z\s]`).MatchString(name) {
		return fmt.Errorf("%s", errContainsSpecial)
	}

	return nil
}
