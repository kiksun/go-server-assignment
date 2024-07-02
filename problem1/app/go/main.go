package main

import (
	"net/http"
	"problem1/configs"
	"problem1/database"
	"strconv"

	"problem1/handler"

	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	defer database.Close()

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "minimal_sns_app")
	})
	e.GET("/get_friend_list", handler.GetFriendList)
	e.GET("/get_friend_of_friend_list", handler.GetFriendOfFriendList)
	e.GET("/get_friend_of_friend_list_paging", handler.GetFriendOfFriendListPaging)
	e.POST("/add_friend", handler.AddFriend)

	e.Logger.Info(e.Start(":" + strconv.Itoa(configs.Get().Server.Port)))
}
