package handler

import (
	"net/http"
	"problem1/database"
	handlerUtil "problem1/handler/util"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo/v4"
)

func GetFriendOfFriendListPaging(c echo.Context) error {
	ID, err := strconv.Atoi(c.QueryParam("ID"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid ID Value")
	}

	l, err := strconv.Atoi(c.QueryParam("limit"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid limit Value")
	}

	p, err := strconv.Atoi(c.QueryParam("page"))
	if err != nil || p <= 0 {
		return c.JSON(http.StatusBadRequest, "Invalid page Value")
	}

	o := (p - 1) * l
	DB := database.GetDB()

	friendOfFriendList, err := handlerUtil.GetFriendsOfFriends(DB, ID, &l, &o)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	return c.JSON(http.StatusOK, friendOfFriendList)
}
