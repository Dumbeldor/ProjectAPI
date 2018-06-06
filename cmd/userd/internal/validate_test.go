package internal

import "testing"

func TestAuthRequest_Validate(t *testing.T) {
	arTest := registerRequest{Login: "", Email: "vincent.glize@live.fr", Password: "testttttt"}
	if arTest.Validate() == nil {
		t.Error("expected not nil")
	}

	arTest.Login = "test"
	arTest.Password = ""
	if arTest.Validate() == nil {
		t.Error("expected not nil")
	}

	arTest.Password = "test"
	if arTest.Validate() == nil {
		t.Error("expected 8 characters is the minimum password lenght")
	}

	arTest.Login = ""
	arTest.Password = ""
	if arTest.Validate() == nil {
		t.Error("expected not nil")
	}

	charactersAllow := "a-z A-Z 0-9 _ -"
	arTest.Login = "azerhjklmnbvcxw_-"
	arTest.Password = "klsdfjklsdkfsdjk45645"
	if arTest.Validate() != nil {
		t.Errorf("expected Login is invalid. Characters allows : %s", charactersAllow)
	}
	arTest.Login = "???????????"
	if arTest.Validate() == nil {
		t.Errorf("expected Login is invalid. Characters allows : %s", charactersAllow)
	}
	arTest.Login = "ççççççççççççç"
	if arTest.Validate() == nil {
		t.Errorf("expected Login is invalid. Characters allows : %s", charactersAllow)
	}
	arTest.Login = "^^^^^^^"
	if arTest.Validate() == nil {
		t.Errorf("expected Login is invalid. Characters allows : %s", charactersAllow)
	}
	arTest.Login = "&&&&&&&&"
	if arTest.Validate() == nil {
		t.Errorf("expected Login is invalid. Characters allows : %s", charactersAllow)
	}
	arTest.Login = "test test test"
	if arTest.Validate() == nil {
		t.Errorf("expected Login is invalid. Characters allows : %s", charactersAllow)
	}
	arTest.Login = "@@@@@@@@"
	if arTest.Validate() == nil {
		t.Errorf("expected Login is invalid. Characters allows : %s", charactersAllow)
	}
	arTest.Login = "!!!!!!!!!!!"
	if arTest.Validate() == nil {
		t.Errorf("expected Login is invalid. Characters allows : %s", charactersAllow)
	}

	arTest.Login = "testatest"
	arTest.Email = "vincent.glizegoogle.inc"
	if arTest.Validate() == nil {
		t.Error("Validate email invalid")
	}

	arTest.Email = "vincent glize@google.inc"
	if arTest.Validate() == nil {
		t.Error("Validate email invalid")
	}

	arTest.Email = "vincent.glize@live.fr"
	if arTest.Validate() != nil {
		t.Errorf("Validate email invalid")
	}
}
