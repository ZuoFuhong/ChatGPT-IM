package service

import (
	"go-IM/logic/dao"
	"go-IM/logic/model"
	"go-IM/pkg/errs"
	"go-IM/pkg/util"
)

type groupService struct{}

var GroupService = new(groupService)

// Get 获取群组信息
func (*groupService) Get(groupId int64) *model.Group {
	group, err := dao.GroupDao.Get(groupId)
	if err != nil {
		panic(err)
	}
	return group
}

// Create 创建群组
func (*groupService) Create(group *model.Group) int64 {
	sf, _ := util.NewSnowflake(0, 0)
	groupId := sf.NextVal()
	group.GroupId = groupId
	affected, err := dao.GroupDao.Add(group)
	if err != nil {
		panic(err)
	}
	if affected == 0 {
		panic(errs.NewHttpErr(errs.Group, "Group already exist"))
	}
	return groupId
}

// Update 更新群组
func (*groupService) Update(group *model.Group) {
	err := dao.GroupDao.Update(group.AppId, group.GroupId, group.Name, group.AvatarUrl, group.Introduction, group.Extra)
	if err != nil {
		panic(err)
	}
}

// AddUser 给群组添加用户
func (*groupService) AddUser(groupUser *model.GroupUser) {
	group := GroupService.Get(groupUser.GroupId)
	if group == nil {
		panic(errs.NewHttpErr(errs.Group, "Group not exist"))
	}
	user := UserService.Get(groupUser.UserId)
	if user == nil {
		panic(errs.NewHttpErr(errs.Group, "user not exist"))
	}
	if group.Type == model.GroupTypeGroup {
		GroupUserService.AddUser(groupUser)
	}
}

// UpdateUser 更新群组用户
func (*groupService) UpdateUser(groupUser *model.GroupUser) {
	group := GroupService.Get(groupUser.GroupId)
	if group == nil {
		panic(errs.NewHttpErr(errs.Group, "Group not exist"))
	}
	if group.Type == model.GroupTypeGroup {
		GroupUserService.Update(groupUser)
	}
}

// DeleteUser 删除用户群组
func (*groupService) DeleteUser(groupId, userId int64) {
	group := GroupService.Get(groupId)
	if group == nil {
		panic(errs.NewHttpErr(errs.Group, "Group not exist"))
	}
	if group.Type == model.GroupTypeGroup {
		GroupUserService.DeleteUser(groupId, userId)
	}
}
