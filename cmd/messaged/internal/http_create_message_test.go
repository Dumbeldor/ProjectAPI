package internal

import (
	"testing"
	"github.com/labstack/echo"
	"net/http/httptest"
	"strings"
	"gitlab.com/projetAPI/easyhttp"
	"github.com/stretchr/testify/assert"
	"fmt"
	"gitlab.com/projetAPI/ProjetAPI/mock"
	"net/http"
)

var (
	messageJSON = `{"message":"Unit test message", "receiver": "Vincent"}`
)

func createRequestCreateMessage(t *testing.T, messageJ string, token string, code int, msg string) {
	e := echo.New()
	req := httptest.NewRequest(echo.POST, "/v1/message/create", strings.NewReader(messageJ))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %s", token))
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	err := httpCreateMessage(c)

	if err != nil {
		t.Errorf("Handler return error : %s", err)
	}

	easyhttp.CheckResponseCode(t, code, rec.Code)

	assert.Equal(t, msg, rec.Body.String())
}

func TestCreateMessageNotJSON_Validate(t *testing.T) {
	messageJSON2 := `fdsfd`
	createRequestCreateMessage(t, messageJSON2, mock.TokenString, http.StatusBadRequest, `{"message":"Request body is not a JSON."}`)
}

func TestCreateMessageWithoutToken_Validate(t *testing.T) {
	createRequestCreateMessage(t, messageJSON, "", http.StatusUnauthorized, `{"message":"Authorization failed."}`)
}

func TestCreateMessageInvalidUser_Validate(t *testing.T) {
	messageJSON2 := `{"message":"Unit test message", "receiver": "InvalidUser"}`
	createRequestCreateMessage(t, messageJSON2, mock.TokenString, http.StatusBadRequest, `{"message":"The user does not exist."}`)
}

func TestCreateMessageWhitoutUser_Validate(t *testing.T) {
	messageJSON2 := `{"message":"Unit test message", "receiver": ""}`
	createRequestCreateMessage(t, messageJSON2, mock.TokenString, http.StatusBadRequest, `{"message":"The user does not exist."}`)
}

func TestCreateMessageWithoutMessage_Validate(t *testing.T) {
	messageJSON2 := `{"message":"", "receiver": "InvalidUser"}`
	createRequestCreateMessage(t, messageJSON2, mock.TokenString, http.StatusBadRequest, `{"message":"1 characters is the minimum message length"}`)
}

func TestCreateMessageWithoutMessage2_Validate(t *testing.T) {
	messageJSON2 := `{"receiver": "InvalidUser"}`
	createRequestCreateMessage(t, messageJSON2, mock.TokenString, http.StatusBadRequest, `{"message":"1 characters is the minimum message length"}`)
}

func TestCreateMessage_Validate(t *testing.T) {
	createRequestCreateMessage(t, messageJSON, mock.TokenString, 200, `{"message":"The sending of the message: Unit test message to the Vincent user has gone well."}`)
}
