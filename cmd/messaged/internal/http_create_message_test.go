package internal

import (
	"testing"
	"github.com/labstack/echo"
	"net/http/httptest"
	"strings"
	"gitlab.com/projetAPI/easyhttp"
	"github.com/stretchr/testify/assert"
)

var (
	messageJSON = `{"message":"Unit test message", "user_sender_id":"kdjfskfdlkfdl", "user_receiver_id": "fsdkfsdjfsd"}`
)

func createRequestCreateMessage(t *testing.T, messageJ string, code int, msg string) {
	e := echo.New()
	req := httptest.NewRequest(echo.POST, "/v1/message/create", strings.NewReader(messageJ))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	err := httpCreateMessage(c)

	if err != nil {
		t.Errorf("Handler return error : %s", err)
	}

	easyhttp.CheckResponseCode(t, code, rec.Code)

	assert.Equal(t, msg, rec.Body.String())
}

func TestCreateMessage_Validate(t *testing.T) {
	createRequestCreateMessage(t, messageJSON, 200, `{"message":"Send successful message."}`)
}
