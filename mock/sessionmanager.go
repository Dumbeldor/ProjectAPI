package mock

import (
	"gitlab.com/projetAPI/auth"
	"fmt"
)

// SessionManagerMock mock session reader & writer
type SessionManagerMock struct {
	sessions map[string]*auth.Session
}

var sessionMock *SessionManagerMock

// NewSessionMock create or retrieve SessionManagerMock
func NewSessionMock() *SessionManagerMock {
	if sessionMock == nil {
		sessionMock = &SessionManagerMock{}
		sessionMock.sessions = make(map[string]*auth.Session)
		return sessionMock
	}
	return sessionMock
}

// LoadSession load session for userID
func (sm *SessionManagerMock) LoadSession(userID string) *auth.Session {
	if sess, ok := sm.sessions[userID]; ok {
		sess.UserID = userID
		return sess
	}

	return nil
}

// LoadValidSessionFromJWT retrieve session for JWT token
func (sm *SessionManagerMock) LoadValidSessionFromJWT(tokenString string) (*auth.Session, error) {
	// @TODO
	session := sessionMock.LoadSession(tokenString)
	if session == nil {
		return nil, fmt.Errorf("no session found")
	}
	return session, nil
}

// Write create session from userID and token
func (sm *SessionManagerMock) Write(userID string, tokenSecret string) bool {
	sm.sessions[userID] = &auth.Session{Secret: tokenSecret }
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
