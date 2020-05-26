package service

import (
	"errors"
	"go-IM/internal/logic/dao"
	"go-IM/internal/logic/model"
)

type userService struct{}

var UserService = new(userService)

// Add 添加用户（将业务账号导入IM系统账户）
// 1.添加用户，2.添加用户消息序列号
func (*userService) Add(user model.User) error {
	affected, e := dao.UserDao.Add(user)
	if e != nil {
		return e
	}
	if affected == 0 {
		return errors.New("user already exist")
	}
	return nil
}

// Get 获取用户信息
func (*userService) Get(appId, userId int64) (*model.User, error) {
	user, err := dao.UserDao.Get(appId, userId)
	return user, err
}

// Get 获取用户信息
func (*userService) Update(user model.User) error {
	err := dao.UserDao.Update(user)
	return err
}
