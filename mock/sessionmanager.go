package mock

import (
	"fmt"
	"gitlab.com/projetAPI/ProjetAPI/service"
)

// SessionManagerMock mock session reader & writer
type SessionManagerMock struct {
	sessions map[string]*service.Session
}

var sessionMock *SessionManagerMock

// NewSessionMock create or retrieve SessionManagerMock
func NewSessionMock() *SessionManagerMock {
	if sessionMock == nil {
		sessionMock = &SessionManagerMock{}
		sessionMock.sessions = make(map[string]*service.Session)
		return sessionMock
	}
	return sessionMock
}

// LoadSession load session for userID
func (sm *SessionManagerMock) LoadSession(userID string) *service.Session {
	if sess, ok := sm.sessions[userID]; ok {
		sess.UserID = userID
		return sess
	}

	return nil
}

// LoadValidSessionFromJWT retrieve session for JWT token
func (sm *SessionManagerMock) LoadValidSessionFromJWT(tokenString string) (*service.Session, error) {
	// @TODO
	session := sessionMock.LoadSession(tokenString)
	if session == nil {
		return nil, fmt.Errorf("no session found")
	}
	return session, nil
}

// Write create session from userID and token
func (sm *SessionManagerMock) Write(userID string, tokenSecret string) bool {
	sm.sessions[userID] = &service.Session{Secret: tokenSecret }
	return true
}

// Destroy destroy the user session
func (sm *SessionManagerMock) Destroy(userID string) bool {
	if _, ok := sm.sessions[userID]; ok {
		delete(sm.sessions, userID)
		return true
	}
	return false
}
