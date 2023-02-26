package service

import (
	"go-IM/logic/model"
	"go-IM/pkg/tinyid"
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
	now := time.Now()
	um.Utime = now.UnixMilli()
	return model.StoreUser(um)
}
