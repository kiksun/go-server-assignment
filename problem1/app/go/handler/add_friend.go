package handler

import (
	"fmt"
	"net/http"
	"problem1/database"
	"problem1/model"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func AddFriend(c echo.Context) error {
	uID, err := strconv.Atoi(c.QueryParam("userID"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "userID is invlid value"})
	}
	fID, err := strconv.Atoi(c.QueryParam("friendID"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "friendID is invlid value"})
	}
	DB := database.GetDB()
	var friendLink model.FriendLink

	err = DB.Table("friend_link").
		Where("user1_id = ? AND user2_id = ? OR user1_id = ? AND user2_id = ?", uID, fID, uID, fID).
		First(&friendLink).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			newFriend := model.FriendLink{User1ID: uID, User2ID: fID}
			DB.Create(&newFriend)
			fmt.Println("New user added:", newFriend)
		} else {
			fmt.Println(err)
		}
	} else {
		fmt.Println("FriendLink already exists:", friendLink)
		return c.JSON(http.StatusOK, "friend already exists")
	}

	return c.JSON(http.StatusOK, "add friend success")
}
