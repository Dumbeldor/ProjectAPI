package internal

import (
	"fmt"
	"regexp"
)

// swagger:parameters registerRequest
type registerRequest struct {
	Login    string `json:"login"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

func (r *registerRequest) Validate() error {
	// Login validate
	if len(r.Login) < 3 {
		return fmt.Errorf("3 characters is the minimum login length")
	}
	if len(r.Login) > 20 {
		return fmt.Errorf("20 characters is the maximum login length")
	}

	if err := validateEmail(r.Email); err != nil {
		return err
	}

	var validLogin = regexp.MustCompile(`^[\w-.]{2,}$`)

	if !validLogin.MatchString(r.Login) {
		return fmt.Errorf("invalid login, allowed characters are a-z A-Z 0-9")
	}

	return validatePassword(r.Password)
}
