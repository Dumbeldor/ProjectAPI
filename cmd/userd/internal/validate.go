package internal

import (
	"fmt"
	"github.com/asaskevich/govalidator"
)

func validateEmail(email string) error {
	if !govalidator.IsEmail(email) {
		return fmt.Errorf("invalid email")
	}

	return nil
}

func validatePassword(password string) error {
	if len(password) < 8 {
		return fmt.Errorf("8 characters is the minimul password length")
	}

	return nil
}