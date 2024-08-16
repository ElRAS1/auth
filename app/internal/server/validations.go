package server

import (
	"fmt"
	"net/mail"
	"regexp"
	"strings"

	"github.com/ELRAS1/auth/pkg/userApi"
)

func (a *AppServer) CreateValidations(req *userApi.CreateRequest) error {
	const nm = "[CreateValidations]"
	if err := req.Validate(); err != nil {
		return fmt.Errorf("%s %v", nm, err.Error())
	}

	if ok := strings.Compare(req.Password, req.PasswordConfirm); ok != 0 {
		return fmt.Errorf("%s passwords don't match", nm)
	}
	if err := CheckPasswordLever(req.Password); err != nil {
		return fmt.Errorf("the password must contain at least one capital letter, a number and at least 8 characters")
	}
	return nil
}

func (a *AppServer) UpdateValidation(req *userApi.UpdateRequest) error {
	const nm = "[UpdateValidation]"
	email, name := req.Email.GetValue(), req.Name.GetValue()
	if email != "" && name != "" {
		if err := req.Validate(); err != nil {
			return fmt.Errorf("%s %v", nm, err.Error())
		}
		return nil
	}
	if email != "" {
		if _, err := mail.ParseAddress(req.Email.String()); err != nil {
			return fmt.Errorf("%s %v", nm, err.Error())
		}
	}
	if name != "" {
		reg := `[A-Z][a-z]*`
		if b, err := regexp.MatchString(reg, name); !b || err != nil {
			return fmt.Errorf("%s %v", nm, err.Error())
		}
		if len(name) <= 8 {
			return fmt.Errorf("%s %v", nm, fmt.Errorf("length name < 8"))
		}
	}
	return nil
}

func CheckPasswordLever(ps string) error {
	if len(ps) <= 8 {
		return fmt.Errorf("password len is < 8")
	}
	num := `[0-9]*`
	a_z := `[a-z]*`
	A_Z := `[A-Z]*`
	symbol := `[!@#~$%^&*()+|_]*`
	if b, err := regexp.MatchString(num, ps); !b || err != nil {
		return fmt.Errorf("password need num :%v", err)
	}
	if b, err := regexp.MatchString(a_z, ps); !b || err != nil {
		return fmt.Errorf("password need a_z :%v", err)
	}
	if b, err := regexp.MatchString(A_Z, ps); !b || err != nil {
		return fmt.Errorf("password need A_Z :%v", err)
	}
	if b, err := regexp.MatchString(symbol, ps); !b || err != nil {
		return fmt.Errorf("password need symbol :%v", err)
	}
	return nil
}
