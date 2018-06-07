package service

import "testing"

func TestSession_Validate(t *testing.T) {
	s := Session{"66777063c115c06cc4d4d5d5a5b0603b63c3fc029af4d19720fe0d4ebb3e4e247edc09a904a9a31eecb7dc4d92dd4b70425dfae99a8716efbba61bf47051d922", ""}
	tokenStr := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjQ2NjQ3ODgyMzcsInN1YiI6ImFlNmQwOTE3LWI0MDItNGUyZi05ZDM5LTZkODEyYWQwMTljNiJ9.OZp5YU3nSp0ePuMoxW1pkgyWfnQT9tr5EpWqZb3Eh4g"
	if s.Validate(tokenStr) != nil {
		t.Error("failed to validate JWT")
	}
}
