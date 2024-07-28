package model

type User struct {
	UserID int    `json:"user_id"`
	Name   string `json:"name"`
}

type FriendLink struct {
	User1ID int `json:"user1_id"`
	User2ID int `json:"user2_id"`
}
