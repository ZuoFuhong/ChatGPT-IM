package defs

// 注册设备
type RegisterDeviceForm struct {
	AppId         int64  `json:"app_id"`
	UserId        int64  `json:"user_id"`
	Type          int32  `json:"type"`
	Brand         string `json:"brand"`
	Model         string `json:"model"`
	SystemVersion string `json:"system_version"`
	SDKVersion    string `json:"sdk_version"`
	Status        int32  `json:"status"`
}

// 注册用户
type RegisterUserForm struct {
	AppId     int64  `json:"app_id"`
	Nickname  string `json:"nickname"`
	Sex       int32  `json:"sex"`
	AvatarUrl string `json:"avatar_url"`
	Extra     string `json:"extra"`
}

// 更新用户
type UpdateUserForm struct {
	AppId     int64  `json:"app_id"`
	UserId    int64  `json:"user_id"`
	Nickname  string `json:"nickname"`
	Sex       int32  `json:"sex"`
	AvatarUrl string `json:"avatar_url"`
	Extra     string `json:"extra"`
}

// 创建群组
type CreateGroupForm struct {
	AppId        int64  `json:"app_id"`
	Name         string `json:"name"`
	Introduction string `json:"introduction"`
	UserNum      int32  `json:"user_num"`
	Type         int32  `json:"type"`
	Extra        string `json:"extra"`
}

// 更新群组
type UpdateGroupForm struct {
	AppId        int64  `json:"app_id"`
	GroupId      int64  `json:"group_id"`
	Name         string `json:"name"`
	Introduction string `json:"introduction"`
	Extra        string `json:"extra"`
}

// 群组添加用户
type GroupAddUserForm struct {
	AppId   int64  `json:"app_id"`
	GroupId int64  `json:"group_id"`
	UserId  int64  `json:"user_id"`
	Label   string `json:"label"`
	Extra   string `json:"extra"`
}

// 更新群组用户
type UpdateGroupUserForm struct {
	AppId   int64  `json:"app_id"`
	GroupId int64  `json:"group_id"`
	UserId  int64  `json:"user_id"`
	Label   string `json:"label"`
	Extra   string `json:"extra"`
}
