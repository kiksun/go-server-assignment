package handler

import (
	"log"
	"net/http"
	"problem1/database"
	"problem1/model"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo/v4"
)

const getFriendOfFriendListQuery = `
			  SELECT user_id, name 
			  FROM users
			  WHERE user_id IN( SELECT user2_id FROM friend_link WHERE user1_id = ANY( SELECT user_id FROM users JOIN friend_link ON (user_id = user2_id AND user1_id = ?) OR (user_id = user1_id AND user2_id = ?)))
			  AND user_id NOT IN (SELECT user2_id FROM block_list WHERE user1_id = ?) 
			  AND user_id Not IN (SELECT user_id FROM users JOIN friend_link ON (user_id = user2_id AND user1_id = ?) OR (user_id = user1_id AND user2_id = ?))
			  AND user_id != ?;
			  `

func GetFriendOfFriendList(c echo.Context) error {
	id, err := strconv.Atoi(c.QueryParam("ID"))
	if err != nil {
		log.Fatal(err)
	}

	db, err := database.ConnectDB()
	if err != nil {
		log.Fatal(err)
	}

	var friendOfFriendList []model.User
	rows, err := db.Query(getFriendOfFriendListQuery, id, id, id, id, id, id)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var friendOfFriend model.User
		err := rows.Scan(&friendOfFriend.UserID, &friendOfFriend.Name)
		if err != nil {
			log.Fatal(err)
		}
		friendOfFriendList = append(friendOfFriendList, friendOfFriend)
	}

	defer database.CloseDB(db)
	return c.JSON(http.StatusOK, friendOfFriendList)
}
