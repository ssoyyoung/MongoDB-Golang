package main

import (
	"github.com/labstack/echo"
	mongodb "github.com/ssoyyoung.p/MongoDB-Golang/mongo"
)

func saveHandler(c echo.Context) error {
	mongodb.MongoDB()
	return nil
}

func main() {
	e := echo.New()
	e.GET("/saveData", saveHandler)
	e.Logger.Fatal(e.Start(":1323"))
}
