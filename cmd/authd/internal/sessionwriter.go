package internal

import (
	"fmt"
	"github.com/go-redis/redis"
	"time"
	"gitlab.com/projetAPI/ProjetAPI/service"
)

type writer struct {
	redisClient *redis.Client
}

func newWriter(redisConfig service.RedisConfig) *writer {
	writer := &writer{}
	writer.redisClient = redis.NewClient(&redis.Options{
		Addr:       fmt.Sprintf("%s:%d", redisConfig.Host, redisConfig.Port),
		Password:   redisConfig.Password,
		DB:         redisConfig.DatabaseID,
		MaxRetries: redisConfig.MaxRetries,
	})
	return writer
}

func (w *writer) Write(userID string, tokenSecret string) bool {
	// Create session object
	sessionObj := service.Session{Secret: tokenSecret}

	//
	jsonSession, err := sessionObj.Serialize()
	if err != nil {
		app.Log.Errorf("Failed to write session JSON for user ID %s. Error was: %s", userID, err)
		return false
	}

	redisErr := w.redisClient.Set(fmt.Sprintf("session-%s", userID), jsonSession,
		time.Duration(gconfig.Session.Duration)*time.Second)
	if redisErr.Err() != nil {
		app.Log.Errorf("Failed to write session for user ID %s. Error was: %s", userID, redisErr.Err())
		return false
	}

	return true
}

func (w *writer) Destroy(userID string) bool {
	redisErr := w.redisClient.Del(fmt.Sprintf("session-%s", userID))
	if redisErr.Err() != nil {
		app.Log.Errorf("Failed to destroy session for user ID %s. Error was: %s", userID, redisErr.Err())
		return false
	}

	return true
}
