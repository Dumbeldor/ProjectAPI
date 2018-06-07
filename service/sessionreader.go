package service

import (
	"encoding/hex"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/go-redis/redis"
)

// RedisConfig configuration object to connect to redis
type RedisConfig struct {
	Host       string `yaml:"host"`
	Port       int16  `yaml:"port"`
	Password   string `yaml:"password"`
	DatabaseID int    `yaml:"database"`
	MaxRetries int    `yaml:"max_retries"`
}

// Reader session reader
type Reader struct {
	redisClient *redis.Client
}

// NewReader create a reader object bound to redis storage
func NewReader(redisConfig RedisConfig) *Reader {
	reader := &Reader{}
	reader.redisClient = redis.NewClient(&redis.Options{
		Addr:       fmt.Sprintf("%s:%d", redisConfig.Host, redisConfig.Port),
		Password:   redisConfig.Password,
		DB:         redisConfig.DatabaseID,
		MaxRetries: redisConfig.MaxRetries,
	})
	return reader
}

// LoadSession load session for userID from redis
// Returns session object if a valid session exists
func (r *Reader) LoadSession(userID string) *Session {
	val, err := r.redisClient.Get(fmt.Sprintf("session-%s", userID)).Result()
	if err == redis.Nil {
		return nil
	}

	session := &Session{}
	err = session.read([]byte(val))
	if err != nil {
		return nil
	}

	// This should be set after reading because it will be emptied by Unmarshal
	session.UserID = userID
	return session
}

// LoadValidSessionFromJWT decode session from JWT
// return error and nil session on failure
func (r *Reader) LoadValidSessionFromJWT(tokenString string) (*Session, error) {
	var session *Session
	session = nil

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		secret := ""

		if claims, ok := token.Claims.(jwt.MapClaims); ok {
			if _, ok := claims["sub"]; !ok {
				return nil, fmt.Errorf("sub claim not found in JWT")
			}

			if _, ok := claims["exp"]; !ok {
				return nil, fmt.Errorf("exp claim not found in JWT")
			}

			sub := fmt.Sprintf("%s", claims["sub"])
			session = r.LoadSession(sub)
			if session == nil {
				return nil, fmt.Errorf("no session found for user ID %s", sub)
			}
			secret = session.Secret
		}

		decodedSecret, err := hex.DecodeString(secret)
		if err != nil {
			return nil, fmt.Errorf("unable to decode secret")
		}

		return decodedSecret, nil
	})

	// If error or token is invalid, return nil session & error
	if err != nil || !token.Valid {
		return session, err
	}

	// Return loaded valid session
	return session, nil
}
