package model

import (
	"encoding/json"
	"fmt"
)

// User 用户信息
type User struct {
	Id        int64   `json:"id"`         // 用户ID
	Nickname  string  `json:"nickname"`   // 昵称
	Sex       int32   `json:"sex"`        // 性别，1:男；2:女
	AvatarUrl string  `json:"avatar_url"` // 用户头像
	Devices   []int64 `json:"devices"`    // 设备列表
	Friends   []int64 `json:"friends"`    // 好友列表
	Extra     string  `json:"extra"`      // 附加属性
	Ctime     int64   `json:"ctime"`      // 创建时间
	Utime     int64   `json:"utime"`      // 更新时间
}

func (*User) BucketName() string {
	return "user"
}

// StoreUser 存储用户信息
func StoreUser(um *User) error {
	bytes, err := json.Marshal(um)
	if err != nil {
		return err
	}
	return WriteToDB(um.BucketName(), fmt.Sprint(um.Id), bytes)
}

// LoadUser 读取用户信息
func LoadUser(uid int64) (*User, error) {
	um := &User{}
	bytes, err := ReadFromDB(um.BucketName(), fmt.Sprint(uid))
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(bytes, um); err != nil {
		return nil, err
	}
	return um, nil
}
