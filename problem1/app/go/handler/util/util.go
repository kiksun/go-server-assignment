package handlerUtil

import (
	"problem1/model"

	_ "github.com/go-sql-driver/mysql"
	"gorm.io/gorm"
)

func GetFriends(db *gorm.DB, userID int) ([]model.User, error) {
	var friendList []model.User

	result := db.Table("users").
		Distinct("user_id , name").
		Joins("JOIN friend_link ON (user_id = user2_id AND user1_id = ?) OR (user_id = user1_id AND user2_id = ?)", userID, userID).
		Where("user_id NOT IN (SELECT user2_id FROM block_list WHERE user1_id = ?)", userID).
		Find(&friendList)

	if result.Error != nil {
		return nil, result.Error
	}

	return friendList, nil
}

func GetFriendsOfFriends(db *gorm.DB, userID int, limit *int, offset *int) ([]model.User, error) {
	var friendsOfFriendList []model.User

	subquery := db.Table("friend_link").
		Distinct("CASE WHEN user1_id = ? THEN user2_id ELSE user1_id END", userID).
		Where("user1_id = ? OR user2_id = ?", userID, userID)

	result := db.Table("users").
		Distinct("user_id, name").
		Joins("JOIN friend_link ON (user_id = user2_id AND user1_id IN (?)) OR (user_id = user1_id AND user2_id IN (?))", subquery, subquery).
		Where("user_id NOT IN (?) AND user_id != ?", subquery, userID).
		Where("user_id NOT IN (SELECT user2_id FROM block_list WHERE user1_id = ?) AND user_id NOT IN (SELECT user1_id FROM block_list WHERE user2_id = ?)", userID, userID)

	if result.Error != nil {
		return nil, result.Error
	}

	if offset != nil && limit != nil {
		result.
			Limit(*limit).
			Offset(*offset)
	}
	result.Find(&friendsOfFriendList)

	return friendsOfFriendList, nil
}
