package model

import "time"

const (
	MessageObjectTypeUser  = 1 // 用户
	MessageObjectTypeGroup = 2 // 群组
)

// Message 消息
type Message struct {
	Id             int64     // 主键
	AppId          int64     // appId
	ObjectType     int       // 所属类型
	ObjectId       int64     // 所属类型ID
	RequestId      int64     // 请求ID
	SenderType     int32     // 发送者类型
	SenderId       int64     // 发送者账户ID
	SenderDeviceId int64     // 发送者设备ID
	ReceiverType   int32     // 接收者类型
	ReceiverId     int64     // 接受者ID，如果是单聊信息，则为user_id，如果是群组消息，则为group_id
	ToUserIds      string    // 需要@的用户ID列表，多个用户 , 隔开
	Type           int       // 消息类型
	Content        string    // 消息内容
	Seq            int64     // 消息同步序列
	SendTime       time.Time // 消息发送时间
	Status         int32     // 创建时间
}
