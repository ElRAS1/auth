package validations

import (
	"fmt"
	"testing"
)

func TestPassword(t *testing.T) {
	t.Parallel()

	t.Run("TestCheckingLength", func(t *testing.T) {

		longPassw := "Ee123456?1111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111"
		if err := CheckPassword(longPassw); err.Error() != fmt.Sprintf("password error: %v", errPasswordTooLong) {
			t.Errorf("failed test to long password: err %v", err)
		}

		shortPassw := "Ee1234"
		if err := CheckPassword(shortPassw); err.Error() != fmt.Sprintf("password error: %v", errPasswordTooShort) {
			t.Errorf("failed test to short password: err %v", err)
		}

	})

	t.Run("TestNoSpecialSymbol", func(t *testing.T) {
		passw := "Ee12345678"
		if err := CheckPassword(passw); err.Error() != fmt.Sprintf("password error: %v", errNoSpecialChar) {
			t.Errorf("failed test to no special symbol: err %v", err)
		}
	})

	t.Run("TestNoDigit", func(t *testing.T) {
		passw := "Eerrrrrr?"
		if err := CheckPassword(passw); err.Error() != fmt.Sprintf("password error: %v", errNoDigit) {
			t.Errorf("failed test to no digit: err %v", err)
		}
	})

	t.Run("TestNoCapitalLetter", func(t *testing.T) {
		passw := "eerrrrrr9?"
		if err := CheckPassword(passw); err.Error() != fmt.Sprintf("password error: %v", errNoCapitalLetter) {
			t.Errorf("failed test to no capital letter: err %v", err)
		}
	})

	t.Run("TestNoSmallLetter", func(t *testing.T) {
		passw := "EEEEEEEE9?"
		if err := CheckPassword(passw); err.Error() != fmt.Sprintf("password error: %v", errNoSmallLetter) {
			t.Errorf("failed test to no small letter: err %v", err)
		}
	})

}
