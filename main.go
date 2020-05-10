package main

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
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
	res := mongodb.ListData()

	return c.String(http.StatusOK, res)
}

func main() {
	e := echo.New()
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"https://localhost", "https://49.247.134.77"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}))
	e.GET("/", goPage)
	e.GET("/saveData", saveHandler)
	e.GET("/getList", getList)
	e.Logger.Fatal(e.Start(":1323"))
}
