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
	e.GET("/get_friend_of_friend_list", handler.GetFriendOfFriendList)
	e.GET("/get_friend_of_friend_list_paging", handler.GetFriendOfFriendListPaging)
	e.POST("/add_friend_list",handler.AddFriend)

	e.Logger.Fatal(e.Start(":" + strconv.Itoa(conf.Server.Port)))
}
