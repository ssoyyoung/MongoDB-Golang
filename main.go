package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
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
	c.Response().Header().Set("Access-Control-Allow-Origin", "*")
	c.Response().Header().Set("Access-Control-Allow-Methods", "POST, GET")
	c.Response().Header().Set("Access-Control-Allow-Headers", "*")
	res := mongodb.ListData()

	return c.String(http.StatusOK, res)
}

///////////////////Meerkat CRUD FUNC///////////////////
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
	c.Response().Header().Set("Access-Control-Allow-Headers", "*")
	id := c.Param("id")
	res := mongodb.DeleteDBbyID(id)

	return c.String(http.StatusOK, res)
}

func updateStreamer(c echo.Context) error {
	c.Response().Header().Set("Access-Control-Allow-Origin", "*")
	c.Response().Header().Set("Access-Control-Allow-Methods", "POST, GET")
	c.Response().Header().Set("Access-Control-Allow-Headers", "*")
	id := c.Param("id")
	fmt.Println(c.FormValue("platform"), c.FormValue("channel"), c.FormValue("channelID"))
	res := mongodb.UpdateDBbyID(id, c.FormValue("platform"), c.FormValue("channel"), c.FormValue("channelID"))

	return c.String(http.StatusOK, res)
}

func createStreamer(c echo.Context) error {
	c.Response().Header().Set("Access-Control-Allow-Origin", "*")
	c.Response().Header().Set("Access-Control-Allow-Methods", "POST, GET")
	c.Response().Header().Set("Access-Control-Allow-Headers", "*")

	fmt.Println("test")
	fmt.Println(c.FormValue("platform"), c.FormValue("channel"), c.FormValue("channelID"))
	res := mongodb.CreateDB(c.FormValue("platform"), c.FormValue("channel"), c.FormValue("channelID"))

	return c.String(http.StatusOK, res)
}

func userInfo2(c echo.Context) error {
	c.Response().Header().Set("Access-Control-Allow-Origin", "*")
	c.Response().Header().Set("Access-Control-Allow-Methods", "*")
	c.Response().Header().Set("Access-Control-Allow-Headers", "*")

	fmt.Println("userInfo")
	fmt.Println(c)
	fmt.Println(c.FormValue("email"), c.FormValue("name"), c.FormValue("googleId"), c.FormValue("imageUrl"))

	return c.String(http.StatusOK, "userinfo")
}

// Login Func

type handler struct{}

func (h *handler) login(c echo.Context) error {
	c.Response().Header().Set("Access-Control-Allow-Origin", "*")
	c.Response().Header().Set("Access-Control-Allow-Methods", "*")
	c.Response().Header().Set("Access-Control-Allow-Headers", "*")

	googleID := c.FormValue("googleId")
	name := c.FormValue("name")
	email := c.FormValue("email")

	res := mongodb.CheckUser(googleID, name, email)
	if res {
		// Create token
		token := jwt.New(jwt.SigningMethodHS256)

		claims := token.Claims.(jwt.MapClaims)
		claims["name"] = name
		claims["googleId"] = googleID
		claims["admin"] = true
		claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

		t, err := token.SignedString([]byte("secret"))
		if err != nil {
			return err
		}
		mongodb.UpdateUser(googleID, t)
		return c.JSON(http.StatusOK, map[string]string{
			"token": t,
		})
	}
	return echo.ErrUnauthorized
}

func (h *handler) private(c echo.Context) error {
	fmt.Println(c.Get("user"))
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	name := claims["name"].(string)

	return c.String(http.StatusOK, "Welcome "+name+"!")
}

//IsLoggedIn FUNC
var isLoggedIn = middleware.JWTWithConfig(middleware.JWTConfig{
	SigningKey: []byte("secret"),
})

func isAdmin(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		user := c.Get("user").(*jwt.Token)
		claims := user.Claims.(jwt.MapClaims)
		isAdmin := claims["admin"].(bool)
		if isAdmin == false {
			return echo.ErrUnauthorized
		}
		return next(c)
	}
}

//ToDo
//Update Token and Refresh Token
//Update Header Authorization

func main() {
	e := echo.New()

	e.GET("/", goPage)
	e.GET("/saveData", saveHandler)
	e.GET("/getList", getList)
	e.GET("/getStreamers", getStreamers)
	e.GET("/getStreamer/:id", getStreamerByID)
	e.GET("/deleteStreamer/:id", deleteStreamer)
	e.POST("/updateStreamer/:id", updateStreamer)
	e.POST("/createStreamer", createStreamer)
	e.POST("/userInfo2", userInfo2)

	h := &handler{}
	//Login Func
	e.POST("/userInfo", h.login)
	e.GET("/private", h.private, isLoggedIn)
	e.GET("/admin", h.private, isLoggedIn, isAdmin)

	e.Logger.Fatal(e.Start(":1323"))
}
