package service

import (
	"ChatGPT-IM/backend/logic/model"
	"ChatGPT-IM/backend/pkg/tinyid"
	"time"
)

type userService struct{}

var UserService = new(userService)

// Add 注册用户
func (*userService) Add(um *model.User) (int64, error) {
	now := time.Now()
	uid := tinyid.NextId()
	um.Id = uid
	um.Ctime = now.UnixMilli()
	um.Utime = now.UnixMilli()
	if err := model.StoreUser(um); err != nil {
		return 0, err
	}
	return uid, nil
}

// UserInfo 获取用户信息
func (*userService) UserInfo(userId int64) (*model.User, error) {
	return model.LoadUser(userId)
}

// Update 更新用户信息
func (*userService) Update(um *model.User) error {
	umInfo, err := model.LoadUser(um.Id)
	if err != nil {
		return err
	}
	umInfo.Nickname = um.Nickname
	umInfo.Sex = um.Sex
	umInfo.AvatarUrl = um.AvatarUrl
	umInfo.Extra = um.Extra
	umInfo.Utime = time.Now().UnixMilli()
	return model.StoreUser(umInfo)
}
