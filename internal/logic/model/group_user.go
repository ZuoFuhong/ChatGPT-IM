package model

import "time"

type GroupUser struct {
	Id         int64     `json:"id,omitempty"` // 自增主键
	AppId      int64     `json:"app_id"`       // app_id
	GroupId    int64     `json:"group_id"`     // 群组id
	UserId     int64     `json:"user_id"`      // 用户id
	Label      string    `json:"label"`        // 用户标签
	Extra      string    `json:"extra"`        // 群组用户附件属性
	CreateTime time.Time `json:"-"`            // 创建时间
	UpdateTime time.Time `json:"-"`            // 更新时间
}
