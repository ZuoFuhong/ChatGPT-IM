package defs

import "go-IM/logic/model"

type PackageType int

const (
	PackageType_SignIn      PackageType = 1
	PackageType_SYNC        PackageType = 2
	PackageType_HEARTBEAT   PackageType = 3
	PackageType_MESSAGE_ACK PackageType = 4
)

type Input struct {
	Type      PackageType
	RequestId int
	Data      string
}

type Output struct {
	Type      PackageType
	RequestId int
	Code      int
	Message   string
	Data      interface{}
}

type SignIn struct {
	AppId    string
	UserId   string
	DeviceId string
	Token    string
}

type SyncInput struct {
	Seq string
}

type SyncOutput struct {
	Messages []model.Message
}

type MessageACK struct {
	DeviceAck   int64
	ReceiveTime int64
}

// 发送者类型
type SenderType int32

const (
	SenderType_ST_SYSTEM SenderType = 1 // 系统
	SenderType_ST_USER   SenderType = 2 // 用户
)

type Sender struct {
	AppId      int64      // appId
	SenderType SenderType // 发送者类型
	SenderId   int64      // 发送者id
	DeviceId   int64      // 发送者设备id
}

// 接收者类型
type ReceiverType int32

const (
	ReceiverType_RT_USER         ReceiverType = 1 // 用户
	ReceiverType_RT_NORMAL_GROUP ReceiverType = 2 // 普通群组
	ReceiverType_RT_LARGE_GROUP  ReceiverType = 3 // 超大群组
)

type SendMessageReq struct {
	ReceiverType   ReceiverType
	ReceiverId     int64
	ToUserIds      []int64
	MessageType    int
	MessageContent []byte
	IsPersist      bool
}

// 消息状态
type MessageStatus int32

const (
	MessageStatus_MS_NORMAL MessageStatus = 1 // 未处理
	MessageStatus_MS_RECALL MessageStatus = 2 // 消息撤回
)
