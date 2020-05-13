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
	c.Response().Header().Set("Access-Control-Allow-Methods", "POST, GET")
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
	id := c.Param("id")
	res := mongodb.SearchDBbyID(id)

	return c.String(http.StatusOK, res)
}

func deleteStreamer(c echo.Context) error {
	c.Response().Header().Set("Access-Control-Allow-Origin", "*")
	c.Response().Header().Set("Access-Control-Allow-Methods", "DELETE, GET")
	c.Response().Header().Set("Access-Control-Allow-Headers", "application/json text/plain */*")
	id := c.Param("id")
	res := mongodb.DeleteDBbyID(id)

	return c.String(http.StatusOK, res)
}

func updateStreamer(c echo.Context) error {
	c.Response().Header().Set("Access-Control-Allow-Origin", "*")
	c.Response().Header().Set("Access-Control-Allow-Methods", "PUT, GET")
	c.Response().Header().Set("Access-Control-Allow-Headers", "application/json text/plain */*")
	id := c.Param("id")
	res := mongodb.UpdateDBbyID(id)

	return c.String(http.StatusOK, res)
}

func createStreamer(c echo.Context) error {
	c.Response().Header().Set("Access-Control-Allow-Origin", "*")
	c.Response().Header().Set("Access-Control-Allow-Methods", "POST, GET")
	c.Response().Header().Set("Access-Control-Allow-Headers", "application/json text/plain */*")

	fmt.Println("test")
	fmt.Println(c.FormValue("platform"), c.FormValue("channel"), c.FormValue("channelID"))
	res := mongodb.CreateDB(c.FormValue("platform"), c.FormValue("channel"), c.FormValue("channelID"))

	return c.String(http.StatusOK, res)
}

func main() {
	e := echo.New()

	e.GET("/", goPage)
	e.GET("/saveData", saveHandler)
	e.GET("/getList", getList)
	e.GET("/getStreamers", getStreamers)
	e.GET("/getStreamer/:id", getStreamerByID)
	e.DELETE("/deleteStreamer/:id", deleteStreamer)
	e.PUT("/updateStreamer/:id", updateStreamer)
	e.POST("/createStreamer", createStreamer)

	e.Logger.Fatal(e.Start(":1323"))
}
