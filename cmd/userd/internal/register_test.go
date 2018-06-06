package internal

import (
	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
	"gitlab.com/projetAPI/easyhttp"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

var (
	userJSON = `{"login":"test_tu", "email":"vincent@live.fr", "password":"test12345"}`
)

func createRequestRegister(t *testing.T, userJ string, code int, msg string) {
	e := echo.New()
	req := httptest.NewRequest(echo.POST, "/v1/user/register", strings.NewReader(userJ))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	err := httpRegister(c)

	easyhttp.CheckResponseCode(t, code, rec.Code)

	if err != nil {
		t.Errorf("Handler return error : %s", err)
	}

	assert.Equal(t, msg, rec.Body.String())
}

func TestRegisterPostBadEmail_Validate(t *testing.T) {
	userJSONInvalide := `{"login":"test_tu", "email":"vincent@live", "password":"test12345"}`

	createRequestRegister(t, userJSONInvalide, http.StatusNotAcceptable, `{"message":"invalid email"}`)
}

func TestRegisterPostBadLogin_Validate(t *testing.T) {
	userJSONInvalide := `{"login":"test$", "email":"vincent@live.fr", "password":"test12345"}`

	createRequestRegister(t, userJSONInvalide, http.StatusNotAcceptable, `{"message":"invalid login, allowed characters are a-z A-Z 0-9"}`)
}

func TestRegisterPostMinimumLogin_Validate(t *testing.T) {
	userJSONInvalide := `{"login":"te", "email":"vincent@live.fr", "password":"test12345"}`

	createRequestRegister(t, userJSONInvalide, http.StatusNotAcceptable, `{"message":"3 characters is the minimum login length"}`)
}

func TestRegisterPostMaxLogin_Validate(t *testing.T) {
	userJSONInvalide := `{"login":"testestestestestestes", "email":"vincent@live.fr", "password":"test12345"}`

	createRequestRegister(t, userJSONInvalide, http.StatusNotAcceptable, `{"message":"20 characters is the maximum login length"}`)
}

func TestRegisterPostMiniPassword_Validate(t *testing.T) {
	userJSONInvalide := `{"login":"test", "email":"vincent@live.fr", "password":"testest"}`

	createRequestRegister(t, userJSONInvalide, http.StatusNotAcceptable, `{"message":"8 characters is the minimal password length"}`)
}

func TestRegisterPost_Validate(t *testing.T) {
	createRequestRegister(t, userJSON, http.StatusOK, `{"message":"Registration succeed."}`)
}

func TestRegisterPostLoginAlreadyTaken_Validate(t *testing.T) {
	createRequestRegister(t, userJSON, http.StatusConflict, `{"message":"Login already taken."}`)
}
