package service

import (
	"go-IM/logic/model"
)

type friendService struct{}

var FriendService = new(friendService)

// AddFriend 新增好友
func (*friendService) AddFriend(uid, fid int64) error {
	if _, err := model.LoadUser(fid); err != nil {
		return err
	}
	um, err := model.LoadUser(uid)
	if err != nil {
		return err
	}
	// 检查重复添加
	for _, uid := range um.Friends {
		if fid == uid {
			return nil
		}
	}
	um.Friends = append(um.Friends, fid)
	return model.StoreUser(um)
}

// QueryFriends 查询好友列表
func (*friendService) QueryFriends(userId int64) ([]*model.User, error) {
	um, err := model.LoadUser(userId)
	if err != nil {
		return nil, err
	}
	flist := make([]*model.User, 0)
	for _, uid := range um.Friends {
		fm, err := model.LoadUser(uid)
		if err != nil {
			continue
		}
		flist = append(flist, fm)
	}
	return flist, nil
}
