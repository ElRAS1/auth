package server

import (
	"fmt"
	"net/mail"
	"strings"

	"github.com/ELRAS1/auth/pkg/userApi"
)

// TODO: исправить на более надежный вариант
// реализован самый простой вариант
func (a *AppServer) validations(req *userApi.Request) error {
	const nm = "[validations]"
	if req.Password == "" || req.PasswordConfirm == "" {
		return fmt.Errorf("%s password cannot be empty", nm)
	}
	if ok := strings.Compare(req.Password, req.PasswordConfirm); ok != 0 {
		return fmt.Errorf("%s passwords don't match", nm)
	}
	if len(req.Password) < 8 {
		return fmt.Errorf("%s password length must be greater than 7", nm)
	}
	if req.Name == "" {
		return fmt.Errorf("%s name cannot be empty", nm)
	}

	if _, err := mail.ParseAddress(req.Email); err != nil {
		return fmt.Errorf("%s %v", nm, err)
	}
	return nil
}
