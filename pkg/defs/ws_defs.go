package defs

import "go-IM/internal/logic/model"

type PackageType int

const (
	PackageType_SYNC      PackageType = 1
	PackageType_HEARTBEAT PackageType = 2
	PackageType_MESSAGE   PackageType = 3
)

type Input struct {
	Type      PackageType
	RequestId int
	Data      []byte
}

type Output struct {
	Type      PackageType
	RequestId int
	Code      int
	Message   string
	Data      []byte
}

type SyncInput struct {
	Seq int64
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
	SenderType_ST_UNKNOWN  SenderType = 0
	SenderType_ST_SYSTEM   SenderType = 1
	SenderType_ST_USER     SenderType = 2
	SenderType_ST_BUSINESS SenderType = 3
)

type Sender struct {
	AppId      int64      // appId
	SenderType SenderType // 发送者类型，1：系统，2：用户，3：业务方
	SenderId   int64      // 发送者id
	DeviceId   int64      // 发送者设备id
}

// 接收者类型
type ReceiverType int32

const (
	ReceiverType_RT_UNKNOWN      ReceiverType = 0
	ReceiverType_RT_USER         ReceiverType = 1
	ReceiverType_RT_NORMAL_GROUP ReceiverType = 2
	ReceiverType_RT_LARGE_GROUP  ReceiverType = 3
)

// 消息类型
type MessageType int32

// 消息体
type MessageBody struct {
	MessageType    MessageType
	MessageContent []byte
}

type SendMessageReq struct {
	MessageId    string
	ReceiverType ReceiverType
	ReceiverId   int64
	ToUserIds    []int64
	MessageBody  *MessageBody
	SendTime     int64
	IsPersist    bool
}
