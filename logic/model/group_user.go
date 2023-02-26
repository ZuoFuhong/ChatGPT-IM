package model

// GroupUser 群组成员
type GroupUser struct {
	GroupId int64  `json:"group_id"` // 群组ID
	UserId  int64  `json:"user_id"`  // 用户ID
	Label   string `json:"label"`    // 用户标签
	Extra   string `json:"extra"`    // 群组用户附件属性
	Ctime   int64  `json:"ctime"`    // 创建时间
	Utime   int64  `json:"utime"`    // 更新时间
}

func (*GroupUser) BucketName() string {
	return "group_user"
}
