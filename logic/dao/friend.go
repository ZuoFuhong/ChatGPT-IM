package dao

import (
	"go-IM/logic/db"
	"go-IM/logic/model"
)

type friendDao struct{}

var FriendDao = new(friendDao)

func (*friendDao) Insert(appId, userId, friendId int64) error {
	_, err := db.Cli.Exec("INSERT INTO friend(app_id, user_id, friend_id) VALUES(?, ?, ?)", appId, userId, friendId)
	return err
}

func (*friendDao) SelectFriends(userId int64) (*[]model.User, error) {
	rows, err := db.Cli.Query(`
		SELECT u.app_id, u.user_id, u.nickname, u.sex, u.avatar_url, u.extra, u.create_time, u.update_time
		FROM friend f
		LEFT JOIN user u ON u.user_id = f.friend_id
		WHERE f.user_id = ?`, userId)
	if err != nil {
		return nil, err
	}
	users := make([]model.User, 0, 5)
	for rows.Next() {
		user := new(model.User)
		err := rows.Scan(&user.AppId, &user.UserId, &user.Nickname, &user.Sex, &user.AvatarUrl, &user.Extra, &user.CreateTime, &user.UpdateTime)
		if err != nil {
			return nil, err
		}
		users = append(users, *user)
	}
	return &users, nil
}
