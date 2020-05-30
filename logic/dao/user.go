package dao

import (
	"database/sql"
	"go-IM/logic/db"
	"go-IM/logic/model"
)

type userDao struct{}

var UserDao = new(userDao)

// 插入一条用户信息
func (*userDao) Add(user *model.User) (int64, error) {
	result, err := db.Cli.Exec("INSERT IGNORE INTO user(app_id,user_id,nickname,sex,avatar_url,extra) VALUES (?,?,?,?,?,?)",
		user.AppId, user.UserId, user.Nickname, user.Sex, user.AvatarUrl, user.Extra)
	if err != nil {
		return 0, err
	}

	affected, err := result.RowsAffected()
	if err != nil {
		return 0, err
	}
	return affected, nil
}

// 获取用户信息
func (*userDao) Get(appId, userId int64) (*model.User, error) {
	row := db.Cli.QueryRow("SELECT nickname,sex,avatar_url,extra,create_time,update_time FROM user WHERE app_id = ? AND user_id = ?",
		appId, userId)
	user := model.User{
		AppId:  appId,
		UserId: userId,
	}

	err := row.Scan(&user.Nickname, &user.Sex, &user.AvatarUrl, &user.Extra, &user.CreateTime, &user.UpdateTime)
	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}

	if err == sql.ErrNoRows {
		return nil, nil
	}

	return &user, err
}

// 更新用户信息
func (*userDao) Update(user *model.User) error {
	_, err := db.Cli.Exec("UPDATE user SET nickname = ?,sex = ?,avatar_url = ?,extra = ? WHERE app_id = ? AND user_id = ?",
		user.Nickname, user.Sex, user.AvatarUrl, user.Extra, user.AppId, user.UserId)
	if err != nil {
		return err
	}
	return nil
}
