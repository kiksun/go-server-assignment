package handler

import (
	"net/http"
	"problem1/database"
	handlerUtil "problem1/handler/util"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo/v4"
)

func GetFriendOfFriendList(c echo.Context) error {
	ID, err := strconv.Atoi(c.QueryParam("ID"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid ID Value")
	}

	DB := database.GetDB()

	friendOfFriendList, err := handlerUtil.GetFriendsOfFriends(DB, ID, nil, nil)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	return c.JSON(http.StatusOK, friendOfFriendList)
}
