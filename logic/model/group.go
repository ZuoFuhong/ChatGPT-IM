package model

import "time"

const (
	GroupTypeGroup    = 1 // 群组
	GroupTypeChatRoom = 2 // 聊天室
)

// Group 群组
type Group struct {
	Id           int64     `json:"-"`            // 群组id
	AppId        int64     `json:"-"`            // appId
	GroupId      int64     `json:"group_id"`     // 群组id
	Name         string    `json:"name"`         // 组名
	AvatarUrl    string    `json:"avatar_url"`   // 群头像
	Introduction string    `json:"introduction"` // 群简介
	UserNum      int32     `json:"user_num"`     // 群组人数
	Type         int32     `json:"type"`         // 群组类型
	Extra        string    `json:"extra"`        // 附加属性
	CreateTime   time.Time `json:"-"`            // 创建时间
	UpdateTime   time.Time `json:"-"`            // 更新时间
}
