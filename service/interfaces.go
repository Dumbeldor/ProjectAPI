package service

// ReaderInterface Authentication reader interface
type ReaderInterface interface {
	LoadSession(userID string) *Session
	LoadValidSessionFromJWT(tokenString string) (*Session, error)
}

// WriterInterface Authenticate writer interface
type WriterInterface interface {
	Write(userID string, tokenSecret string) bool
	Destroy(userID string) bool
}
