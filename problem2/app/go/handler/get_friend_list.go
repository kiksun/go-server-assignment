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

const getFriendListQuery = `
              SELECT user_id , name
			  FROM users 
			  JOIN friend_link ON (user_id = user2_id AND user1_id = ?) OR (user_id = user1_id AND user2_id = ?)
			  `

func GetFriendList(c echo.Context) error {
	id, err := strconv.Atoi(c.QueryParam("ID"))
	if err != nil {
		log.Fatal(err)
	}

	db, err := database.ConnectDB()
	if err != nil {
		log.Fatal(err)
	}
	defer database.CloseDB(db)

	rows, err := db.Query(getFriendListQuery, id, id)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var friendList []model.User
	for rows.Next() {
		var user model.User
		err := rows.Scan(&user.UserID, &user.Name)
		if err != nil {
			log.Fatal(err)
		}
		friendList = append(friendList, user)
	}

	return c.JSON(http.StatusOK, friendList)
}
