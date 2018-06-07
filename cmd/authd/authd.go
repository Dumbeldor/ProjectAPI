// authd
//     Schemes: http, https
//     Host: localhost
//     BasePath: /
//     Version: 1.0
//     License: Proprietary
//     Contact: Support<support@glizou.fr> https://glizou.fr
//
//     Consumes:
//     - application/json
//     Produces:
//     - application/json
// swagger:meta
package main

import (
	"fmt"
	"github.com/pborman/getopt/v2"
	"gitlab.com/projetAPI/ProjetAPI/cmd/authd/internal"
)

var configFile = ""
var showVersion bool

func init() {
	showVersion = false
	getopt.FlagLong(&configFile, "config", 'c', "Configuration file")
	getopt.FlagLong(&showVersion, "version", 'V', "Show application version")
}

func main() {
	getopt.Parse()

	if showVersion {
		fmt.Printf("version: %s\nbuild date: %s\n", internal.AppVersion, internal.AppBuildDate)
		return
	}

	internal.StartApp(configFile)
}
