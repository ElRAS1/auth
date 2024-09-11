package server

import (
	"fmt"
	"net/mail"
	"regexp"
	"strings"

	"github.com/ELRAS1/auth/pkg/userApi"
)

const (
	numbers = `[0-9]+`
	letters = `[a-z]+`
	uppers  = `[A-Z]+`
	symbol  = `[!@#~$%^&*()+|_]+`
)

func (a *AppServer) CreateValidations(req *userApi.CreateRequest) error {
	const nm = "[CreateValidations]"

	if ok := strings.Compare(req.Password, req.PasswordConfirm); ok != 0 {
		return fmt.Errorf("%s passwords don't match", nm)
	}

	if err := req.Validate(); err != nil {
		return fmt.Errorf("%s %w", nm, err)
	}

	if err := CheckPasswordLever(req.Password); err != nil {
		return fmt.Errorf("%s %w", nm, err)
	}

	return nil
}

func (a *AppServer) UpdateValidation(req *userApi.UpdateRequest) error {
	const nm = "[UpdateValidation]"

	email, name := req.Email.GetValue(), req.Name.GetValue()
	if email != "" && name != "" {
		if err := req.Validate(); err != nil {
			return fmt.Errorf("%s %w", nm, err)
		}
		return nil

	} else if email == "" && name == "" {
		return fmt.Errorf("%s at least one field must be filled in", nm)

	}

	if email != "" {
		if _, err := mail.ParseAddress(req.Email.Value); err != nil {
			return fmt.Errorf("%s %w", nm, err)
		}
	}

	if name != "" {
		reg := `[A-Z][a-z]*`

		if b, err := regexp.MatchString(reg, name); !b || err != nil {
			return fmt.Errorf("%s %w", nm, err)
		}

		if len(name) <= 8 {
			return fmt.Errorf("%s %w", nm, fmt.Errorf("length name < 8"))
		}
	}

	return nil
}

func CheckPasswordLever(ps string) error {
	if len(ps) <= 8 {
		return fmt.Errorf("password len is < 8")
	}

	if b, err := regexp.MatchString(numbers, ps); !b || err != nil {
		return fmt.Errorf("there are no numbers")
	}

	if b, err := regexp.MatchString(letters, ps); !b || err != nil {
		return fmt.Errorf("there are no lowercase letters")
	}

	if b, err := regexp.MatchString(uppers, ps); !b || err != nil {
		return fmt.Errorf("there are no capital letters")
	}

	if b, err := regexp.MatchString(symbol, ps); !b || err != nil {
		return fmt.Errorf("password need symbol !@#~$%%^&*()+|_")
	}

	return nil
}
