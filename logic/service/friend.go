package service

import (
	"go-IM/logic/dao"
	"go-IM/logic/model"
)

type friendService struct{}

var FriendService = new(friendService)

// 新增好友
func (*friendService) AddFriend(appId, userId, friendId int64) {
	err := dao.FriendDao.Insert(appId, userId, friendId)
	if err != nil {
		panic(err)
	}
}

// 查询好友列表
func (*friendService) QueryFriends(appId, userId int64) *[]model.User {
	friends, err := dao.FriendDao.SelectFriends(appId, userId)
	if err != nil {
		panic(err)
	}
	return friends
}
