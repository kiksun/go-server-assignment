package handler

import (
	"fmt"
	"net/http"
	"problem1/cache"
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

	cacheKey := fmt.Sprintf("getFriendOfFriend_ID:%d", ID)
	cacheValue, f := cache.GetCacheValue(cacheKey)
	if f {
		return c.JSON(http.StatusOK, cacheValue)
	}

	DB := database.GetDB()
	friendOfFriendList, err := handlerUtil.GetFriendsOfFriends(DB, ID, nil, nil)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	cache.SetCacheValue(cacheKey, friendOfFriendList)

	return c.JSON(http.StatusOK, friendOfFriendList)
}
