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
	id, err := strconv.Atoi(c.QueryParam("ID"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid ID Value")
	}

	limit, err := strconv.Atoi(c.QueryParam("limit"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid limit Value")
	}

	page, err := strconv.Atoi(c.QueryParam("page"))
	if err != nil || page <= 0 {
		return c.JSON(http.StatusBadRequest, "Invalid page Value")
	}

	offset := (page - 1) * limit
	db := database.GetDB()

	friendOfFriendList, err := handlerUtil.GetFriendsOfFriends(db, id, &limit, &offset)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	return c.JSON(http.StatusOK, friendOfFriendList)
}
