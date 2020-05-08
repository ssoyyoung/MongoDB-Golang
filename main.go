package main

import (
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

func main() {
	e := echo.New()
	e.GET("/", goPage)
	e.GET("/saveData", saveHandler)
	e.Logger.Fatal(e.Start(":1323"))
}
