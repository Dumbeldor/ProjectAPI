package service

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/dgrijalva/jwt-go"
)

// Session The user session object
type Session struct {
	Secret string `json:"jwt_secret"`
	UserID string
}

// Serialize transforms session object to a json string
func (s *Session) Serialize() ([]byte, error) {
	jsonSession, err := json.Marshal(s)
	if err != nil {
		return nil, err
	}

	return jsonSession, nil
}

func (s *Session) read(data []byte) error {
	return json.Unmarshal(data, s)
}

// Validate valite JWT tokenString
// Returns error when JWT is invalid
func (s *Session) Validate(tokenString string) error {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		decodedSecret, err := hex.DecodeString(s.Secret)
		if err != nil {
			return nil, fmt.Errorf("unable to decode secret")
		}

		return decodedSecret, nil
	})

	if err != nil {
		return err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		if _, ok := claims["sub"]; !ok {
			return fmt.Errorf("sub claim not found in JWT")
		}

		if _, ok := claims["exp"]; !ok {
			return fmt.Errorf("exp claim not found in JWT")
		}

		return nil
	}

	return fmt.Errorf("invalid JWT")
}
