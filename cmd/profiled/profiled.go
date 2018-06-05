// main Package
//     Schemes: https
//     Host: minegame.fr
//     BasePath: /
//     Version: 1.0
//     License: Proprietary
//     Contact: Support<vincent.glize@live.fr> https://minegame.fr
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
	"gitlab.com/projetAPI/ProjetAPI/cmd/profiled/internal"
)

var configFile = ""
var showVersion bool

func init() {
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
