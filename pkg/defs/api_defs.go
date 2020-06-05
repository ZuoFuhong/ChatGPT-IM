package defs

// 注册设备
type RegisterDeviceForm struct {
	AppId         string `json:"app_id"`
	UserId        string `json:"user_id"`
	Type          int32  `json:"type"`
	Brand         string `json:"brand"`
	Model         string `json:"model"`
	SystemVersion string `json:"system_version"`
	SDKVersion    string `json:"sdk_version"`
	Status        int32  `json:"status"`
}

// 注册用户
type RegisterUserForm struct {
	AppId     string `json:"app_id"`
	Nickname  string `json:"nickname"`
	Sex       int32  `json:"sex"`
	AvatarUrl string `json:"avatar_url"`
	Extra     string `json:"extra"`
}

// 更新用户
type UpdateUserForm struct {
	AppId     string `json:"app_id"`
	UserId    string `json:"user_id"`
	Nickname  string `json:"nickname"`
	Sex       int32  `json:"sex"`
	AvatarUrl string `json:"avatar_url"`
	Extra     string `json:"extra"`
}

// 创建群组
type CreateGroupForm struct {
	AppId        string `json:"app_id"`
	Name         string `json:"name"`
	AvatarUrl    string `json:"avatar_url"`
	Introduction string `json:"introduction"`
	UserNum      int32  `json:"user_num"`
	Type         int32  `json:"type"`
	Extra        string `json:"extra"`
}

// 更新群组
type UpdateGroupForm struct {
	AppId        string `json:"app_id"`
	GroupId      string `json:"group_id"`
	Name         string `json:"name"`
	AvatarUrl    string `json:"avatar_url"`
	Introduction string `json:"introduction"`
	Extra        string `json:"extra"`
}

// 群组添加用户
type GroupAddUserForm struct {
	AppId   string `json:"app_id"`
	GroupId string `json:"group_id"`
	UserId  string `json:"user_id"`
	Label   string `json:"label"`
	Extra   string `json:"extra"`
}

// 更新群组用户
type UpdateGroupUserForm struct {
	AppId   string `json:"app_id"`
	GroupId string `json:"group_id"`
	UserId  string `json:"user_id"`
	Label   string `json:"label"`
	Extra   string `json:"extra"`
}

// 新增好友
type AddFriendForm struct {
	AppId    string `json:"app_id"`
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
