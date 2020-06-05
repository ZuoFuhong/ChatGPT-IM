package service

import (
	"go-IM/logic/dao"
	"go-IM/logic/model"
)

type groupUserService struct{}

var GroupUserService = new(groupUserService)

// ListByUserId 获取用户所加入的群组
func (*groupUserService) ListByUserId(appId, userId int64) []model.Group {
	groups, err := dao.GroupUserDao.ListByUserId(appId, userId)
	if err != nil {
		panic(err)
	}
	return groups
}

// GetUsers 获取群组的所有用户信息
func (*groupUserService) GetUsers(groupId int64) []model.GroupUser {
	users, err := dao.GroupUserDao.ListUser(groupId)
	if err != nil {
		panic(err)
	}
	return users
}

// AddUser 给群组添加用户
func (*groupUserService) AddUser(groupUser *model.GroupUser) {
	err := dao.GroupUserDao.Add(groupUser.AppId, groupUser.GroupId, groupUser.UserId, groupUser.Label, groupUser.Extra)
	if err != nil {
		panic(err)
	}

	err = dao.GroupDao.UpdateUserNum(groupUser.GroupId, 1)
	if err != nil {
		panic(err)
	}
}

// DeleteUser 从群组移除用户
func (*groupUserService) DeleteUser(groupId, userId int64) {
	err := dao.GroupUserDao.Delete(groupId, userId)
	if err != nil {
		panic(err)
	}

	err = dao.GroupDao.UpdateUserNum(groupId, -1)
	if err != nil {
		panic(err)
	}
}

// Update 更新群组用户信息
func (*groupUserService) Update(groupUser *model.GroupUser) {
	err := dao.GroupUserDao.Update(groupUser.AppId, groupUser.GroupId, groupUser.UserId, groupUser.Label, groupUser.Extra)
	if err != nil {
		panic(err)
	}
}
