package handler

import (
	"log"
	"net/http"
	"problem1/database"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo/v4"
)

const addFriendQuery = `INSERT INTO friend_link  (user1_id, user2_id) VALUES (?,?)`

func AddFriend(c echo.Context) error {
	userId, err := strconv.Atoi(c.QueryParam("userID"))
	if err != nil {
		log.Fatal(err)
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "userID is invlid value"})
	}
	friendId, err := strconv.Atoi(c.QueryParam("friendID"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "friendID is invlid value"})
	}

	db, err := database.ConnectDB()
	if err != nil {
		log.Fatal(err)
	}

	_, err = db.Exec(addFriendQuery, userId, friendId)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Failed to insert into database"})
	}

	return c.JSON(http.StatusOK, map[string]string{"message": "Friend added successfully"})
}
