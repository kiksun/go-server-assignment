package handler

import (
	"problem1/model"

	_ "github.com/go-sql-driver/mysql"
	"gorm.io/gorm"
)

func GetFriends(db *gorm.DB, userID int) ([]model.User, error) {
	var friendList []model.User

	r := db.Table("users").
		Distinct("user_id , name").
		Joins("JOIN friend_link ON (user_id = user2_id AND user1_id = ?) OR (user_id = user1_id AND user2_id = ?)", userID, userID).
		Where("user_id NOT IN (SELECT user2_id FROM block_list WHERE user1_id = ?)", userID).
		Find(&friendList)

	if r.Error != nil {
		return nil, r.Error
	}

	return friendList, nil
}

func GetFriendsOfFriends(db *gorm.DB, userID int, limit *int, o *int) ([]model.User, error) {
	var friendsOfFriendList []model.User

	sq := db.Table("friend_link").
		Distinct("CASE WHEN user1_id = ? THEN user2_id ELSE user1_id END", userID).
		Where("user1_id = ? OR user2_id = ?", userID, userID)

	r := db.Table("users").
		Distinct("user_id, name").
		Joins("JOIN friend_link ON (user_id = user2_id AND user1_id IN (?)) OR (user_id = user1_id AND user2_id IN (?))", sq, sq).
		Where("user_id NOT IN (?) AND user_id != ? ", sq, userID).
		Where("user_id NOT IN (SELECT user2_id FROM block_list WHERE user1_id = ?) AND user_id NOT IN (SELECT user1_id FROM block_list WHERE user2_id = ?)", userID, userID)

	if r.Error != nil {
		return nil, r.Error
	}

	if o != nil && limit != nil {
		r.Limit(*limit).Offset(*o)
	}
	r.Find(&friendsOfFriendList)

	return friendsOfFriendList, nil
}
