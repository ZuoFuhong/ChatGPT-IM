package dao

import (
	"database/sql"
	"go-IM/logic/db"
	"go-IM/logic/model"
)

type groupUserDao struct{}

var GroupUserDao = new(groupUserDao)

// 获取用户加入的群组信息
func (*groupUserDao) ListByUserId(appId, userId int64) ([]model.Group, error) {
	rows, err := db.Cli.Query(
		"SELECT g.group_id,g.name,g.introduction,g.user_num,g.type,g.extra,g.create_time,g.update_time "+
			"FROM group_user u "+
			"LEFT JOIN `group` g on u.app_id = g.app_id and u.group_id = g.group_id "+
			"WHERE u.app_id = ? and u.user_id = ?",
		appId, userId)
	if err != nil {
		return nil, err
	}
	var groups []model.Group
	var group model.Group
	for rows.Next() {
		err := rows.Scan(&group.GroupId, &group.Name, &group.Introduction, &group.UserNum, &group.Type, &group.Extra, &group.CreateTime, &group.UpdateTime)
		if err != nil {
			return nil, err
		}
		groups = append(groups, group)
	}
	return groups, nil
}

// ListGroupUser 获取群组用户信息
func (*groupUserDao) ListUser(appId, groupId int64) ([]model.GroupUser, error) {
	rows, err := db.Cli.Query(`
		SELECT user_id,label,extra,create_time,update_time 
		FROM group_user
		WHERE app_id = ?
		AND group_id = ?`, appId, groupId)
	if err != nil {
		return nil, err
	}
	groupUsers := make([]model.GroupUser, 0, 5)
	for rows.Next() {
		var groupUser = model.GroupUser{
			AppId:   appId,
			GroupId: groupId,
		}
		err := rows.Scan(&groupUser.UserId, &groupUser.Label, &groupUser.Extra, &groupUser.CreateTime, &groupUser.UpdateTime)
		if err != nil {
			return nil, err
		}
		groupUsers = append(groupUsers, groupUser)
	}
	return groupUsers, nil
}

// 获取群组用户信息,用户不存在返回nil
func (*groupUserDao) Get(appId, groupId, userId int64) (*model.GroupUser, error) {
	var groupUser = model.GroupUser{
		AppId:   appId,
		GroupId: groupId,
		UserId:  userId,
	}
	err := db.Cli.QueryRow("SELECT label,extra FROM group_user WHERE app_id = ? AND group_id = ? AND user_id = ?",
		appId, groupId, userId).
		Scan(&groupUser.Label, &groupUser.Extra)
	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}

	if err == sql.ErrNoRows {
		return nil, nil
	}

	return &groupUser, nil
}

// 将用户添加到群组
func (*groupUserDao) Add(appId, groupId, userId int64, label, extra string) error {
	_, err := db.Cli.Exec("INSERT IGNORE INTO group_user(app_id,group_id,user_id,label,extra) VALUES (?,?,?,?,?)",
		appId, groupId, userId, label, extra)
	if err != nil {
		return err
	}
	return nil
}

// 将用户从群组删除
func (d *groupUserDao) Delete(appId int64, groupId int64, userId int64) error {
	_, err := db.Cli.Exec("DELETE FROM group_user WHERE app_id = ? AND group_id = ? AND user_id = ?",
		appId, groupId, userId)
	return err
}

// 更新用户群组信息
func (*groupUserDao) Update(appId, groupId, userId int64, label string, extra string) error {
	_, err := db.Cli.Exec("UPDATE group_user SET label = ?,extra = ? WHERE app_id = ? AND group_id = ? AND user_id = ?",
		label, extra, appId, groupId, userId)
	return err
}
