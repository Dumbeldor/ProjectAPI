package main

import (
	"net/http"

	"encoding/json"
	"fmt"
	"github.com/labstack/echo"
	"io/ioutil"
)

type resAircraft struct {
	FlgH int `json:"flgH"`
}

func main() {
	e := echo.New()
	e.GET("/:lat/:lng/:dstU", func(c echo.Context) error {
		lat := c.Param("lat")
		lng := c.Param("lng")
		dstL := "0"
		dstU := c.Param("dstU")

		url := "https://public-api.adsbexchange.com/VirtualRadar/AircraftList.json?lat=" + lat + "&lng=" + lng + "&fDstL=" + dstL + "&fDstU=" + dstU

		fmt.Printf("url : %s", url)

		response, err := http.Get(url)
		if err != nil {
			e.Logger.Fatalf("Erreur : %s", err)
			return c.String(http.StatusInternalServerError, "Erreur Interne")
		}
		defer response.Body.Close()

		contents, err := ioutil.ReadAll(response.Body)
		fmt.Println(string(contents))

		var dat map[string]interface{}
		if err := json.Unmarshal(contents, &dat); err != nil {
			e.Logger.Fatalf("Erreur : %s", err)
			return c.String(http.StatusInternalServerError, "Erreur Interne")
		}
		fmt.Println(dat)

		feeds := dat["feeds"].([]interface{})
		feed := feeds[0].(map[string]interface{})
		id := feed["name"].(string)
		//id := feed[0].(float64)
		fmt.Println(id)

		//acLists := dat["acList"].([]interface{})
		//for

		return c.JSON(http.StatusOK, string(contents))
	})
	e.Logger.Fatal(e.Start(":1323"))
}
