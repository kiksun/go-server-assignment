package main

import (
	"net/http"
	"problem1/configs"
	"strconv"

	"problem1/handler"

	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	conf := configs.Get()

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "minimal_sns_app")
	})
	e.GET("/get_friend_list", handler.GetFriendList)
	e.GET("/get_friend_of_friend_list", func(c echo.Context) error {
		// FIXME
		return nil
	})

	e.GET("/get_friend_of_friend_list_paging", func(c echo.Context) error {
		// FIXME
		return nil
	})

	e.Logger.Fatal(e.Start(":" + strconv.Itoa(conf.Server.Port)))
}
