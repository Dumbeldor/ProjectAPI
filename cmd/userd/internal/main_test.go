package internal

import (
	"flag"
	"gitlab.com/projetAPI/ProjetAPI/service"
	"os"
	"testing"
)

const (
	insertUserTest = `INSERT INTO users(user_id, login, password, salt1, salt2, email) VALUES('a1472317-f1e2-413e-b55d-d412b9f4953f',	'test',
                         'c5f9317c2b2d0a4ee2c5471993d7825342f2a510937a06c09e4774349393545f3deaca60cbb60e1e5f1ef8d7683ad1934bbdcf1c1081243981d53efdbc714142',
                         'Vt>*Sw%USNYUHvRk)f:{/>!caBF]#yD^KASeWyD!(~;r#Z(rJzw]FdtWP&T}[K}/,xvbau,[/l&oHI%/%^(TwWpoQJDOMWVX(mS[Y#R]:C)#=OHqM%<!%NJehawX<B,X*(D/ZAa{ysqWkkkeyEIie{S)Z}?,|pLV=bw?HCjTU&DL^-CwlH[|k/{I/<d;h$}l[*<O*Np-FYtbxhUHjwZjCZ$-%:R/W*;TPSxnGM{WPX,#YmXZvj*J)lM}ddEfT^fg',
                         '&BMZx^Avx%]-AGpe)MEIZIGK*m%)f{uAw:>UEW%&;-MMe?e]eQ-aij~O#iJY,vnV>}LJJe(EBct<Q}(bENED-A(fKGNg%Ac:x{RSu[m|Ko)bf$&fuXakO&%O*%Pg;xx*NEPVCHS/d:lI%BLdVyth*)/U}FG[eVjMPZZj/Gv^D^Pf:kz>NNTHVap.aS^b?MRu.=E;yLtT;>mAhD)$S:_czVpstDrf/$TNj:Jugue*,S/M]IFfk>g>omy]MjopKdAx',
                         'test@test.fr')`
)

var (
	configFile = ""
)

func init() {
	flag.StringVar(&configFile, "config", "", "Configuration file")
}

// Init TU
func TestMain(m *testing.M) {
	flag.Parse()
	configLoaded := gconfig.load(configFile)

	if !configLoaded {
		gconfig.UsersDB.URL = "host=127.0.0.1 dbname=postgres user=postgres password=example sslmode=disable"
	}

	app = service.New(
		"glizou-user-TEST",
		&gconfig.HTTP,
		&gconfig.Log,
		AppVersion,
		AppBuildDate)

	if !verifyUserDB() {
		app.Log.Fatalf("Critical server error. Can't connect to user database")
	}

	gUserDB.ClearTable()

	insertUser()

	code := m.Run()

	os.Exit(code)
}

func insertUser() {
	gUserDB.ExecSql(insertUserTest)
}
