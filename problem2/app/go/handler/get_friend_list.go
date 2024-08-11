package handler

import (
	"fmt"
	"net/http"
	"problem1/cache"
	"problem1/database"
	handler_util "problem1/handler/util"

	"strconv"

	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo/v4"
)

func GetFriendList(c echo.Context) error {
	ID, err := strconv.Atoi(c.QueryParam("ID"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid ID Value")
	}

	cacheKey := fmt.Sprintf("getFriend_ID:%d", ID)
	cacheValue, f := cache.GetCacheValue(cacheKey)
	if f {
		fmt.Println("cache return")
		return c.JSON(http.StatusOK, cacheValue)
	}

	DB := database.GetDB()
	friendList, err := handler_util.GetFriends(DB, ID)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	cache.SetCacheValue(cacheKey, friendList)

	return c.JSON(http.StatusOK, friendList)
}
