package internal

import (
	"crypto/sha512"
	"encoding/hex"
	"github.com/dgrijalva/jwt-go"
	"github.com/satori/go.uuid"
	"time"
)

func createJWT(rinfo authRequest, linfo *loginInfos) (string, string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": linfo.UserID,
		"exp": time.Now().Unix() + gconfig.Session.Duration,
	})

	// JWT secret is uuid + @ + current time string
	hasher := sha512.New()
	hasher.Write([]byte(uuid.NewV4().String() + "@" + time.Now().String()))

	secret := hasher.Sum(nil)

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString(secret)
	if err != nil {
		app.Log.Errorf("Unable to create JWT token for user %s. Error was: %s", rinfo.Body.Login, err.Error())
		return "", "", err
	}

	return tokenString, hex.EncodeToString(secret), nil
}
