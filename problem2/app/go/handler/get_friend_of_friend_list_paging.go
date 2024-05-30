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

const getFriendOfFriendListPagingQuery = `
			  SELECT user_id, name 
			  FROM users
			  WHERE user_id IN( SELECT user2_id FROM friend_link WHERE user1_id = ANY( SELECT user_id FROM users JOIN friend_link ON (user_id = user2_id AND user1_id = ?) OR (user_id = user1_id AND user2_id = ?)))
			  AND user_id NOT IN (SELECT user2_id FROM block_list WHERE user1_id = ?) 
			  AND user_id Not IN (SELECT user_id FROM users JOIN friend_link ON (user_id = user2_id AND user1_id = ?) OR (user_id = user1_id AND user2_id = ?))
			  AND user_id != ?
			  LIMIT ? OFFSET ?
			  `

func GetFriendOfFriendListPaging(c echo.Context) error {
	id, err := strconv.Atoi(c.QueryParam("ID"))
	if err != nil {
		log.Fatal(err)
	}

	limit, err := strconv.Atoi(c.QueryParam("limit"))
	if err != nil {
		log.Fatal(err)
	}

	page, err := strconv.Atoi(c.QueryParam("page"))
	if err != nil {
		log.Fatal(err)
	}
	if page <= 0 {
		log.Fatal("Invalid Value")
	}

	db, err := database.ConnectDB()
	if err != nil {
		log.Fatal(err)
	}

	var friendOfFriendList []model.User
	var offset = (page - 1) * limit
	rows, err := db.Query(getFriendOfFriendListPagingQuery, id, id, id, id, id, id, limit, offset)
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
