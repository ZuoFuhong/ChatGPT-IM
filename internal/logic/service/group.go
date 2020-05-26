package service

import (
	"context"
	"errors"
	"go-IM/internal/logic/dao"
	"go-IM/internal/logic/model"
)

type groupService struct{}

var GroupService = new(groupService)

// Get 获取群组信息
func (*groupService) Get(ctx context.Context, appId, groupId int64) (*model.Group, error) {
	group, err := dao.GroupDao.Get(appId, groupId)
	if err != nil {
		return nil, err
	}
	if group == nil {
		return nil, nil
	}
	return group, nil
}

// Create 创建群组
func (*groupService) Create(ctx context.Context, group model.Group) error {
	affected, err := dao.GroupDao.Add(group)
	if err != nil {
		return err
	}

	if affected == 0 {
		return errors.New("group already exist")
	}
	return nil
}

// Update 更新群组
func (*groupService) Update(ctx context.Context, group model.Group) error {
	err := dao.GroupDao.Update(group.AppId, group.GroupId, group.Name, group.Introduction, group.Extra)
	return err
}

// AddUser 给群组添加用户
func (*groupService) AddUser(ctx context.Context, appId, groupId, userId int64, label, extra string) error {
	group, err := GroupService.Get(ctx, appId, groupId)
	if err != nil {
		return err
	}
	if group == nil {
		return errors.New("group not exist")
	}

	user, err := UserService.Get(appId, userId)
	if err != nil {
		return err
	}
	if user == nil {
		return errors.New("user not exist")
	}

	if group.Type == model.GroupTypeGroup {
		err = GroupUserService.AddUser(ctx, appId, groupId, userId, label, extra)
		if err != nil {
			return err
		}
	}
	return nil
}

// UpdateUser 更新群组用户
func (*groupService) UpdateUser(ctx context.Context, appId, groupId, userId int64, label, extra string) error {
	group, err := GroupService.Get(ctx, appId, groupId)
	if err != nil {
		return err
	}

	if group == nil {
		return errors.New("group not exist")
	}

	if group.Type == model.GroupTypeGroup {
		err = GroupUserService.Update(ctx, appId, groupId, userId, label, extra)
		if err != nil {
			return err
		}
	}
	return nil
}

// DeleteUser 删除用户群组
func (*groupService) DeleteUser(ctx context.Context, appId, groupId, userId int64) error {
	group, err := GroupService.Get(ctx, appId, groupId)
	if err != nil {
		return err
	}
	if group == nil {
		return errors.New("group not exist")
	}
	if group.Type == model.GroupTypeGroup {
		err = GroupUserService.DeleteUser(ctx, appId, groupId, userId)
		if err != nil {
			return err
		}
	}
	return nil
}
