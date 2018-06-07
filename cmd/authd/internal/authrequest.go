package internal

import (
	"fmt"
)

// swagger:parameters login
type authRequest struct {
	// in: body
	Body struct {
		// User login
		// required: true
		Login string `json:"login,required"`
		// User password
		// required: true
		Password string `json:"password,required"`
	}
}

func (a *authRequest) Validate() error {
	if len(a.Body.Login) == 0 {
		return fmt.Errorf("login field is empty")
	}

	if len(a.Body.Password) == 0 {
		return fmt.Errorf("password field is empty")
	}

	return nil
}
