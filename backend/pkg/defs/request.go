package defs

// RegisterDeviceForm 注册设备
type RegisterDeviceForm struct {
	UserId string `json:"user_id"`
	Type   int32  `json:"type"`
}

// RegisterUserForm 注册用户
type RegisterUserForm struct {
	Nickname  string `json:"nickname"`
	Sex       int32  `json:"sex"`
	AvatarUrl string `json:"avatar_url"`
	Extra     string `json:"extra"`
}

// UpdateUserForm 更新用户
type UpdateUserForm struct {
	UserId    string `json:"user_id"`
	Nickname  string `json:"nickname"`
	Sex       int32  `json:"sex"`
	AvatarUrl string `json:"avatar_url"`
	Extra     string `json:"extra"`
}

// AddFriendForm 新增好友
type AddFriendForm struct {
	UserId   string `json:"user_id"`
	FriendId string `json:"friend_id"`
}

type UserVO struct {
	UserId    string `json:"user_id"`    // 唯一用户ID
	Nickname  string `json:"nickname"`   // 昵称
	AvatarUrl string `json:"avatar_url"` // 用户头像
	Extra     string `json:"extra"`      // 附加属性
}

type GroupVO struct {
	GroupId      string `json:"group_id"`     // 群组id
	Name         string `json:"name"`         // 组名
	AvatarUrl    string `json:"avatar_url"`   // 群头像
	Introduction string `json:"introduction"` // 群简介
	Type         int32  `json:"type"`         // 群组类型
}
