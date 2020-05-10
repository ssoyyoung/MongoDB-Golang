package main

import (
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
	res := mongodb.ListData()

	return c.String(http.StatusOK, res)
}

func main() {
	e := echo.New()
	
	//e.Use(CORSMiddlewareWrapper)
	//e.Use(middleware.CORS())
	//e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
	//	AllowOrigins: []string{"*"},
	//	AllowHeaders: []string{http.MethodGet, http.MethodHead, http.MethodPut, http.MethodPatch, http.MethodPost, http.MethodDelete},
	//}))

	e.GET("/", goPage)
	e.GET("/saveData", saveHandler)
	e.GET("/getList", getList)
	e.Logger.Fatal(e.Start(":1323"))
}

// CORSMiddlewareWrapper func
/* func CORSMiddlewareWrapper(next echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		req := ctx.Request()
		dynamicCORSConfig := middleware.CORSConfig{
			AllowOrigins: []string{req.Header.Get("Origin")},
			AllowHeaders: []string{"Accept", "Cache-Control", "Content-Type", "X-Requested-With"},
		}
		CORSMiddleware := middleware.CORSWithConfig(dynamicCORSConfig)
		CORSHandler := CORSMiddleware(next)
		return CORSHandler(ctx)
	}
} */
