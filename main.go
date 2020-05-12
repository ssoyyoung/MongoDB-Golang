package main

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo"
	mongodb "github.com/ssoyyoung.p/MongoDB-Golang/mongo"
)

func goPage(c echo.Context) error {
	return c.File("home.html")
}

func saveHandler(c echo.Context) error {
	mongodb.MongoDB()
	return nil
}

func getList(c echo.Context) error {
	c.Response().Header().Set("Access-Control-Allow-Origin", "*")
	c.Response().Header().Set("Access-Control-Allow-Methods", "POST")
	c.Response().Header().Set("Access-Control-Allow-Headers", "application/json text/plain */*")
	res := mongodb.ListData()

	return c.String(http.StatusOK, res)
}

///////////////////ADMIN FUNC///////////////////
func getStreamers(c echo.Context) error {
	c.Response().Header().Set("Access-Control-Allow-Origin", "*")
	res := mongodb.CrawlList()

	return c.String(http.StatusOK, res)
}

func getStreamerByID(c echo.Context) error {
	c.Response().Header().Set("Access-Control-Allow-Origin", "*")
	_id := c.Param("_id")
	fmt.Println(_id)
	res := mongodb.SearchDBbyID(_id)

	return c.String(http.StatusOK, res)
}

func main() {
	e := echo.New()

	e.GET("/", goPage)
	e.GET("/saveData", saveHandler)
	e.GET("/getList", getList)
	e.GET("/getStreamers", getStreamers)
	e.GET("/getStreamer/:id", getStreamerByID)

	e.Logger.Fatal(e.Start(":1323"))
}
