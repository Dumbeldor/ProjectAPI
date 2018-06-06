package internal

import (
	"testing"
	"github.com/labstack/echo"
	"net/http/httptest"
	"gitlab.com/projetAPI/easyhttp"
	"net/http"
	"github.com/stretchr/testify/assert"
	"strings"
)

var (
	userJSON   = `{"login":"test_tu", "email":"vincent@live.fr", "password":"test12345"}`
)



func TestRegisterPostBadEmail_Validate(t *testing.T) {
	e := echo.New()

	userJSON   = `{"login":"test_tu", "email":"vincent@live", "password":"test12345"}`

	req := httptest.NewRequest(echo.POST, "/v1/user/register", strings.NewReader(userJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	err := httpRegister(c)

	easyhttp.CheckResponseCode(t, http.StatusNotAcceptable, rec.Code)

	if err != nil {
		t.Errorf("Handler return error : %s", err)
	}

	assert.Equal(t, `{"message":"invalid email"}`, rec.Body.String())
}

func TestRegisterPost_Validate(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(echo.POST, "/v1/user/register", strings.NewReader(userJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	err := httpRegister(c)

	easyhttp.CheckResponseCode(t, http.StatusOK, rec.Code)

	if err != nil {
		t.Errorf("Handler return error : %s", err)
	}

	assert.Equal(t, `{"message":"Registration succeed."}`, rec.Body.String())
}