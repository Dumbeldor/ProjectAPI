package internal

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/labstack/echo"
	"gitlab.com/projetAPI/easyhttp"
	"gitlab.com/projetAPI/ProjetAPI/service"
	"log"
	"net/http"
	"os"
	"strings"
	"testing"
	"github.com/stretchr/testify/assert"
)

const (
	tableCreationQuery = `CREATE TABLE IF NOT EXISTS users (
    	user_id UUID NOT NULL,
    	login TEXT NOT NULL,
		email TEXT NOT NULL,
    	password TEXT NOT NULL,
    	creation_date TIMESTAMP NOT NULL DEFAULT NOW(),
    	locked BOOL NOT NULL DEFAULT('f'),
    	salt1 TEXT NOT NULL,
    	salt2 TEXT NOT NULL,
    	PRIMARY KEY (user_id)
	);`

	extension = `CREATE EXTENSION "uuid-ossp";`

	insertUserTest = `INSERT INTO users (user_id, login, email, password, salt1, salt2)
		VALUES (uuid_generate_v4(), 'test', 'vincent@dskjgf.inc',
		'b93bfda11b54035348905c83bb5fbeb2319c00e5f2867d7ccece135f227202091a0ac9194ad4e5163bed76d77680f4e53878710d72f23ab429404c7b0d83f1a2',
		']~>{QPjkCp~rFQ(/(PSJL<,[afOb|gr>-_Aj^DD[Pxz;M*jEWjroARO;|.nHl)TtGLEzaPOzW:%r[rsVh=lJjf[ENk]euu%t_,adZ.j?IGI{{e/zC=H^%Vc{-R:h)^O*Z$I$d]Wqq^%&feCg_hckzlifOCIW<t%YQP_:L,$$%HZpxycEKd!Fy/&NaGU^]MoSX*|#,-~R#r?UBzx$avz,Rnf^LvIRBnpjsB<hH)[]m[%#uFGuEB$hboyH/~kPfOrH',
		'Go]uWPZ%UvL%OJL_tfH{G~).qeu#iRr/xCr~j(v]iGm}&o>cCKU_L!C?EQYt|FSa!IRx*vz[}YoI!VYn:MQKegvQKxZe%}-O!Va%ID*Ipo:,vj[rJRE|]pu.]W%^AcJ)A)L:^>AOKPrQ?sfkoZvKLan{^s*YRsOzjy$<E.{oWWywoB<h^.w),<<KUyxm_EXerpyJ$[hb(>q)-aN%tr=O(Zq?_]Z<~ieIjnYmb$^<U#>:#v#AJE(q;hSbcb;PcT!.')`
)

var (
	configFile = ""
	userJSON   = `{"login":"test","password":"test12345"}`
	session    map[string]interface{}
)

func init() {
	flag.StringVar(&configFile, "config", "", "Configuration file")
}

func TestMain(m *testing.M) {
	flag.Parse()
	gconfig.load(configFile)

	app = service.New(
		"uc-auth",
		&gconfig.HTTP,
		&gconfig.Log,
		AppVersion,
		AppBuildDate)

	connectDB()

	createSessionWriter()

	createSessionReader()

	ensureTableExists()

	createExtension()

	insertUser()

	code := m.Run()

	clearTable()

	os.Exit(code)
}

func connectDB() {
	if !verifyAuthDB() {
		log.Fatalf("Critical server error. Can't connect to database")
	}
}

func createSessionWriter() {
	sessionWriter = newWriter(gconfig.Redis)
}

func createSessionReader() {
	sessionReader = service.NewReader(gconfig.Redis)
}

func ensureTableExists() {
	if _, err := gAuthDB.nativeDB.Exec(tableCreationQuery); err != nil {
		log.Fatal(err)
	}
}

func createExtension() {
	gAuthDB.nativeDB.Exec(extension)
}

func insertUser() {
	if _, err := gAuthDB.nativeDB.Exec(insertUserTest); err != nil {
		log.Fatal(err)
	}
}

func clearTable() {
	gAuthDB.nativeDB.Exec("DELETE FROM users")
}

func TestAuthRequest_Validate(t *testing.T) {
	arTest1 := authRequest{}
	arTest1.Body.Login = ""
	arTest1.Body.Password = "test"
	if arTest1.Validate() == nil {
		t.Error("expected not nil")
	}

	arTest2 := authRequest{}
	arTest2.Body.Login = "test"
	arTest2.Body.Password = ""
	if arTest2.Validate() == nil {
		t.Error("expected not nil")
	}

	arTest3 := authRequest{}
	arTest3.Body.Login = ""
	arTest3.Body.Password = ""
	if arTest3.Validate() == nil {
		t.Error("expected not nil")
	}

	arTest4 := authRequest{}
	arTest4.Body.Login = "test"
	arTest4.Body.Password = "test"
	if arTest4.Validate() != nil {
		t.Error("expected nil")
	}
}

func TestLoginPostInvalideMdp_Validate(t *testing.T) {
	e := echo.New()
	userJSONBad   := `{"login":"test","password":"badmdp"}`
	req, err := http.NewRequest("POST", "/v1/auth/login", strings.NewReader(userJSONBad))
	if err != nil {
		t.Error("Error /v1/auth/login")
	}
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	rr, err := easyhttp.ExecuteRequest(req, httpAuthLogin, e)
	easyhttp.CheckResponseCode(t, http.StatusForbidden, rr.Code)

	if err != nil {
		t.Errorf("Handler return error : %s", err)
	}

	assert.Equal(t, `{"message":"Invalid user/password."}`, rr.Body.String())
}


func TestLoginPost_Validate(t *testing.T) {
	e := echo.New()
	req, err := http.NewRequest("POST", "/v1/auth/login", strings.NewReader(userJSON))
	if err != nil {
		t.Error("Error /v1/auth/login")
	}
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	rr, err := easyhttp.ExecuteRequest(req, httpAuthLogin, e)
	easyhttp.CheckResponseCode(t, http.StatusOK, rr.Code)

	if err != nil {
		t.Errorf("Handler return error : %s", err)
	}

	err = json.Unmarshal(rr.Body.Bytes(), &session)
	if err != nil {
		t.Errorf("Error to get session : %s", err)
	}

	if session["token"] == nil || session["token"] == "" {
		t.Errorf("Error token is invalid !")
	}
}

func TestLogout_Validate(t *testing.T) {
	e := echo.New()
	req, err := http.NewRequest("POST", "/v1/auth/logout", strings.NewReader(``))
	if err != nil {
		t.Error("Error /v1/auth/logout")
	}
	req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %s", session["token"]))

	rr, err := easyhttp.ExecuteRequest(req, httpAuthLogout, e)
	easyhttp.CheckResponseCode(t, http.StatusOK, rr.Code)

	if err != nil {
		t.Errorf("Handler return error : %s", err)
	}
}
