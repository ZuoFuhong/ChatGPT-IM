package service

import (
	"go-IM/logic/dao"
	"go-IM/logic/model"
	"go-IM/pkg/errs"
	"go-IM/pkg/util"
)

type userService struct{}

var UserService = new(userService)

// Add 添加用户（唯一用户ID）
// 1.添加用户，2.添加用户消息序列号
func (*userService) Add(user *model.User) int64 {
	sf, _ := util.NewSnowflake(0, 0)
	userId := sf.NextVal()
	user.UserId = userId
	affected, e := dao.UserDao.Add(user)
	if e != nil {
		panic(e)
	}
	if affected == 0 {
		panic(errs.NewHttpErr(errs.User, "user already exist"))
	}
	return userId
}

// Get 获取用户信息
func (*userService) Get(appId, userId int64) *model.User {
	user, err := dao.UserDao.Get(appId, userId)
	if err != nil {
		panic(err)
	}
	return user
}

// Get 更新用户信息
func (*userService) Update(user *model.User) {
	err := dao.UserDao.Update(user)
	if err != nil {
		panic(err)
	}
}
