package model

import "time"

// Friend 好友
type Friend struct {
	Id         int64     `json:"-"`         // 主键
	AppId      int64     `json:"app_id"`    // AppID
	UserId     int64     `json:"user_id"`   // 用户ID
	FriendId   int64     `json:"friend_id"` // 好友ID
	CreateTime time.Time `json:"-"`
	UpdateTime time.Time `json:"-"`
}
